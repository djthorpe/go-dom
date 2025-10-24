package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	// Packages
	etc "github.com/djthorpe/go-wasmbuild/etc"
	yaml "gopkg.in/yaml.v3"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Config struct {
	Vars   map[string]string `yaml:"vars,omitempty" json:"vars,omitempty"`
	Assets []string          `yaml:"assets,omitempty" json:"assets,omitempty"`
}

type BuildContext struct {
	Config

	// Source path for build
	Path string `json:"input,omitempty"`

	// Output path for build
	Output string `json:"output,omitempty"`

	// Go tool command and flags
	GoCmd  string   `json:"go_cmd,omitempty"`
	GoRoot string   `json:"go_root,omitempty"`
	GoArgs []string `json:"go_args,omitempty"`
	GoEnv  []string `json:"go_env,omitempty"`

	// WasmExec Javascript path
	WasmExecJS   *File `json:"wasm_exec_js,omitempty"`
	WasmExecHTML *File `json:"wasm_exec_html,omitempty"`
	FavIcon      *File `json:"favicon,omitempty"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// ParseYAML parses YAML configuration from an io.Reader
func ParseYAML(r io.Reader) (*Config, error) {
	var config Config
	if err := yaml.NewDecoder(r).Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ParseYAMLPath parses a YAML configuration from a path (relative to a base path)
func ParseYAMLPath(path, base string) (*Config, error) {
	if filepath.IsAbs(path) == false {
		path = filepath.Join(base, path)
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ParseYAML(f)
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *Config) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (c *BuildContext) String() string {
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// BuildContext creates a BuildContext from the Config, returning all the
// information needed to build a WASM application.
func (c Config) BuildContext(ctx *Context, path, output string) (*BuildContext, error) {
	// Make input path absolute
	if filepath.IsAbs(path) == false {
		var err error
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, fmt.Errorf("failed to determine absolute path: %w", err)
		}
	}

	// Determine the output path
	if output == "" {
		// Use temporary directory
		var err error
		output, err = os.MkdirTemp("", fmt.Sprintf("wasmbuild-%s-*", filepath.Base(path)))
		if err != nil {
			return nil, fmt.Errorf("failed to create temp directory: %w", err)
		}
	} else {
		// Make output path absolute
		if filepath.IsAbs(output) == false {
			var err error
			output, err = filepath.Abs(output)
			if err != nil {
				return nil, fmt.Errorf("failed to determine absolute path: %w", err)
			}
		}

		// Ensure build directory exists
		if err := os.MkdirAll(output, 0755); err != nil {
			return nil, fmt.Errorf("failed to create build directory: %w", err)
		}
	}

	// Get GOROOT from environment
	goroot := os.Getenv("GOROOT")
	if goroot == "" {
		// If GOROOT is not set, try to determine it from the go tool
		if !filepath.IsAbs(ctx.Go) {
			var err error
			ctx.Go, err = exec.LookPath(ctx.Go)
			if err != nil {
				return nil, fmt.Errorf("failed to locate go executable: %w", err)
			}
		}

		// Run 'go env GOROOT' to get GOROOT
		cmd := exec.Command(ctx.Go, "env", "GOROOT")
		output, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("failed to determine GOROOT: %w", err)
		}
		goroot = strings.TrimSpace(string(output))
	}

	// wasm_exec.js
	wasmPathExecJS := RegularFileFromPathList(ctx.WasmExec, goroot)
	if wasmPathExecJS == "" {
		return nil, fmt.Errorf("wasm_exec.js not found in GOROOT")
	}
	wasmExecJS, err := NewFileFromSource(wasmPathExecJS, "wasm_exec.js")
	if err != nil {
		return nil, fmt.Errorf("failed to read wasm_exec.js: %w", err)
	}

	//  wasm_exec.html
	funcs := template.FuncMap{
		"Title": func() string {
			if title, ok := c.Vars["Title"]; ok {
				return title
			}
			return filepath.Base(path)
		},
		"Header": func() string {
			if head, ok := c.Vars["Header"]; ok {
				return head
			}
			return ""
		},
		"Footer": func() string {
			if foot, ok := c.Vars["Footer"]; ok {
				return foot
			}
			return ""
		},
		"Notify": func() string {
			if notify, ok := c.Vars["Notify"]; ok {
				return notify
			}
			return ""
		},
		"WasmFile": func() string {
			if wasmFile, ok := c.Vars["WasmFile"]; ok {
				return wasmFile
			}
			return filepath.Base(path) + ".wasm"
		},
	}
	wasmExecHTML, err := NewFileFromTemplate(etc.WasmExecHTML, "wasm_exec.html", c.Vars, funcs)
	if err != nil {
		return nil, fmt.Errorf("failed to create wasm_exec.html: %w", err)
	}

	// Return build context
	return &BuildContext{
		Config: c,
		Path:   path,
		Output: output,
		GoCmd:  ctx.Go,
		GoRoot: goroot,
		GoArgs: append([]string{
			"build",
		}, strings.Fields(ctx.GoFlags)...),
		GoEnv: []string{
			"GOOS=js",
			"GOARCH=wasm",
		},
		WasmExecJS:   wasmExecJS,
		WasmExecHTML: wasmExecHTML,
		FavIcon:      NewFile(etc.FaviconPNG, "favicon.png"),
	}, nil
}

// Return a exec.Cmd for building the WASM application
func (bc *BuildContext) GoBuildCmd(args ...string) *exec.Cmd {
	cmd := exec.Command(bc.GoCmd, append(bc.GoArgs, args...)...)
	cmd.Dir = bc.Path
	cmd.Env = append(os.Environ(), bc.GoEnv...)
	return cmd
}

// Return the path to file wasm_exec.js
func RegularFileFromPathList(path, base string) string {
	for _, path := range strings.Split(path, string(filepath.ListSeparator)) {
		if filepath.IsAbs(path) == false {
			path = filepath.Join(base, path)
		}
		if info, err := os.Stat(path); err == nil && info.Mode().IsRegular() {
			return path
		}
	}
	return ""
}
