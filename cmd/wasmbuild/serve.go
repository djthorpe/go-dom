package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type ServeCmd struct {
	BuildPath
	Listen string `default:"localhost:9090" help:"Address to listen on (e.g., localhost:9090 or 0.0.0.0:9090)"`
	Watch  bool   `short:"w" help:"Watch for changes and recompile automatically"`
}

type ServeContext struct {
	BuildContext
	Listen string `json:"listen"`
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *ServeCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveDir(ctx.Config, c.Path, true)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create the server context from the configuration
	serveContext, err := config.ServeContext(ctx, c.Path, c.Listen)
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
func (c Config) ServeContext(ctx *Context, path, listen string) (*ServeContext, error) {
	// Create a compiler context from the configuration
	buildContext, err := c.BuildContext(ctx, path, "")
	if err != nil {
		return nil, err
	}

	return &ServeContext{
		BuildContext: *buildContext,
		Listen:       listen,
	}, nil
}

func (c *ServeContext) Serve(ctx *Context, files ...*File) error {
	handler := http.NewServeMux()

	// TODO: Build dependency watcher

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

	// Start HTTP server
	return http.ListenAndServe(c.Listen, logging(handler, ctx.log))
}

// logging middleware
func logging(next http.Handler, logger *Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
