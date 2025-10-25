package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type File struct {
	Data []byte
	Path string
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewFile(data []byte, dest string) *File {
	return &File{
		Data: data,
		Path: dest,
	}
}

func NewFileFromSource(source, dest string) (*File, error) {
	data, err := os.ReadFile(source)
	if err != nil {
		return nil, err
	}
	return &File{
		Data: data,
		Path: dest,
	}, nil
}

func NewFileFromTemplate(data []byte, dest string, vars map[string]string, funcs template.FuncMap) (*File, error) {
	tmpl := template.New(filepath.Base(dest))

	// Define functions before parsing
	if funcs != nil {
		tmpl = tmpl.Funcs(funcs)
	}

	// Expand environment variables in template data
	for key, value := range vars {
		vars[key] = os.ExpandEnv(value)
	}

	// Parse the template
	tmpl, err := tmpl.Parse(string(data))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, vars); err != nil {
		return nil, err
	}
	return &File{
		Data: buf.Bytes(),
		Path: dest,
	}, nil
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (f *File) MarshalJSON() ([]byte, error) {
	return []byte(`"` + f.Path + `"`), nil
}

func (f *File) String() string {
	return f.Path
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (f *File) WriteTo(dir string) error {
	dest := filepath.Join(dir, f.Path)
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return os.WriteFile(dest, f.Data, 0o644)
}

func (f *File) URL() string {
	return "/" + f.Path
}

func (f *File) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set the Content-Type and Content-Length headers
		w.Header().Set("Content-Type", http.DetectContentType(f.Data))
		w.Header().Set("Content-Length", fmt.Sprintf("%d", len(f.Data)))

		// Set no-cache headers
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		// Output the data
		w.Write(f.Data)
	})
}
