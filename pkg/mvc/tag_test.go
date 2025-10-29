package mvc_test

import (
	"testing"

	// Packages
	assert "github.com/stretchr/testify/assert"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild/pkg/mvc"
)

func TestTagType(t *testing.T) {
	assert := assert.New(t)

	tag := Tag("UL")
	assert.NotNil(tag)

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(ViewTag, tag.Name())
		assert.NotNil(tag.Root())
		assert.Equal("UL", tag.Root().TagName())
	})

	t.Run("NewViewWithElement", func(t *testing.T) {
		tag2 := NewViewWithElement(tag.Root())
		assert.NotNil(tag2)
		assert.Equal(ViewTag, tag2.Name())
		assert.NotNil(tag2.Root())
		assert.Equal("UL", tag2.Root().TagName())
	})
}
