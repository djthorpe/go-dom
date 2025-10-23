package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/alecthomas/kong"
)

// Global context shared across commands
type Context struct {
	GoRoot       string
	GoPath       string
	WasmExecPath string
}

var CLI struct {
	Go       string `default:"go" help:"Path to go tool"`
	WasmExec string `default:"lib/wasm/wasm_exec.js" help:"Path to wasm_exec.js relative to GOROOT"`
	Output   string `help:"Output directory (uses temp dir if not specified)"`
	GoFlags  string `help:"Additional flags to pass to go build"`
	Verbose  bool   `short:"v" help:"Enable verbose output"`
	Config   string `default:"wasmbuild.yaml" help:"Path to configuration file"`

	Compile CompileCmd `cmd:"" help:"Compile a WASM application"`
	Serve   ServeCmd   `cmd:"" help:"Serve a WASM application"`
}

type CompileCmd struct {
	Path string `arg:"" default:"." help:"Path for the wasm application"`
}

type ServeCmd struct {
	Path   string `arg:"" default:"." help:"Path to the wasm application source"`
	Listen string `short:"l" default:"localhost:9090" help:"Address to listen on (e.g., localhost:9090 or 0.0.0.0:9090)"`
	Watch  bool   `short:"w" help:"Watch for changes and recompile automatically (includes local dependencies)"`
	BS5    bool   `help:"Include Bootstrap 5 library in the HTML template"`
}

func main() {
	ctx := kong.Parse(&CLI)

	// Initialize logger based on verbose flag
	InitLogger(CLI.Verbose)

	// Initialize shared context
	sharedCtx, err := initContext()
	if err != nil {
		ctx.Fatalf("Initialization failed: %v", err)
	}

	// Run the selected command
	err = ctx.Run(sharedCtx)
	ctx.FatalIfErrorf(err)
}

func initContext() (*Context, error) {
	sharedCtx := &Context{}

	// Get GOROOT from environment
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		// If GOROOT is not set, try to determine it from the go tool
		goPath := CLI.Go
		if !filepath.IsAbs(goPath) {
			var err error
			goPath, err = exec.LookPath(goPath)
			if err != nil {
				return nil, fmt.Errorf("failed to locate go executable: %w", err)
			}
		}

		// Run 'go env GOROOT' to get GOROOT
		cmd := exec.Command(goPath, "env", "GOROOT")
		output, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("failed to determine GOROOT: %w", err)
		}
		goroot = string(output[:len(output)-1]) // Remove trailing newline
	}
	logger.Infof("Using GOROOT: %s", goroot)
	sharedCtx.GoRoot = goroot

	// Determine wasm_exec.js location (absolute or relative to GOROOT)
	wasmExecPath := CLI.WasmExec
	if !filepath.IsAbs(wasmExecPath) {
		wasmExecPath = filepath.Join(goroot, wasmExecPath)
	}

	// Verify wasm_exec.js exists
	if _, err := os.Stat(wasmExecPath); err != nil {
		return nil, fmt.Errorf("wasm_exec.js not found at %s: %w", wasmExecPath, err)
	}
	logger.Infof("Using wasm_exec.js: %s", wasmExecPath)
	sharedCtx.WasmExecPath = wasmExecPath

	// Locate the go executable
	goPath := CLI.Go
	if !filepath.IsAbs(goPath) {
		var err error
		goPath, err = exec.LookPath(goPath)
		if err != nil {
			return nil, fmt.Errorf("failed to locate go executable: %w", err)
		}
	}
	logger.Infof("Using go executable: %s", goPath)
	sharedCtx.GoPath = goPath

	return sharedCtx, nil
}

func (c *CompileCmd) Run(ctx *Context) error {
	// Use command-specific Path, default to "." if not specified
	path := c.Path
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

	// Parse the configuration file
	config, err := ParseYAMLPath(CLI.Config, absPath)
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	fmt.Println(config)

	// Compile application
	compile := &CompileCommand{
		Path:     path,
		DirName:  dirName,
		BuildDir: CLI.Output,
		GoPath:   ctx.GoPath,
		GoFlags:  CLI.GoFlags,
	}

	buildDir, err := compile.RunAndGetBuildDir()
	if err != nil {
		return err
	}

	// Show the build directory when not in verbose mode
	if !CLI.Verbose {
		fmt.Println(buildDir)
	}

	return nil
}
