package main

import (
	"context"
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type DepContext struct {
	BuildContext

	// Modified channel
	modified chan struct{}
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
	buildContext, err := config.BuildContext(ctx, c.Path, "")
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
				case <-dep.modified:
					fmt.Println("Modified")
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

// DepContext creates a DepContext from the Config, returning all the
// information needed to build a WASM application.
func (b BuildContext) DepContext(ctx *Context) (*DepContext, error) {
	// Return the DepContext
	return &DepContext{
		BuildContext: b,
		modified:     make(chan struct{}),
	}, nil
}

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

// Return a map of modification times on all files
func (d *DepContext) ModTimes(paths []string, cap int) map[string]time.Time {
	modTimes := make(map[string]time.Time, cap)
	for _, k := range paths {
		info, err := os.Stat(k)
		if err != nil {
			continue
		}
		if info.IsDir() {
			// Read files in the directory
			entries, err := os.ReadDir(k)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), ".") {
					continue
				}
				entryInfo, err := entry.Info()
				if err != nil {
					continue
				}
				if !entryInfo.IsDir() {
					modTimes[filepath.Join(k, entry.Name())] = entryInfo.ModTime()
				}
			}
		} else {
			modTimes[k] = info.ModTime()
		}
	}
	return modTimes
}

// Run a watcher for dependencies
func (d *DepContext) Run(ctx context.Context) error {
	// Get paths of dependencies
	paths, err := d.Dependencies()
	if err != nil {
		return err
	}

	// Track modification times
	modTimes := d.ModTimes(paths, len(paths))

	// Check for dependency changes every 5 seconds
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			// Check each dependency for changes
			modTimes2 := d.ModTimes(paths, len(modTimes))

			// Compare the two maps
			if !maps.Equal(modTimes2, modTimes) {
				// Redisover new paths
				if paths2, err := d.Dependencies(); err == nil {
					paths = paths2
				}

				// Send modification signal
				d.modified <- struct{}{}
			}

			// Set modTimes to modTimes2 for next iteration
			modTimes = modTimes2
		}
	}
}
