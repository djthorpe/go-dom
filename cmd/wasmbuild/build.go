package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type BuildPath struct {
	Path string `arg:"" default:"." help:"Path to the go application"`
}

type BuildCmd struct {
	BuildPath
	Output string `short:"o" help:"Output directory (uses temp dir if not specified)"`
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *BuildCmd) Run(ctx *Context) error {
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

	// Compile
	file, err := buildContext.CompileExec(ctx)
	if err != nil {
		return err
	}

	// Copy files to output directory
	for _, files := range []*File{
		file,
		buildContext.WasmExecHTML,
		buildContext.WasmExecJS,
		buildContext.FavIcon,
	} {
		// Write file
		ctx.log.Info("cp ", files.Path, " ", buildContext.Output)
		if err := files.WriteTo(buildContext.Output); err != nil {
			return fmt.Errorf("failed to copy %s: %w", files.Path, err)
		}
	}

	// Copy assets to output directory
	for _, asset := range config.Assets {
		if filepath.IsAbs(asset) == false {
			asset = filepath.Join(c.Path, asset)
		}
		dest := fmt.Sprintf("%s/%s", buildContext.Output, filepath.Base(asset))

		// Walk the asset path and copy files
		err := filepath.Walk(asset, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Relative path from asset root
			relPath, err := filepath.Rel(asset, path)
			if err != nil {
				return err
			}

			destPath := filepath.Join(dest, relPath)
			if info.Mode().IsDir() {
				ctx.log.Info("mkdir ", destPath)
				if err := os.MkdirAll(destPath, 0755); err != nil {
					return err
				}
				// Walk into directory
				return nil
			}

			// Copy file across
			ctx.log.Info("cp ", path, " ", filepath.Dir(destPath))
			if err := CopyFile(path, destPath); err != nil {
				return err
			}

			// Return success
			return nil
		})
		if err != nil {
			return fmt.Errorf("failed to copy asset %s: %w", asset, err)
		}
	}

	// Print out the destination to stdout
	fmt.Println(buildContext.Output)

	// Return success
	return nil
}

func (c *BuildContext) CompileExec(ctx *Context) (*File, error) {
	// Create temporary directory for build
	tmpDir, err := os.MkdirTemp("", "wasmbuild-compile-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	// Build to temporary location
	wasmPath := filepath.Join(tmpDir, filepath.Base(c.Path)+".wasm")
	cmd := c.GoBuildCmd("-o", wasmPath)
	// Update the output path to temp location
	for i, arg := range cmd.Args {
		if arg == "-o" && i+1 < len(cmd.Args) {
			cmd.Args[i+1] = wasmPath
			break
		}
	}

	// Log the compile
	ctx.log.Info(cmd.String())

	// Capture stderr for error messages
	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	// Run the command
	if err := cmd.Run(); err != nil {
		stderr := stderrBuf.String()
		return nil, fmt.Errorf("compilation failed: %w\n%s", err, stderr)
	}

	// Read the compiled wasm file into memory
	wasmData, err := os.ReadFile(wasmPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read compiled wasm file: %w", err)
	}

	// Return as File object
	return NewFile(wasmData, filepath.Base(c.Path)+".wasm"), nil
}

func CopyFile(src, dest string) error {
	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy data
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	// Return success
	return nil
}
