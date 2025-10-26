package mvc

import (
	"testing"
)

func TestWithClass(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		add      []string
		expected []string
	}{
		{
			name:     "add single class",
			initial:  []string{},
			add:      []string{"btn"},
			expected: []string{"btn"},
		},
		{
			name:     "add multiple classes",
			initial:  []string{},
			add:      []string{"btn", "btn-primary"},
			expected: []string{"btn", "btn-primary"},
		},
		{
			name:     "add to existing classes",
			initial:  []string{"container"},
			add:      []string{"mx-auto", "mt-3"},
			expected: []string{"container", "mx-auto", "mt-3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{class: tt.initial}
			optFunc := WithClass(tt.add...)
			if err := optFunc(o); err != nil {
				t.Fatalf("WithClass() error = %v", err)
			}
			if len(o.class) != len(tt.expected) {
				t.Errorf("WithClass() got %d classes, want %d", len(o.class), len(tt.expected))
			}
			for i, class := range tt.expected {
				if i >= len(o.class) || o.class[i] != class {
					t.Errorf("WithClass() class[%d] = %v, want %v", i, o.class, tt.expected)
					break
				}
			}
		})
	}
}

func TestWithoutClass(t *testing.T) {
	tests := []struct {
		name     string
		initial  []string
		remove   []string
		expected []string
	}{
		{
			name:     "remove single class",
			initial:  []string{"btn", "btn-primary"},
			remove:   []string{"btn-primary"},
			expected: []string{"btn"},
		},
		{
			name:     "remove multiple classes",
			initial:  []string{"btn", "btn-primary", "active"},
			remove:   []string{"btn-primary", "active"},
			expected: []string{"btn"},
		},
		{
			name:     "remove non-existent class",
			initial:  []string{"btn"},
			remove:   []string{"btn-danger"},
			expected: []string{"btn"},
		},
		{
			name:     "remove all classes",
			initial:  []string{"btn", "active"},
			remove:   []string{"btn", "active"},
			expected: []string{},
		},
		{
			name:     "remove duplicate classes",
			initial:  []string{"btn", "btn", "active"},
			remove:   []string{"btn"},
			expected: []string{"active"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{class: tt.initial}
			optFunc := WithoutClass(tt.remove...)
			if err := optFunc(o); err != nil {
				t.Fatalf("WithoutClass() error = %v", err)
			}
			if len(o.class) != len(tt.expected) {
				t.Errorf("WithoutClass() got %d classes, want %d: got %v, want %v", len(o.class), len(tt.expected), o.class, tt.expected)
			}
			for i, class := range tt.expected {
				if i >= len(o.class) || o.class[i] != class {
					t.Errorf("WithoutClass() class[%d] = %v, want %v", i, o.class, tt.expected)
					break
				}
			}
		})
	}
}

func TestWithAttr(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]string
		key      string
		value    string
		expected map[string]string
	}{
		{
			name:     "add single attribute",
			initial:  nil,
			key:      "href",
			value:    "/home",
			expected: map[string]string{"href": "/home"},
		},
		{
			name:     "add to existing attributes",
			initial:  map[string]string{"class": "btn"},
			key:      "type",
			value:    "button",
			expected: map[string]string{"class": "btn", "type": "button"},
		},
		{
			name:     "overwrite existing attribute",
			initial:  map[string]string{"href": "/old"},
			key:      "href",
			value:    "/new",
			expected: map[string]string{"href": "/new"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{attr: tt.initial}
			optFunc := WithAttr(tt.key, tt.value)
			if err := optFunc(o); err != nil {
				t.Fatalf("WithAttr() error = %v", err)
			}
			if len(o.attr) != len(tt.expected) {
				t.Errorf("WithAttr() got %d attributes, want %d", len(o.attr), len(tt.expected))
			}
			for key, value := range tt.expected {
				if o.attr[key] != value {
					t.Errorf("WithAttr() attr[%s] = %v, want %v", key, o.attr[key], value)
				}
			}
		})
	}
}

func TestWithoutAttr(t *testing.T) {
	tests := []struct {
		name     string
		initial  map[string]string
		keys     []string
		expected map[string]string
	}{
		{
			name:     "remove single attribute",
			initial:  map[string]string{"href": "/home", "target": "_blank"},
			keys:     []string{"target"},
			expected: map[string]string{"href": "/home"},
		},
		{
			name:     "remove multiple attributes",
			initial:  map[string]string{"href": "/home", "target": "_blank", "rel": "noopener"},
			keys:     []string{"target", "rel"},
			expected: map[string]string{"href": "/home"},
		},
		{
			name:     "remove non-existent attribute",
			initial:  map[string]string{"href": "/home"},
			keys:     []string{"target"},
			expected: map[string]string{"href": "/home"},
		},
		{
			name:     "remove from nil map",
			initial:  nil,
			keys:     []string{"href"},
			expected: nil,
		},
		{
			name:     "remove all attributes",
			initial:  map[string]string{"href": "/home"},
			keys:     []string{"href"},
			expected: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{attr: tt.initial}
			optFunc := WithoutAttr(tt.keys...)
			if err := optFunc(o); err != nil {
				t.Fatalf("WithoutAttr() error = %v", err)
			}
			if tt.expected == nil && o.attr == nil {
				return
			}
			if len(o.attr) != len(tt.expected) {
				t.Errorf("WithoutAttr() got %d attributes, want %d", len(o.attr), len(tt.expected))
			}
			for key, value := range tt.expected {
				if o.attr[key] != value {
					t.Errorf("WithoutAttr() attr[%s] = %v, want %v", key, o.attr[key], value)
				}
			}
		})
	}
}

func TestWithID(t *testing.T) {
	tests := []struct {
		name     string
		id       string
		expected string
	}{
		{
			name:     "set simple id",
			id:       "my-element",
			expected: "my-element",
		},
		{
			name:     "set empty id",
			id:       "",
			expected: "",
		},
		{
			name:     "set complex id",
			id:       "nav-item-123",
			expected: "nav-item-123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{}
			optFunc := WithID(tt.id)
			if err := optFunc(o); err != nil {
				t.Fatalf("WithID() error = %v", err)
			}
			if o.id != tt.expected {
				t.Errorf("WithID() id = %v, want %v", o.id, tt.expected)
			}
		})
	}
}

func TestOptClasses(t *testing.T) {
	tests := []struct {
		name     string
		classes  []string
		expected []string
	}{
		{
			name:     "no duplicates",
			classes:  []string{"btn", "btn-primary"},
			expected: []string{"btn", "btn-primary"},
		},
		{
			name:     "with duplicates",
			classes:  []string{"btn", "btn", "active"},
			expected: []string{"btn", "active"},
		},
		{
			name:     "all duplicates",
			classes:  []string{"btn", "btn", "btn"},
			expected: []string{"btn"},
		},
		{
			name:     "empty list",
			classes:  []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &opt{class: tt.classes}
			result := o.Classes()
			if len(result) != len(tt.expected) {
				t.Errorf("Classes() got %d classes, want %d: got %v, want %v", len(result), len(tt.expected), result, tt.expected)
			}
			for i, class := range tt.expected {
				if i >= len(result) || result[i] != class {
					t.Errorf("classes() class[%d] = %v, want %v", i, result, tt.expected)
					break
				}
			}
		})
	}
}
