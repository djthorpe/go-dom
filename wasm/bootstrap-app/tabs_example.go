package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

func AddTabsExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Basic tabs
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Basic Tabs")).Element)
	card1Body := card1.Body()

	basicTabs := app.Tabs("basic-tabs")

	homePane := basicTabs.AddTab("Home", true, app)
	homePane.SetPadding("3")
	h3Home := app.CreateElement("h3")
	h3Home.AppendChild(app.CreateTextNode("Home Content"))
	homePane.AppendChild(h3Home)
	pHome := app.CreateElement("p")
	pHome.AppendChild(app.CreateTextNode("This is the home tab content. It's shown by default."))
	homePane.AppendChild(pHome)

	profilePane := basicTabs.AddTab("Profile", false, app)
	profilePane.SetPadding("3")
	h3Profile := app.CreateElement("h3")
	h3Profile.AppendChild(app.CreateTextNode("Profile Content"))
	profilePane.AppendChild(h3Profile)
	pProfile := app.CreateElement("p")
	pProfile.AppendChild(app.CreateTextNode("This is the profile tab content. Click the Profile tab to see it."))
	profilePane.AppendChild(pProfile)

	contactPane := basicTabs.AddTab("Contact", false, app)
	contactPane.SetPadding("3")
	h3Contact := app.CreateElement("h3")
	h3Contact.AppendChild(app.CreateTextNode("Contact Content"))
	contactPane.AppendChild(h3Contact)
	pContact := app.CreateElement("p")
	pContact.AppendChild(app.CreateTextNode("This is the contact tab content. Click the Contact tab to see it."))
	contactPane.AppendChild(pContact)

	card1Body.Element.AppendChild(basicTabs.Element)
	container.AppendChild(card1.Element)

	// Example 2: Tabs with Pills style
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Pills Style Tabs")).Element)
	card2Body := card2.Body()

	pillsTabs := app.Tabs("pills-tabs").SetStyle(bs5.TabStylePills)

	tab1 := pillsTabs.AddTab("Overview", true, app)
	tab1.SetPadding("3")
	p1 := app.CreateElement("p")
	p1.AppendChild(app.CreateTextNode("Pills style tabs have a more rounded, button-like appearance."))
	tab1.AppendChild(p1)

	tab2 := pillsTabs.AddTab("Features", false, app)
	tab2.SetPadding("3")
	ul := app.CreateElement("ul")
	li1 := app.CreateElement("li")
	li1.AppendChild(app.CreateTextNode("Easy navigation"))
	ul.AppendChild(li1)
	li2 := app.CreateElement("li")
	li2.AppendChild(app.CreateTextNode("Clean design"))
	ul.AppendChild(li2)
	li3 := app.CreateElement("li")
	li3.AppendChild(app.CreateTextNode("Responsive layout"))
	ul.AppendChild(li3)
	tab2.AppendChild(ul)

	tab3 := pillsTabs.AddTab("Settings", false, app)
	tab3.SetPadding("3")
	p3 := app.CreateElement("p")
	p3.AppendChild(app.CreateTextNode("Configure your preferences here."))
	tab3.AppendChild(p3)

	card2Body.Element.AppendChild(pillsTabs.Element)
	container.AppendChild(card2.Element)

	// Example 3: Justified tabs (fill width)
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Justified Tabs (Fill Width)")).Element)
	card3Body := card3.Body()

	justifiedTabs := app.Tabs("justified-tabs").SetJustified(true)

	j1 := justifiedTabs.AddTab("Dashboard", true, app)
	j1.SetPadding("3")
	pJ1 := app.CreateElement("p")
	pJ1.AppendChild(app.CreateTextNode("Welcome to your dashboard. These tabs fill the entire width."))
	j1.AppendChild(pJ1)

	j2 := justifiedTabs.AddTab("Analytics", false, app)
	j2.SetPadding("3")
	pJ2 := app.CreateElement("p")
	pJ2.AppendChild(app.CreateTextNode("View your analytics data here."))
	j2.AppendChild(pJ2)

	j3 := justifiedTabs.AddTab("Reports", false, app)
	j3.SetPadding("3")
	pJ3 := app.CreateElement("p")
	pJ3.AppendChild(app.CreateTextNode("Generate and view reports."))
	j3.AppendChild(pJ3)

	card3Body.Element.AppendChild(justifiedTabs.Element)
	container.AppendChild(card3.Element)

	return container.Element
}
