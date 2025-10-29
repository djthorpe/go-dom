package mvc_test

import (
	"testing"

	// Packages
	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TestSpanType(t *testing.T) {
	assert := assert.New(t)

	span := Span()
	assert.NotNil(span)

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(ViewSpan, span.Name())
		assert.NotNil(span.Root())
		assert.Equal("SPAN", span.Root().TagName())
	})

	t.Run("NewViewWithElement", func(t *testing.T) {
		span2 := NewViewWithElement(span.Root())
		assert.NotNil(span2)
		assert.Equal(ViewSpan, span2.Name())
		assert.NotNil(span2.Root())
		assert.Equal("SPAN", span2.Root().TagName())
	})
}
