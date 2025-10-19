package dom

import "io"

///////////////////////////////////////////////////////////////////////////////
// TYPES

type NodeType int

///////////////////////////////////////////////////////////////////////////////
// INTERFACES

// Node implements https://developer.mozilla.org/en-US/docs/Web/API/Node
type Node interface {
	// Properties
	ChildNodes() []Node
	Contains(Node) bool
	Equals(Node) bool
	FirstChild() Node
	HasChildNodes() bool
	IsConnected() bool
	LastChild() Node
	NextSibling() Node
	NodeName() string
	NodeType() NodeType
	OwnerDocument() Document
	ParentElement() Element
	ParentNode() Node
	PreviousSibling() Node
	TextContent() string

	// Methods
	AppendChild(Node) Node
	CloneNode(bool) Node
	InsertBefore(Node, Node) Node
	RemoveChild(Node)
	ReplaceChild(Node, Node)
}

type Element interface {
	Node

	// Properties
	InnerHTML() string
	OuterHTML() string
	TagName() string
	Attributes() []Attr
	Style() Style

	// Attribute Methods
	//RemoveAttrbute(string)
	//RemoveAttributeNode(Attr) Attr
	SetAttribute(string, string) Attr
	//SetAttributeNode(Attr) Attr
	GetAttribute(string) Attr
	//GetAttributeNames() []string
	//GetAttributeNode(string) Attr
	//HasAttribute(string) bool
	HasAttributes() bool

	// Class Methods
	AddClass(string)
	RemoveClass(string)
}

// Document implements https://developer.mozilla.org/en-US/docs/Web/API/Document
type Document interface {
	Node

	// Properties
	Body() Element
	//CharacterSet() string
	//ContentType() string
	Doctype() DocumentType
	//DocumentElement() Element
	//DocumentURI() string
	//Head() Element
	Title() string

	// Methods
	CreateElement(string) Element
	CreateAttribute(string) Attr
	CreateComment(string) Comment
	CreateTextNode(string) Text
}

type Text interface {
	Node

	// Properties
	Data() string
	Length() int
}

type Comment interface {
	Node

	// Properties
	Data() string
	Length() int
}

type Attr interface {
	Node

	// Properties
	OwnerElement() Element
	Name() string
	Value() string
	SetValue(string)
}

// Style implements https://developer.mozilla.org/en-US/docs/Web/API/CSSStyleDeclaration
type Style interface {
	// Methods
	Get(string) string
	Set(string, string)
}

// Document implements https://developer.mozilla.org/en-US/docs/Web/API/DocumentType
type DocumentType interface {
	Node

	// Properties
	Name() string
	PublicId() string
	SystemId() string
}

type Window interface {
	// Properties
	Document() Document

	// Methods
	Write(io.Writer, Node) (int, error)
	//Read(io.Reader, string) (Document, error)
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
