package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type ServeCmd struct {
	BuildPath
	WatchFlag
	Listen string `default:"localhost:9090" help:"Address to listen on (e.g., localhost:9090 or 0.0.0.0:9090)"`
}

type WatchFlag struct {
	Watch bool `short:"w" help:"Watch for changes in dependencies"`
}

type ServeContext struct {
	BuildContext
	Listen string `json:"listen"`
	Watch  bool   `json:"watch"`

	// Broadcast notifications to clients
	broadcaster ServeBroadcaster `json:"-"`
}

// ServeMessage represents a message sent to SSE clients
type ServeMessage struct {
	Type string // "reload" or "error"
	Data string
}

type ServeBroadcaster struct {
	mu      sync.Mutex
	clients map[chan ServeMessage]bool
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewServeBroadcaster() *ServeBroadcaster {
	return &ServeBroadcaster{
		clients: make(map[chan ServeMessage]bool),
	}
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *ServeCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveFile(ctx.Config, c.Path)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create the server context from the configuration
	serveContext, err := config.ServeContext(ctx, c.Path, c.Listen, c.Watch)
	if err != nil {
		return err
	}

	// Compile the .wasm file
	file, err := serveContext.CompileExec(ctx)
	if err != nil {
		return err
	}

	// Log the serve context
	ctx.log.Info("Serve: ", serveContext)

	// Start the server
	return serveContext.Serve(ctx,
		file,
		serveContext.WasmExecJS,
		serveContext.WasmExecHTML,
		serveContext.FavIcon,
	)
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *ServeContext) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// ServeContext creates a ServeContext from the Config, returning all the
// information needed to build and serve a WASM application.
func (c Config) ServeContext(ctx *Context, path, listen string, watch bool) (*ServeContext, error) {
	// Create a compiler context from the configuration
	buildContext, err := c.BuildContext(ctx, path, "")
	if err != nil {
		return nil, err
	}

	return &ServeContext{
		BuildContext: *buildContext,
		Listen:       listen,
		Watch:        watch,
	}, nil
}

func (c *ServeContext) Serve(ctx *Context, files ...*File) error {
	handler := http.NewServeMux()

	// Serve files
	for _, f := range files {
		if f != nil {
			handler.Handle(f.URL(), f.Handler())
		}
	}

	// Serve assets
	for _, asset := range c.Assets {
		if filepath.IsAbs(asset) == false {
			asset = filepath.Join(c.Path, asset)
		}
		if info, err := os.Stat(asset); err != nil {
			return err
		} else if info.Mode().IsRegular() {
			file, err := NewFileFromSource(asset, filepath.Base(asset))
			if err != nil {
				return err
			} else {
				handler.Handle(file.URL(), file.Handler())
			}
		}
	}

	// Output listening info
	url, err := url.Parse(fmt.Sprintf("http://%s%s", c.Listen, c.WasmExecHTML.URL()))
	if err != nil {
		return err
	}
	fmt.Println(url.String())

	// If watch flag is set, we build a dependency watcher
	var wg sync.WaitGroup
	if c.Watch {
		dep, err := c.BuildContext.DepContext(ctx)
		if err != nil {
			return err
		}

		// Watch for dependency changes
		wg.Add(1)
		go func() {
			defer wg.Done()
			dep.Run(ctx.ctx)
		}()

		// Respond to modification events
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.ctx.Done():
					return
				case <-dep.modified:
					// Re-compile the .wasm file
					file, err := c.CompileExec(ctx)
					if err != nil {
						c.broadcaster.broadcast(ServeMessage{
							Type: "error",
							Data: fmt.Sprintf("Compilation error: %v", err),
						})
						ctx.log.Error("Failed to compile after modification: ", err)
						continue
					} else {
						// TODO: Update the handler with the new file data
						c.broadcaster.broadcast(ServeMessage{
							Type: "reload",
							Data: file.Path,
						})
					}
				}
			}
		}()

	}

	// Start HTTP server
	err = http.ListenAndServe(c.Listen, logging(handler, ctx.log))

	// Wait for all go-routines to end
	wg.Wait()

	// Return any errors
	return err
}

func (rb *ServeBroadcaster) register(client chan ServeMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.clients[client] = true
}

func (rb *ServeBroadcaster) unregister(client chan ServeMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	delete(rb.clients, client)
	close(client)
}

func (rb *ServeBroadcaster) broadcast(msg ServeMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	for client := range rb.clients {
		select {
		case client <- msg:
		default:
			// Client not ready, skip
		}
	}
}

// logging middleware
func logging(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// notify handler
func (c *ServeContext) NotifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// We need to be able to flush the data
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	// Create a client channel and register it
	notify := make(chan ServeMessage)
	c.broadcaster.register(notify)
	defer c.broadcaster.unregister(notify)

	// Send initial connection message
	fmt.Fprintf(w, "data: connected\n\n")
	flusher.Flush()

	// Keep connection alive and wait for reload signals
	for {
		select {
		case msg := <-notify:
			switch msg.Type {
			case "reload":
				fmt.Fprintf(w, "data: reload\n\n")
				flusher.Flush()
			case "error":
				// For multi-line error messages, prefix each line with "data: "
				lines := strings.Split(msg.Data, "\n")
				fmt.Fprintf(w, "event: error\n")
				for _, line := range lines {
					fmt.Fprintf(w, "data: %s\n", line)
				}
				fmt.Fprintf(w, "\n")
				flusher.Flush()
			}
		case <-r.Context().Done():
			return
		}
	}
}
