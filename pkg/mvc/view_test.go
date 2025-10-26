package mvc

import (
	"testing"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

func TestRegisterView(t *testing.T) {
	// Save original views map and restore it after test
	originalViews := views
	defer func() {
		views = originalViews
	}()

	// Reset views for testing
	views = make(map[string]ViewConstructorFunc, 50)

	tests := []struct {
		name        string
		viewName    string
		constructor ViewConstructorFunc
		shouldPanic bool
	}{
		{
			name:        "register new view",
			viewName:    "test-view",
			constructor: nil,
			shouldPanic: false,
		},
		{
			name:        "register duplicate view",
			viewName:    "test-view",
			constructor: nil,
			shouldPanic: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("RegisterView() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
			}()
			RegisterView(tt.viewName, tt.constructor)
		})
	}
}

func TestDataComponentAttrKey(t *testing.T) {
	expected := "data-wasmbuild-component"
	if DataComponentAttrKey != expected {
		t.Errorf("DataComponentAttrKey = %v, want %v", DataComponentAttrKey, expected)
	}
}

func TestNewView(t *testing.T) {
	// Save and restore original views
	originalViews := views
	defer func() {
		views = originalViews
	}()

	tests := []struct {
		name        string
		viewName    string
		tagName     string
		setup       func()
		shouldPanic bool
		panicMsg    string
	}{
		{
			name:     "create registered view",
			viewName: "test-view",
			tagName:  "DIV",
			setup: func() {
				views = make(map[string]ViewConstructorFunc, 50)
				RegisterView("test-view", nil)
			},
			shouldPanic: false,
		},
		{
			name:     "create unregistered view",
			viewName: "unregistered-view",
			tagName:  "DIV",
			setup: func() {
				views = make(map[string]ViewConstructorFunc, 50)
			},
			shouldPanic: true,
			panicMsg:    "View not registered",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}

			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("NewView() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
				if r != nil && tt.panicMsg != "" {
					// Check panic message contains expected text
					if msg, ok := r.(string); ok {
						if len(msg) < len(tt.panicMsg) || msg[:len(tt.panicMsg)] != tt.panicMsg {
							t.Errorf("NewView() panic message = %v, want to contain %v", msg, tt.panicMsg)
						}
					}
				}
			}()

			v := NewView(tt.viewName, tt.tagName)
			if !tt.shouldPanic {
				if v == nil {
					t.Fatal("NewView() returned nil")
				}
				if v.Name() != tt.viewName {
					t.Errorf("NewView().Name() = %v, want %v", v.Name(), tt.viewName)
				}
			}
		})
	}
}

func TestViewName(t *testing.T) {
	v := &view{name: "test-component"}
	if got := v.Name(); got != "test-component" {
		t.Errorf("Name() = %v, want %v", got, "test-component")
	}
}

func TestViewString(t *testing.T) {
	// This test verifies the String() method exists
	// We can't test actual HTML output without a real DOM environment
	t.Skip("String() requires DOM environment - tested in integration tests")
}

func TestNodeFromAny(t *testing.T) {
	tests := []struct {
		name        string
		input       any
		shouldPanic bool
		wantType    string
	}{
		{
			name:        "unsupported type - int",
			input:       123,
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - bool",
			input:       true,
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - slice",
			input:       []string{"a", "b"},
			shouldPanic: true,
			wantType:    "",
		},
		{
			name:        "unsupported type - map",
			input:       map[string]string{"key": "value"},
			shouldPanic: true,
			wantType:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) != tt.shouldPanic {
					t.Errorf("NodeFromAny() panic = %v, wantPanic %v", r != nil, tt.shouldPanic)
				}
			}()

			_ = NodeFromAny(tt.input)
		})
	}
}

func TestViewChaining(t *testing.T) {
	// Test that methods return the view for chaining (compile-time check)
	t.Run("method chaining pattern", func(t *testing.T) {
		// This is a compile-time check that the methods have the right signatures
		// for method chaining. We can't actually test them without a DOM.

		// These should all compile - they return View for chaining
		_ = func(v View) View { return v.Empty() }
		_ = func(v View) View { return v.Insert("test") }
		_ = func(v View) View { return v.Append("test") }
		_ = func(v View) View { return v.AddEventListener("click", nil) }
		_ = func(v View) View { return v.Apply() }

		// Body doesn't return View, so it breaks the chain intentionally
		_ = func(v View) { v.Body(nil) }
	})
}

func TestViewInterfaceImplementation(t *testing.T) {
	// Compile-time check that view implements View interface
	var _ View = (*view)(nil)

	// Runtime check
	v := &view{name: "test"}
	if _, ok := any(v).(View); !ok {
		t.Error("view does not implement View interface")
	}
}

func TestViewConstructorFunc(t *testing.T) {
	// Test that ViewConstructorFunc type works as expected
	var constructor ViewConstructorFunc = func(e Element) View {
		return &view{name: "constructed"}
	}

	if constructor == nil {
		t.Error("ViewConstructorFunc should not be nil")
	}
}

func TestViewStructFields(t *testing.T) {
	v := &view{
		name: "test-view",
		root: nil,
		body: nil,
	}

	if v.name != "test-view" {
		t.Errorf("view.name = %v, want %v", v.name, "test-view")
	}

	if v.root != nil {
		t.Errorf("view.root should be nil")
	}

	if v.body != nil {
		t.Errorf("view.body should be nil")
	}
}
