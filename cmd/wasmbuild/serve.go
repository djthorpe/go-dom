package main

import "fmt"

///////////////////////////////////////////////////////////////////////////////
// TYPES

type ServeCmd struct {
	CompilePath
	Listen string `default:"localhost:9090" help:"Address to listen on (e.g., localhost:9090 or 0.0.0.0:9090)"`
	Watch  bool   `short:"w" help:"Watch for changes and recompile automatically"`
}

///////////////////////////////////////////////////////////////////////////////
// COMMANDS

func (c *ServeCmd) Run(ctx *Context) error {
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
	buildContext, err := config.BuildContext(ctx, c.Path, "")
	if err != nil {
		return err
	}

	// Log the compile
	ctx.log.Info("Compile: ", buildContext)

	fmt.Println("Starting server on", c.Listen)
	return nil
}
