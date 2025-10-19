package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type CompileCommand struct {
	Path     string
	DirName  string
	BuildDir string
	GoPath   string
	GoFlags  string
}

type CompileResult struct {
	BuildDir string
	Error    error
	Stderr   string
}

func (c *CompileCommand) Run() error {
	result := c.RunAndGetResult()
	return result.Error
}

func (c *CompileCommand) RunAndGetBuildDir() (string, error) {
	result := c.RunAndGetResult()
	return result.BuildDir, result.Error
}

func (c *CompileCommand) RunAndGetResult() CompileResult {
	// Determine build directory
	buildDir := c.BuildDir
	if buildDir == "" {
		// Use temporary directory
		tmpDir, err := os.MkdirTemp("", "wasm-build-*")
		if err != nil {
			return CompileResult{Error: fmt.Errorf("failed to create temp directory: %w", err)}
		}
		buildDir = tmpDir
		logger.Infof("Using temporary build directory: %s", buildDir)
	} else {
		// Ensure build directory exists
		if err := os.MkdirAll(buildDir, 0755); err != nil {
			return CompileResult{Error: fmt.Errorf("failed to create build directory: %w", err)}
		}
		logger.Infof("Using build directory: %s", buildDir)
	}

	// Use the provided directory name for the output filename
	outputFile := filepath.Join(buildDir, c.DirName+".wasm")

	// Build the command: GOOS=js GOARCH=wasm go build -o output.wasm -tags js [goflags] ./path
	cmd := exec.Command(c.GoPath, "build", "-o", outputFile, "-tags", "js")

	// Add additional go flags if specified
	if c.GoFlags != "" {
		flags := strings.Fields(c.GoFlags)
		// Insert flags before the path argument
		cmd.Args = append(cmd.Args, flags...)
	}

	// Add the path to build - ensure it's treated as a file path, not a package path
	buildPath := c.Path
	if !filepath.IsAbs(buildPath) && !strings.HasPrefix(buildPath, "./") && !strings.HasPrefix(buildPath, "../") {
		buildPath = "./" + buildPath
	}
	cmd.Args = append(cmd.Args, buildPath)

	// Set environment variables
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "GOOS=js", "GOARCH=wasm")

	// Capture stderr for error messages
	var stderrBuf bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderrBuf

	logger.Infof("Executing: GOOS=js GOARCH=wasm %s", strings.Join(cmd.Args, " "))

	// Run the command
	if err := cmd.Run(); err != nil {
		stderr := stderrBuf.String()
		// Also print to stderr for console visibility
		fmt.Fprint(os.Stderr, stderr)
		return CompileResult{
			BuildDir: buildDir,
			Error:    fmt.Errorf("compilation failed: %w", err),
			Stderr:   stderr,
		}
	}

	logger.Infof("Successfully compiled to: %s", outputFile)
	return CompileResult{BuildDir: buildDir}
}
