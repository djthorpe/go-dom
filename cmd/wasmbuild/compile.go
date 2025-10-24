package main

import (
	"bytes"
	"fmt"
	"os"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type CompilePath struct {
	Path string `arg:"" default:"." help:"Path to the go application"`
}

type CompileCmd struct {
	CompilePath
	Output string `short:"o" help:"Output directory (uses temp dir if not specified)"`
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *CompileCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveDir(ctx.Config, c.Path, true)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create a compiler context from the configuration
	buildContext, err := config.BuildContext(ctx, c.Path, c.Output)
	if err != nil {
		return err
	}

	// Prepare the command for exec
	cmd := buildContext.GoBuildCmd()

	// Log the compile
	ctx.log.Info("Compile: ", buildContext)

	// Capture stderr for error messages
	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	// Run the command
	if err := cmd.Run(); err != nil {
		stderr := stderrBuf.String()
		return fmt.Errorf("compilation failed: %w\n%s", err, stderr)
	}

	// Copy files to output directory
	for _, files := range []*File{
		buildContext.WasmExecHTML,
		buildContext.WasmExecJS,
		buildContext.FavIcon,
	} {
		// Write file
		if err := files.WriteTo(buildContext.Output); err != nil {
			return fmt.Errorf("failed to copy %s: %w", files.Path, err)
		}
	}

	// Return success
	return nil
}
