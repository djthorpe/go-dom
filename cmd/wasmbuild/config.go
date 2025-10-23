package main

import (
	"io"
	"os"
	"path/filepath"

	// Packages
	yaml "gopkg.in/yaml.v3"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Config struct {
	Vars   map[string]string `yaml:"vars,omitempty"`
	Assets []string          `yaml:"assets,omitempty"`
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
