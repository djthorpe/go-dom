package bs5

import (
	"strconv"

	// Packages
	"github.com/djthorpe/go-dom/pkg/dom"

	// Namespace imports
	. "github.com/djthorpe/go-dom"
)

////////////////////////////////////////////////////////////////////////
// TYPES

type App struct {
	Document
}

////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func New(title string) *App {
	doc := dom.GetWindowWithTitle(title).Document()
	return &App{
		doc,
	}
}

////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (app *App) String() string {
	str := "<bs5-app"
	str += " title=" + strconv.Quote(app.Title())
	return str + ">"
}
