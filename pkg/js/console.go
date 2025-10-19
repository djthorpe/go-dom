//go:build js

package js

import (
	"syscall/js"
)

// Console provides access to the browser's console API.
type Console struct {
	value js.Value
}

// GetConsole returns a Console instance for the global console object.
func GetConsole() *Console {
	return &Console{
		value: js.Global().Get("console"),
	}
}

// Log writes a message to the console.
func (c *Console) Log(args ...any) {
	c.value.Call("log", args...)
}

// Warn writes a warning message to the console.
func (c *Console) Warn(args ...any) {
	c.value.Call("warn", args...)
}

// Error writes an error message to the console.
func (c *Console) Error(args ...any) {
	c.value.Call("error", args...)
}

// Info writes an informational message to the console.
func (c *Console) Info(args ...any) {
	c.value.Call("info", args...)
}

// Debug writes a debug message to the console.
func (c *Console) Debug(args ...any) {
	c.value.Call("debug", args...)
}

// Clear clears the console.
func (c *Console) Clear() {
	c.value.Call("clear")
}
