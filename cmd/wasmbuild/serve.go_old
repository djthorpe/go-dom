package main

import (
	"bytes"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/djthorpe/go-dom/etc"
)

// SSEMessage represents a message sent to SSE clients
type SSEMessage struct {
	Type string // "reload" or "error"
	Data string // error message for "error" type
}

// reloadBroadcaster manages SSE connections and broadcasts reload events
type reloadBroadcaster struct {
	mu      sync.Mutex
	clients map[chan SSEMessage]bool
}

func newReloadBroadcaster() *reloadBroadcaster {
	return &reloadBroadcaster{
		clients: make(map[chan SSEMessage]bool),
	}
}

func (rb *reloadBroadcaster) register(client chan SSEMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	rb.clients[client] = true
}

func (rb *reloadBroadcaster) unregister(client chan SSEMessage) {
	rb.mu.Lock()
	defer rb.mu.Unlock()
	delete(rb.clients, client)
	close(client)
}

func (rb *reloadBroadcaster) broadcast(msg SSEMessage) {
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

func (s *ServeCmd) Run(ctx *Context) error {
	// Use command-specific Path, default to "." if not specified
	path := s.Path
	if path == "" {
		path = "."
	}

	// Resolve the path to get the directory name
	absPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("failed to resolve path: %w", err)
	}

	// Get the directory name for the output filename
	dirName := filepath.Base(absPath)

	// Determine build directory - create it once if not specified
	buildDir := CLI.Output
	if buildDir == "" {
		// Create a single temporary directory for all compilations
		tmpDir, err := os.MkdirTemp("", "wasm-build-*")
		if err != nil {
			return fmt.Errorf("failed to create temp directory: %w", err)
		}
		buildDir = tmpDir
		logger.Infof("Using temporary build directory: %s", buildDir)
	}

	// Compile application first
	compile := &CompileCommand{
		Path:     path,
		DirName:  dirName,
		BuildDir: buildDir, // Always use the same build directory
		GoPath:   ctx.GoPath,
		GoFlags:  CLI.GoFlags,
	}

	_, err = compile.RunAndGetBuildDir()
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	// Verify the build directory exists
	if _, err := os.Stat(buildDir); err != nil {
		return fmt.Errorf("build directory does not exist: %w", err)
	}

	logger.Infof("Serving directory: %s", buildDir)
	logger.Infof("wasm_exec.js from: %s", ctx.WasmExecPath)

	// Always show listening address (with timestamp if verbose)
	if CLI.Verbose {
		logger.Infof("Listening on: %s", s.Listen)
	} else {
		fmt.Printf("Listening on: %s\n", s.Listen)
	}

	// Create a broadcaster for reload events
	broadcaster := newReloadBroadcaster()

	// Start file watcher if --watch flag is set
	if s.Watch {
		// Build list of paths to watch - start with the main application path
		watchPaths := []string{absPath}
		logger.Infof("Watching for changes in: %s", absPath)

		// Discover and add local package dependencies
		localDeps, err := discoverLocalDependencies(ctx.GoPath, absPath)
		if err != nil {
			logger.Errorf("Failed to discover local dependencies: %v", err)
		} else if len(localDeps) > 0 {
			logger.Infof("Discovered %d local dependencies to watch:", len(localDeps))
			for _, dep := range localDeps {
				logger.Infof("  - %s", dep)
			}
			watchPaths = append(watchPaths, localDeps...)
		}

		go s.watchAndRecompile(watchPaths, compile, broadcaster)
	}

	// Prepare bootstrap HTML with the actual WASM filename
	wasmFileName := dirName + ".wasm"
	bootstrapHTML := bytes.Replace(etc.BootstrapHTML, []byte("{{WASM_FILE}}"), []byte(wasmFileName), 1)

	// Replace {{LIBRARY}} with Bootstrap 5 HTML or empty string
	var libraryHTML []byte
	if s.BS5 {
		libraryHTML = etc.Bootstrap5
	}
	bootstrapHTML = bytes.Replace(bootstrapHTML, []byte("{{LIBRARY}}"), libraryHTML, 1)

	// Create custom handler that serves bootstrap.html at root and wasm_exec.js from GOROOT
	fs := http.FileServer(http.Dir(buildDir))
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Handle SSE endpoint for live reload
		if r.URL.Path == "/_notify" {
			s.handleSSE(w, r, broadcaster)
			return
		}
		// Serve bootstrap.html for root path and /index.html
		if r.URL.Path == "/" || r.URL.Path == "/index.html" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.Write(bootstrapHTML)
			return
		}
		// Serve favicon.png
		if r.URL.Path == "/favicon.png" {
			w.Header().Set("Content-Type", "image/png")
			w.Write(etc.FaviconPNG)
			return
		}
		// Serve notify.js
		if r.URL.Path == "/notify.js" {
			w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
			w.Write(etc.NotifyJS)
			return
		}
		// Serve notify.css
		if r.URL.Path == "/notify.css" {
			w.Header().Set("Content-Type", "text/css; charset=utf-8")
			w.Write(etc.NotifyCSS)
			return
		}
		// Serve wasm_exec.js from GOROOT
		if r.URL.Path == "/wasm_exec.js" {
			http.ServeFile(w, r, ctx.WasmExecPath)
			return
		}
		// Otherwise serve from the file system
		fs.ServeHTTP(w, r)
	})

	// Wrap handler with logging middleware
	loggedHandler := loggingMiddleware(handler)

	// Start HTTP server
	if err := http.ListenAndServe(s.Listen, loggedHandler); err != nil {
		return fmt.Errorf("server failed: %w", err)
	}

	return nil
}

