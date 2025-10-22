package bootstrap

import (
	"strings"
	"testing"

	. "github.com/djthorpe/go-wasmbuild"
	"github.com/stretchr/testify/assert"
)

func TestSelect(t *testing.T) {
	assert := assert.New(t)

	sel := Select()
	assert.NotNil(sel)
	assert.NotNil(sel.Element())
	assert.Equal("SELECT", sel.Element().TagName())
	assert.True(sel.Element().ClassList().Contains("form-select"))
}

func TestSelectWithOptions(t *testing.T) {
	assert := assert.New(t)

	sel := Select(WithClass("my-select"))
	classList := sel.Element().ClassList()
	assert.True(classList.Contains("form-select"))
	assert.True(classList.Contains("my-select"))
}

func TestSelectWithName(t *testing.T) {
	assert := assert.New(t)

	sel := Select(WithName("country"))
	assert.Equal("country", sel.Element().GetAttribute("name"))
}

func TestSelectWithRequired(t *testing.T) {
	assert := assert.New(t)

	sel := Select(WithRequired())
	assert.True(sel.Element().HasAttribute("required"))
}

func TestSelectWithDisabled(t *testing.T) {
	assert := assert.New(t)

	sel := Select(WithDisabled())
	assert.True(sel.Element().HasAttribute("disabled"))
}

func TestSelectWithMultiple(t *testing.T) {
	assert := assert.New(t)

	sel := Select(WithMultiple())
	assert.True(sel.Element().HasAttribute("multiple"))
}

func TestSelectOnChange(t *testing.T) {
	assert := assert.New(t)

	sel := Select().OnChange(func(e Event) {
		// Handler
	})

	assert.NotNil(sel)
}

func TestSelectAppendOptions(t *testing.T) {
	assert := assert.New(t)

	sel := Select()
	opt1 := Option(WithValue("1")).Append("Option 1")
	opt2 := Option(WithValue("2")).Append("Option 2")

	sel.Append(opt1, opt2)

	assert.Equal(2, len(sel.Element().ChildNodes()))
}

func TestOption(t *testing.T) {
	assert := assert.New(t)

	opt := Option()
	assert.NotNil(opt)
	assert.NotNil(opt.Element())
	assert.Equal("OPTION", opt.Element().TagName())
}

func TestOptionWithValue(t *testing.T) {
	assert := assert.New(t)

	opt := Option(WithValue("test-value"))
	assert.Equal("test-value", opt.Element().GetAttribute("value"))
}

func TestOptionWithSelected(t *testing.T) {
	assert := assert.New(t)

	opt := Option(WithSelected())
	assert.True(opt.Element().HasAttribute("selected"))
}

func TestOptionWithDisabled(t *testing.T) {
	assert := assert.New(t)

	opt := Option(WithDisabled())
	assert.True(opt.Element().HasAttribute("disabled"))
}

func TestOptionWithContent(t *testing.T) {
	assert := assert.New(t)

	opt := Option(WithValue("1")).Append("First Option")
	assert.Equal("First Option", opt.Element().TextContent())
}

func TestSelectChaining(t *testing.T) {
	assert := assert.New(t)

	sel := Select(
		WithName("size"),
		WithRequired(),
		WithClass("form-select-lg"),
	)

	assert.Equal("size", sel.Element().GetAttribute("name"))
	assert.True(sel.Element().HasAttribute("required"))
	assert.True(sel.Element().ClassList().Contains("form-select"))
	assert.True(sel.Element().ClassList().Contains("form-select-lg"))
}

func TestSelectValue(t *testing.T) {
	assert := assert.New(t)

	sel := Select()
	sel.SetValue("test")
	assert.Equal("test", sel.Value())
}

func TestSelectOuterHTML(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *selectElement
		contains []string
	}{
		{
			name:     "empty select",
			setup:    func() *selectElement { return Select() },
			contains: []string{`<select`, `class="form-select"`, `</select>`},
		},
		{
			name: "select with name",
			setup: func() *selectElement {
				return Select(WithName("country"))
			},
			contains: []string{`<select`, `name="country"`, `class="form-select"`},
		},
		{
			name: "select with options",
			setup: func() *selectElement {
				sel := Select()
				sel.Append(
					Option(WithValue("1")).Append("One"),
					Option(WithValue("2")).Append("Two"),
				)
				return sel
			},
			contains: []string{`<select`, `<option`, `value="1"`, `>one<`, `value="2"`, `>two<`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sel := tt.setup()
			html := strings.ToLower(sel.Element().OuterHTML())
			for _, substr := range tt.contains {
				assert.Contains(t, html, strings.ToLower(substr), "HTML should contain: %s", substr)
			}
		})
	}
}

func TestOptionOuterHTML(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *option
		contains []string
	}{
		{
			name:     "empty option",
			setup:    func() *option { return Option() },
			contains: []string{`<option`, `</option>`},
		},
		{
			name: "option with value",
			setup: func() *option {
				return Option(WithValue("test"))
			},
			contains: []string{`<option`, `value="test"`},
		},
		{
			name: "option with selected",
			setup: func() *option {
				opt := Option(WithValue("1"), WithSelected())
				opt.Append("Selected")
				return opt
			},
			contains: []string{`<option`, `value="1"`, `selected`, `>selected<`},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opt := tt.setup()
			html := strings.ToLower(opt.Element().OuterHTML())
			for _, substr := range tt.contains {
				assert.Contains(t, html, strings.ToLower(substr), "HTML should contain: %s", substr)
			}
		})
	}
}

func TestSelectComponentInterface(t *testing.T) {
	sel := Select()

	// Verify it implements the Component interface
	assert.NotNil(t, sel.Element())
	assert.Equal(t, "select", sel.Name())
}

func TestOptionComponentInterface(t *testing.T) {
	opt := Option()

	// Verify it implements the Component interface
	assert.NotNil(t, opt.Element())
	assert.Equal(t, "option", opt.Name())
}
