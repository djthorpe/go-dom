package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"

	// Packages
	"github.com/fsnotify/fsnotify"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type DepContext struct {
	BuildContext

	// Modified channel - returns nil or an error
	modified chan error

	// The WebAssembly file that was compiled
	wasm *File
}

type DepCmd struct {
	BuildPath
	WatchFlag
}

// DepPackageInfo represents the JSON output from go list
type DepPackageInfo struct {
	ImportPath string   `json:"ImportPath"`
	Dir        string   `json:"Dir"`
	Deps       []string `json:"Deps"`
	Module     *struct {
		Path string `json:"Path"`
	} `json:"Module"`
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// DepContext creates a DepContext from the Config, returning all the
// information needed to build a WASM application.
func (b BuildContext) DepContext(ctx *Context) (*DepContext, error) {
	// Return the DepContext
	return &DepContext{
		BuildContext: b,
		modified:     make(chan error),
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (d *DepContext) String() string {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *DepCmd) Run(ctx *Context) error {
	// Read the configuration file
	configPath, err := ResolveFile(ctx.Config, c.Path)
	if err != nil {
		return err
	}
	config, err := ParseYAMLPath(configPath, c.Path)
	if err != nil {
		return err
	}

	// Create a build context from the configuration
	buildContext, err := config.BuildContext(ctx, c.Path, "", false)
	if err != nil {
		return err
	}

	// Create a dependency context from the build context
	dep, err := buildContext.DepContext(ctx)
	if err != nil {
		return err
	} else {
		ctx.log.Info(dep)
	}

	// If watch flag is set, run a watcher
	if c.Watch {
		// Watch for dependency changes
		var wg sync.WaitGroup
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
				case event := <-dep.modified:
					if event != nil {
						ctx.log.Error(event)
						continue
					}

					// Compile the code
					if wasm, err := buildContext.CompileExec(ctx); err != nil {
						ctx.log.Error("Compilation error after modification: ", err)
						continue
					} else {
						dep.wasm = wasm
					}

					// Indicate success
					ctx.log.Info("Re-compiled wasm: ", dep.wasm.Path)
				}
			}
		}()

		// Wait for all go-routines to end
		wg.Wait()
	} else if paths, err := dep.Dependencies(); err != nil {
		return err
	} else {
		// Print out dependencies
		for _, p := range paths {
			fmt.Println(p)
		}
	}

	// Return success
	return nil
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// Return a list of dependencies
func (d *DepContext) Dependencies() ([]string, error) {
	// Get package information including all dependencies
	cmd := exec.Command(d.GoCmd, "list", "-json", d.Path)

	var stderr strings.Builder
	cmd.Stderr = &stderr
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list package dependencies: %w", err)
	}

	var pkgInfo DepPackageInfo
	if err := json.Unmarshal(output, &pkgInfo); err != nil {
		return nil, fmt.Errorf("failed to parse package info: %w", err)
	}

	// Check if we're in a module
	if pkgInfo.Module == nil {
		return nil, fmt.Errorf("not in a Go module, skipping dependency discovery")
	}

	// Find local dependencies (those that start with the module path)
	deps := make(map[string]bool)
	for _, dep := range pkgInfo.Deps {
		if strings.HasPrefix(dep, pkgInfo.Module.Path) && dep != pkgInfo.ImportPath {
			// Get the directory for this dependency
			depCmd := exec.Command(d.GoCmd, "list", "-json", dep)
			depOutput, err := depCmd.Output()
			if err != nil {
				return nil, fmt.Errorf("could not get info for dependency %s: %v", dep, err)
			}

			var depInfo DepPackageInfo
			if err := json.Unmarshal(depOutput, &depInfo); err != nil {
				return nil, fmt.Errorf("could not parse info for dependency %s: %v", dep, err)
			}

			if depInfo.Dir != "" {
				absDir, err := filepath.Abs(depInfo.Dir)
				if err == nil {
					deps[absDir] = true
				}
			}
		}
	}

	// Append the input path as a dependency
	absPath, err := filepath.Abs(d.Path)
	if err == nil {
		deps[absPath] = true
	}

	// Append assets as dependencies
	for _, asset := range d.Assets {
		if filepath.IsAbs(asset) == false {
			asset = filepath.Join(d.Path, asset)
		}
		absAsset, err := filepath.Abs(asset)
		if err != nil {
			continue
		}

		// Asset files are directly added, directories are recursively walked
		info, err := os.Stat(absAsset)
		if err != nil {
			continue
		}
		if info.Mode().IsDir() {
			err := filepath.Walk(absAsset, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.HasPrefix(info.Name(), ".") {
					return nil
				}
				if info.Mode().IsDir() {
					absPath, err := filepath.Abs(path)
					if err == nil {
						deps[absPath] = true
					}
				}
				return nil
			})
			if err != nil {
				continue
			}
		} else {
			deps[absAsset] = true
		}
	}

	// Return the paths
	paths := make([]string, 0, len(deps))
	for dep := range deps {
		paths = append(paths, dep)
	}
	return paths, nil
}

// Run a watcher for dependencies using fsnotify
func (d *DepContext) Run(ctx context.Context) error {
	// Get paths of dependencies
	paths, err := d.Dependencies()
	if err != nil {
		return err
	}

	// Create fsnotify watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("failed to create watcher: %w", err)
	}
	defer watcher.Close()

	// Add all dependency paths to the watcher
	for _, path := range paths {
		if err := watcher.Add(path); err != nil {
			return fmt.Errorf("failed to watch %s: %w", path, err)
		}
	}

	// Track modification times for debouncing
	last := time.Now()
	debounceDelay := 500 * time.Millisecond

	for {
		select {
		case <-ctx.Done():
			return nil

		case event := <-watcher.Events:
			// Filter out events we don't care about
			if event.Has(fsnotify.Chmod) {
				continue
			}

			// Debounce: ignore events that occur too quickly
			now := time.Now()
			if now.Sub(last) < debounceDelay {
				continue
			} else {
				last = now
				d.modified <- nil
			}

			// TODO: If the event is a Create event on a directory, re-compute watche paths

		case err := <-watcher.Errors:
			d.modified <- fmt.Errorf("watcher error: %w", err)
		}
	}
}
