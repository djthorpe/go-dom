package bootstrap

import (
	"strings"
	"testing"

	. "github.com/djthorpe/go-wasmbuild"
	"github.com/stretchr/testify/assert"
)

func TestForm(t *testing.T) {
	assert := assert.New(t)

	form := Form()
	assert.NotNil(form)
	assert.NotNil(form.Element())
	assert.Equal("FORM", form.Element().TagName())

	// Check defaults
	assert.True(form.Element().HasAttribute("novalidate"))
	assert.True(form.Element().ClassList().Contains("needs-validation"))
}

func TestFormWithOptions(t *testing.T) {
	assert := assert.New(t)

	form := Form(WithClass("my-form"))
	classList := form.Element().ClassList()
	assert.True(classList.Contains("my-form"))
}

func TestFormAction(t *testing.T) {
	assert := assert.New(t)

	form := Form(WithAction("/submit"))
	assert.Equal("/submit", form.Element().GetAttribute("action"))
}

func TestFormMethod(t *testing.T) {
	assert := assert.New(t)

	form := Form(WithMethod("POST"))
	assert.Equal("POST", form.Element().GetAttribute("method"))
}

func TestFormEnctype(t *testing.T) {
	assert := assert.New(t)

	form := Form(WithEnctype("multipart/form-data"))
	assert.Equal("multipart/form-data", form.Element().GetAttribute("enctype"))
}

func TestFormNoValidate(t *testing.T) {
	assert := assert.New(t)

	// WithoutValidation() removes the novalidate attribute to enable browser validation
	form := Form(WithoutValidation())
	assert.False(form.Element().HasAttribute("novalidate"))
}

func TestFormDefaultValidation(t *testing.T) {
	assert := assert.New(t)

	// Default form has needs-validation class
	form := Form()
	assert.True(form.Element().ClassList().Contains("needs-validation"))
}

func TestFormChaining(t *testing.T) {
	assert := assert.New(t)

	form := Form(
		WithAction("/api/login"),
		WithMethod("POST"),
	)

	assert.Equal("/api/login", form.Element().GetAttribute("action"))
	assert.Equal("POST", form.Element().GetAttribute("method"))
	assert.True(form.Element().HasAttribute("novalidate"))
	assert.True(form.Element().ClassList().Contains("needs-validation"))
}

func TestFormAppend(t *testing.T) {
	assert := assert.New(t)

	form := Form()
	form.Append("Text content")

	assert.Equal(1, len(form.Element().ChildNodes()))
	assert.Equal("Text content", form.Element().TextContent())
}

func TestFormAppendComponents(t *testing.T) {
	assert := assert.New(t)

	form := Form()
	heading := Heading(2)
	heading.Append("Form Title")
	para := Para()
	para.Append("Form description")

	form.Append(heading, para)

	assert.Equal(2, len(form.Element().ChildNodes()))
}

func TestFormOnSubmit(t *testing.T) {
	assert := assert.New(t)

	form := Form().OnSubmit(func(e Event) {
		e.PreventDefault()
	})

	assert.NotNil(form)
	// Note: In non-WASM environment, the callback won't actually be triggered
	// This just verifies the OnSubmit method can be called without error
}

func TestFormOuterHTML(t *testing.T) {
	tests := []struct {
		name      string
		setupForm func() *form
		contains  []string // Check for these substrings instead of exact match
	}{
		{
			name:      "empty form",
			setupForm: func() *form { return Form() },
			contains:  []string{`class="needs-validation"`, `novalidate=`},
		},
		{
			name: "form with action and method",
			setupForm: func() *form {
				return Form(WithAction("/submit"), WithMethod("POST"))
			},
			contains: []string{`action="/submit"`, `method="post"`, `class="needs-validation"`, `novalidate=`},
		},
		{
			name: "form with class",
			setupForm: func() *form {
				return Form(WithClass("my-custom-class"))
			},
			contains: []string{`class="`, `my-custom-class`, `needs-validation`, `novalidate=`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := tt.setupForm()
			html := strings.ToLower(form.Element().OuterHTML())
			for _, substr := range tt.contains {
				assert.Contains(t, html, strings.ToLower(substr), "HTML should contain: %s", substr)
			}
		})
	}
}

func TestFormComponentInterface(t *testing.T) {
	form := Form()

	// Verify it implements the Component interface
	assert.NotNil(t, form.Element())
	assert.Equal(t, "container", form.Name())
}
