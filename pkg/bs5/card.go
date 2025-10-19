package bs5

import (
	"github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type Card struct {
	dom.Element
	header *CardHeader
	body   *CardBody
	footer *CardFooter
	img    *CardImg
}

type CardHeader struct {
	dom.Element
	card *Card
}

type CardBody struct {
	dom.Element
	card *Card
}

type CardFooter struct {
	dom.Element
	card *Card
}

type CardImg struct {
	dom.Element
	card *Card
}

type CardTitle struct {
	dom.Element
}

type CardText struct {
	dom.Element
}

type CardImgPosition string

////////////////////////////////////////////////////////////////////////
// CONSTANTS

const (
	CardImgTop    CardImgPosition = "top"
	CardImgBottom CardImgPosition = "bottom"
)

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// Card creates a Bootstrap 5 card component
func (app *App) Card() *Card {
	// <div class="card">
	card := app.CreateElement("div")
	card.AddClass("card")

	return &Card{
		Element: card,
	}
}

////////////////////////////////////////////////////////////////////////
// CARD METHODS

// Header creates or returns the card header
func (c *Card) Header(children ...dom.Node) *CardHeader {
	if c.header == nil {
		// <div class="card-header">
		header := c.Element.OwnerDocument().CreateElement("div")
		header.AddClass("card-header")

		// Add children
		for _, child := range children {
			header.AppendChild(child)
		}

		// Insert at the beginning (after image if present)
		if c.img != nil && c.img.Element.ParentNode() != nil {
			c.Element.InsertBefore(header, c.img.Element.NextSibling())
		} else {
			if c.Element.FirstChild() != nil {
				c.Element.InsertBefore(header, c.Element.FirstChild())
			} else {
				c.Element.AppendChild(header)
			}
		}

		c.header = &CardHeader{
			Element: header,
			card:    c,
		}
	}
	return c.header
}

// Body creates or returns the card body
func (c *Card) Body(children ...dom.Node) *CardBody {
	if c.body == nil {
		// <div class="card-body">
		body := c.Element.OwnerDocument().CreateElement("div")
		body.AddClass("card-body")

		// Add children
		for _, child := range children {
			body.AppendChild(child)
		}

		c.Element.AppendChild(body)
		c.body = &CardBody{
			Element: body,
			card:    c,
		}
	}
	return c.body
}

// Footer creates or returns the card footer
func (c *Card) Footer(children ...dom.Node) *CardFooter {
	if c.footer == nil {
		// <div class="card-footer">
		footer := c.Element.OwnerDocument().CreateElement("div")
		footer.AddClass("card-footer")

		// Add children
		for _, child := range children {
			footer.AppendChild(child)
		}

		c.Element.AppendChild(footer)
		c.footer = &CardFooter{
			Element: footer,
			card:    c,
		}
	}
	return c.footer
}

// Image adds an image to the card
func (c *Card) Image(src string, alt string, position CardImgPosition) *CardImg {
	if c.img == nil {
		// <img src="..." class="card-img-top" alt="...">
		img := c.Element.OwnerDocument().CreateElement("img")
		img.SetAttribute("src", src)
		img.SetAttribute("alt", alt)

		if position == CardImgTop {
			img.AddClass("card-img-top")
			// Insert at the beginning
			if c.Element.FirstChild() != nil {
				c.Element.InsertBefore(img, c.Element.FirstChild())
			} else {
				c.Element.AppendChild(img)
			}
		} else {
			img.AddClass("card-img-bottom")
			// Append at the end
			c.Element.AppendChild(img)
		}

		c.img = &CardImg{
			Element: img,
			card:    c,
		}
	}
	return c.img
}

// SetWidth sets the card width using inline style
func (c *Card) SetWidth(width string) *Card {
	style := c.Element.Style()
	if style != nil {
		style.Set("width", width)
	}
	return c
}

// SetTextAlign sets the text alignment
func (c *Card) SetTextAlign(align string) *Card {
	c.Element.AddClass("text-" + align)
	return c
}

// SetBorder sets the border color
func (c *Card) SetBorder(color ColorVariant) *Card {
	c.Element.AddClass("border-" + string(color))
	return c
}

// SetBackground sets the background color
func (c *Card) SetBackground(color ColorVariant) *Card {
	c.Element.AddClass("bg-" + string(color))
	return c
}

// AddClass adds a CSS class to the card
func (c *Card) AddClass(className string) *Card {
	c.Element.AddClass(className)
	return c
}

// RemoveClass removes a CSS class from the card
func (c *Card) RemoveClass(className string) *Card {
	c.Element.RemoveClass(className)
	return c
}

////////////////////////////////////////////////////////////////////////
// CARD HEADER METHODS

