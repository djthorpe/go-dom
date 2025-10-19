package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// PackageInfo represents the JSON output from go list
type PackageInfo struct {
	ImportPath string   `json:"ImportPath"`
	Dir        string   `json:"Dir"`
	Deps       []string `json:"Deps"`
	Module     *struct {
		Path string `json:"Path"`
	} `json:"Module"`
}

// discoverLocalDependencies finds all local package dependencies for the given path
// and returns their absolute directory paths
func discoverLocalDependencies(goPath, sourcePath string) ([]string, error) {

	// Get package information including all dependencies
	cmd := exec.Command(goPath, "list", "-json", ".")
	cmd.Dir = sourcePath
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list package dependencies: %w", err)
	}

	var pkgInfo PackageInfo
	if err := json.Unmarshal(output, &pkgInfo); err != nil {
		return nil, fmt.Errorf("failed to parse package info: %w", err)
	}

	// Check if we're in a module
	if pkgInfo.Module == nil {
		logger.Infof("Not in a Go module, skipping dependency discovery")
		return nil, nil
	}

	modulePath := pkgInfo.Module.Path
	logger.Infof("Module path: %s", modulePath)
	logger.Infof("Found %d total dependencies for %s", len(pkgInfo.Deps), pkgInfo.ImportPath)

	// Find local dependencies (those that start with the module path)
	localDeps := make(map[string]bool) // Use map to deduplicate
	for _, dep := range pkgInfo.Deps {
		if strings.HasPrefix(dep, modulePath) && dep != pkgInfo.ImportPath {
			// Get the directory for this dependency
			depCmd := exec.Command(goPath, "list", "-json", dep)
			depCmd.Dir = sourcePath
			depOutput, err := depCmd.Output()
			if err != nil {
				logger.Infof("Could not get info for dependency %s: %v", dep, err)
				continue
			}

			var depInfo PackageInfo
			if err := json.Unmarshal(depOutput, &depInfo); err != nil {
				logger.Infof("Could not parse info for dependency %s: %v", dep, err)
				continue
			}

			if depInfo.Dir != "" {
				absDir, err := filepath.Abs(depInfo.Dir)
				if err == nil {
					localDeps[absDir] = true
				}
			}
		}
	}

	// Convert map to slice
	result := make([]string, 0, len(localDeps))
	for dir := range localDeps {
		result = append(result, dir)
	}

	return result, nil
}
