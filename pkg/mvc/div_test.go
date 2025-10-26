package mvc

import (
	"testing"
)

func TestViewDivConstant(t *testing.T) {
	expected := "wasmbuild-mvc-div"
	if ViewDiv != expected {
		t.Errorf("ViewDiv = %v, want %v", ViewDiv, expected)
	}
}

func TestDivRegistered(t *testing.T) {
	// Check if the div view is registered during init
	if _, exists := views[ViewDiv]; !exists {
		t.Errorf("ViewDiv should be registered in init(), but was not found")
	}
}

// Note: Full testing of Div() requires a WASM environment with DOM access
// These tests verify the basic structure and registration only
func TestDivType(t *testing.T) {
	// Verify the div type exists
	var _ *div = &div{}
}