// loggingMiddleware logs HTTP requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Infof("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// handleSSE handles Server-Sent Events for live reload
func (s *ServeCmd) handleSSE(w http.ResponseWriter, r *http.Request, broadcaster *reloadBroadcaster) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	logger.Infof("SSE client connected: %s", r.RemoteAddr)

	// Create a client channel and register it
	clientChan := make(chan SSEMessage, 10)
	broadcaster.register(clientChan)
	defer broadcaster.unregister(clientChan)

	// Send initial connection message
	fmt.Fprintf(w, "data: connected\n\n")
	flusher.Flush()

	// Keep connection alive and wait for reload signals
	for {
		select {
		case msg := <-clientChan:
			if msg.Type == "reload" {
				fmt.Fprintf(w, "data: reload\n\n")
				flusher.Flush()
				logger.Infof("Sent reload notification to: %s", r.RemoteAddr)
			} else if msg.Type == "error" {
				// For multi-line error messages, prefix each line with "data: "
				lines := strings.Split(msg.Data, "\n")
				fmt.Fprintf(w, "event: compileerror\n")
				for _, line := range lines {
					fmt.Fprintf(w, "data: %s\n", line)
				}
				fmt.Fprintf(w, "\n")
				flusher.Flush()
				logger.Infof("Sent error notification to: %s", r.RemoteAddr)
			}
		case <-r.Context().Done():
			logger.Infof("SSE client disconnected: %s", r.RemoteAddr)
			return
		}
	}
}

// watchAndRecompile monitors the source directories for .go file changes and recompiles
func (s *ServeCmd) watchAndRecompile(sourcePaths []string, compile *CompileCommand, broadcaster *reloadBroadcaster) {
	// Track modification times of .go files
	modTimes := make(map[string]time.Time)
	var mu sync.Mutex

	// Helper function to update modification times for all paths
	updateModTimes := func(paths []string) {
		for _, sourcePath := range paths {
			filepath.WalkDir(sourcePath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return nil
				}
				if !d.IsDir() && filepath.Ext(path) == ".go" {
					if info, err := d.Info(); err == nil {
						mu.Lock()
						modTimes[path] = info.ModTime()
						mu.Unlock()
					}
				}
				return nil
			})
		}
	}

	// Initial scan to populate modification times for all watched paths
	updateModTimes(sourcePaths)

	// Keep track of current watch paths
	currentWatchPaths := sourcePaths

	// Poll for changes every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		changed := false

		// Check all watched paths for changes
		for _, sourcePath := range currentWatchPaths {
			filepath.WalkDir(sourcePath, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return nil
				}
				if !d.IsDir() && filepath.Ext(path) == ".go" {
					info, err := d.Info()
					if err != nil {
						return nil
					}

					mu.Lock()
					oldTime, exists := modTimes[path]
					if !exists || info.ModTime().After(oldTime) {
						modTimes[path] = info.ModTime()
						changed = true
					}
					mu.Unlock()
				}
				return nil
			})
		}

		if changed {
			logger.Info("Change detected, recompiling...")
			result := compile.RunAndGetResult()
			if result.Error != nil {
				logger.Errorf("Recompilation failed: %v", result.Error)
				// Broadcast error to all connected SSE clients
				broadcaster.broadcast(SSEMessage{
					Type: "error",
					Data: result.Stderr,
				})
				logger.Info("Error notification broadcasted")
			} else {
				logger.Info("Recompilation successful")

				// Rediscover dependencies after successful compilation
				absPath := sourcePaths[0] // The main source path
				localDeps, err := discoverLocalDependencies(compile.GoPath, absPath)
				if err != nil {
					logger.Errorf("Failed to rediscover local dependencies: %v", err)
				} else {
					// Build new watch paths list
					newWatchPaths := []string{absPath}
					if len(localDeps) > 0 {
						logger.Infof("Rediscovered %d local dependencies to watch:", len(localDeps))
						for _, dep := range localDeps {
							logger.Infof("  - %s", dep)
						}
						newWatchPaths = append(newWatchPaths, localDeps...)
					}

					// Update current watch paths and scan new dependencies
					currentWatchPaths = newWatchPaths
					updateModTimes(currentWatchPaths)
				}

				// Broadcast reload to all connected SSE clients
				broadcaster.broadcast(SSEMessage{Type: "reload"})
				logger.Info("Reload notification broadcasted")
			}
		}
	}
}
