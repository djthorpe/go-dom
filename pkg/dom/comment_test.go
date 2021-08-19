package dom_test

import (
	"testing"

	// Modules
	"github.com/djthorpe/go-dom"
	. "github.com/djthorpe/go-dom/pkg/dom"
)

func Test_Comment_001(t *testing.T) {
	comment := GetWindow().Document().CreateComment("test")
	if comment.NodeType() != dom.COMMENT_NODE {
		t.Error("Expected COMMENT_NODE")
	} else if comment.NodeName() != "#comment" {
		t.Errorf("Expected #comment, got %q", comment.NodeName())
	} else {
		t.Log(comment)
	}
}