// AddClass adds a CSS class to the card header
func (h *CardHeader) AddClass(className string) *CardHeader {
	h.Element.AddClass(className)
	return h
}

// RemoveClass removes a CSS class from the card header
func (h *CardHeader) RemoveClass(className string) *CardHeader {
	h.Element.RemoveClass(className)
	return h
}

////////////////////////////////////////////////////////////////////////
// CARD BODY METHODS

// Title creates a card title element (h5)
func (b *CardBody) Title(text string) *CardTitle {
	// <h5 class="card-title">
	title := b.Element.OwnerDocument().CreateElement("h5")
	title.AddClass("card-title")
	title.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
	b.Element.AppendChild(title)

	return &CardTitle{
		Element: title,
	}
}

// Subtitle creates a card subtitle element (h6)
func (b *CardBody) Subtitle(text string) *CardTitle {
	// <h6 class="card-subtitle mb-2 text-muted">
	subtitle := b.Element.OwnerDocument().CreateElement("h6")
	subtitle.AddClass("card-subtitle")
	subtitle.AddClass("mb-2")
	subtitle.AddClass("text-muted")
	subtitle.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
	b.Element.AppendChild(subtitle)

	return &CardTitle{
		Element: subtitle,
	}
}

// Text creates a card text paragraph
func (b *CardBody) Text(text string) *CardText {
	// <p class="card-text">
	p := b.Element.OwnerDocument().CreateElement("p")
	p.AddClass("card-text")
	p.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
	b.Element.AppendChild(p)

	return &CardText{
		Element: p,
	}
}

// Link creates a card link
func (b *CardBody) Link(text string, href string) dom.Element {
	// <a href="#" class="card-link">
	link := b.Element.OwnerDocument().CreateElement("a")
	link.AddClass("card-link")
	link.SetAttribute("href", href)
	link.AppendChild(b.Element.OwnerDocument().CreateTextNode(text))
	b.Element.AppendChild(link)
	return link
}

// AddClass adds a CSS class to the card body
func (b *CardBody) AddClass(className string) *CardBody {
	b.Element.AddClass(className)
	return b
}

// RemoveClass removes a CSS class from the card body
func (b *CardBody) RemoveClass(className string) *CardBody {
	b.Element.RemoveClass(className)
	return b
}

////////////////////////////////////////////////////////////////////////
// CARD FOOTER METHODS

// AddClass adds a CSS class to the card footer
func (f *CardFooter) AddClass(className string) *CardFooter {
	f.Element.AddClass(className)
	return f
}

// RemoveClass removes a CSS class from the card footer
func (f *CardFooter) RemoveClass(className string) *CardFooter {
	f.Element.RemoveClass(className)
	return f
}

// SetMuted sets the text to muted color
func (f *CardFooter) SetMuted(muted bool) *CardFooter {
	if muted {
		f.Element.AddClass("text-muted")
	} else {
		f.Element.RemoveClass("text-muted")
	}
	return f
}

////////////////////////////////////////////////////////////////////////
// CARD IMAGE METHODS

// SetAlt updates the alt text
func (img *CardImg) SetAlt(alt string) *CardImg {
	img.Element.SetAttribute("alt", alt)
	return img
}

// AddClass adds a CSS class to the card image
func (img *CardImg) AddClass(className string) *CardImg {
	img.Element.AddClass(className)
	return img
}

// RemoveClass removes a CSS class from the card image
func (img *CardImg) RemoveClass(className string) *CardImg {
	img.Element.RemoveClass(className)
	return img
}

////////////////////////////////////////////////////////////////////////
// CARD TITLE METHODS

// AddClass adds a CSS class to the card title
func (t *CardTitle) AddClass(className string) *CardTitle {
	t.Element.AddClass(className)
	return t
}

// RemoveClass removes a CSS class from the card title
func (t *CardTitle) RemoveClass(className string) *CardTitle {
	t.Element.RemoveClass(className)
	return t
}

////////////////////////////////////////////////////////////////////////
// CARD TEXT METHODS

// AddClass adds a CSS class to the card text
func (t *CardText) AddClass(className string) *CardText {
	t.Element.AddClass(className)
	return t
}

// RemoveClass removes a CSS class from the card text
func (t *CardText) RemoveClass(className string) *CardText {
	t.Element.RemoveClass(className)
	return t
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (c *Card) String() string {
	return "<bs5-card>"
}

func (h *CardHeader) String() string {
	return "<bs5-card-header>"
}

func (b *CardBody) String() string {
	return "<bs5-card-body>"
}

func (f *CardFooter) String() string {
	return "<bs5-card-footer>"
}

func (img *CardImg) String() string {
	return "<bs5-card-img>"
}

func (t *CardTitle) String() string {
	return "<bs5-card-title>"
}

func (txt *CardText) String() string {
	return "<bs5-card-text>"
}
