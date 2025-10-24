package main

import (
	// Packages
	"context"
	"os"
	"os/signal"

	"github.com/alecthomas/kong"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Context struct {
	Go       string `default:"go" help:"Path to go tool"`
	WasmExec string `default:"lib/wasm/wasm_exec.js:misc/wasm/wasm_exec.js" help:"Path to wasm_exec.js relative to GOROOT"`
	GoFlags  string `help:"Additional flags to pass to go build"`
	Config   string `default:"wasmbuild.yaml" help:"Path to configuration YAML file (relative to source path)"`
	Verbose  bool   `short:"v" help:"Enable verbose output"`

	// Private
	log *Logger
	ctx context.Context
}

type CLI struct {
	Context
	Build BuildCmd `cmd:"" help:"Build a WASM application"`
	Serve ServeCmd `cmd:"" help:"Serve a WASM application"`
	Dep   DepCmd   `cmd:"" help:"Show dependencies of a WASM application"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func main() {
	cli := new(CLI)
	kong := kong.Parse(cli)

	// Additional context setup
	cli.Context.log = NewLogger(cli.Verbose)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	cli.Context.ctx = ctx

	// Run the selected command
	kong.FatalIfErrorf(kong.Run(&cli.Context))
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

/*
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
	// Resolve the path to get the directory name
	path, err := filepath.Abs(c.Path)
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

	// Create a compiler context from the configuration
	compile_context := config.CompileContext(path, CLI.Output, CLI.GoPath, CLI.GoFlags)

	/*
		// Compile application
		compile := &CompileCommand{
			Path:      path,
			BuildPath: CLI.Output,
			GoPath:    ctx.GoPath,
			GoFlags:   CLI.GoFlags,
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
*/
