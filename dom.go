package dom

///////////////////////////////////////////////////////////////////////////////
// TYPES

type NodeType int

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// Node implements https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	// Properties
	BaseURI() string
	ChildNodes() []Node
	FirstChild() Node
	IsConnected() bool
	LastChild() Node
	NextSibling() Node
	NodeName() string
	NodeType() NodeType
	OwnerDocument() Document
	ParentNode() Node
	ParentElement() Element
	PreviousSibling() Node
	TextContent() string

	// Methods
	Equals(Node) bool
	AppendChild(Node) Node
	CloneNode() Node
	Contains(Node) bool
	HasChildNodes() bool
	InsertBefore(Node, Node) Node
	RemoveChild(Node)
	ReplaceChild(Node, Node)
}

type Element interface {
	Node

	// Properties
	InnerHTML() string
	OuterHTML() string
}

// Document implements https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document interface {
	Node

	// Properties
	Body() Element

	// Methods
	CreateElement(string) Element
}

type Window interface {
	// Properties
	Document() Document
}

///////////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	UNKNOWN_NODE NodeType = iota
	ELEMENT_NODE
	ATTRIBUTE_NODE
	TEXT_NODE
	CDATA_SECTION_NODE
	ENTITY_REFERENCE_NODE
	ENTITY_NODE
	PROCESSING_INSTRUCTION_NODE
	COMMENT_NODE
	DOCUMENT_NODE
	DOCUMENT_TYPE_NODE
	DOCUMENT_FRAGMENT_NODE
	NOTATION_NODE
)

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (t NodeType) String() string {
	switch t {
	case ELEMENT_NODE:
		return "ELEMENT_NODE"
	case ATTRIBUTE_NODE:
		return "ATTRIBUTE_NODE"
	case TEXT_NODE:
		return "TEXT_NODE"
	case CDATA_SECTION_NODE:
		return "CDATA_SECTION_NODE"
	case ENTITY_REFERENCE_NODE:
		return "ENTITY_REFERENCE_NODE"
	case ENTITY_NODE:
		return "ENTITY_NODE"
	case PROCESSING_INSTRUCTION_NODE:
		return "PROCESSING_INSTRUCTION_NODE"
	case COMMENT_NODE:
		return "COMMENT_NODE"
	case DOCUMENT_NODE:
		return "DOCUMENT_NODE"
	case DOCUMENT_TYPE_NODE:
		return "DOCUMENT_TYPE_NODE"
	case DOCUMENT_FRAGMENT_NODE:
		return "DOCUMENT_FRAGMENT_NODE"
	case NOTATION_NODE:
		return "NOTATION_NODE"
	default:
		return "UNKNOWN_NODE"
	}
}
