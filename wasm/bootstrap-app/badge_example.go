package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddBadgeExamples adds badge component examples to the app
func AddBadgeExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Heading with badges with icons
	headingCard := app.Card()
	headingCard.AddClass("mb-4")
	headingCard.Header(app.H4(app.CreateTextNode("Headings with Icon Badges")).Element)

	headingCardBody := headingCard.Body()

	h1Content := app.H1()
	h1Content.Element.AppendChild(app.CreateTextNode("Main Title "))

	newBadgeContent := app.CreateElement("span")
	starIcon := app.Icon("star-fill")
	starIcon.AddClass("me-1")
	newBadgeContent.AppendChild(starIcon.Element)
	newBadgeContent.AppendChild(app.CreateTextNode("New"))
	h1Content.Element.AppendChild(app.Badge(bs5.ColorSecondary, newBadgeContent).Element)
	h1Content.Element.AppendChild(app.CreateTextNode(" "))

	featuredBadgeContent := app.CreateElement("span")
	trophyIcon := app.Icon("trophy-fill")
	trophyIcon.AddClass("me-1")
	featuredBadgeContent.AppendChild(trophyIcon.Element)
	featuredBadgeContent.AppendChild(app.CreateTextNode("Featured"))
	h1Content.Element.AppendChild(app.Badge(bs5.ColorPrimary, featuredBadgeContent).Element)

	headingCardBody.Element.AppendChild(h1Content.Element)
	container.AppendChild(headingCard.Element)

	// Example 2: Inline badges with icons
	inlineCard := app.Card()
	inlineCard.AddClass("mb-4")
	inlineCard.Header(app.H4(app.CreateTextNode("Inline Badges with Icons")).Element)

	inlineCardBody := inlineCard.Body()

	inlineContent := app.CreateElement("div")
	inlineContent.AppendChild(app.CreateTextNode("Hello, World! "))

	successBadgeContent := app.CreateElement("span")
	checkIcon := app.Icon("check-circle-fill")
	checkIcon.AddClass("me-1")
	successBadgeContent.AppendChild(checkIcon.Element)
	successBadgeContent.AppendChild(app.CreateTextNode("Success"))
	inlineContent.AppendChild(app.Badge(bs5.ColorSuccess, successBadgeContent).Element)
	inlineContent.AppendChild(app.CreateTextNode(" "))

	dangerBadgeContent := app.CreateElement("span")
	fireIcon := app.Icon("fire")
	fireIcon.AddClass("me-1")
	dangerBadgeContent.AppendChild(fireIcon.Element)
	dangerBadgeContent.AppendChild(app.CreateTextNode("Hot"))
	inlineContent.AppendChild(app.Badge(bs5.ColorDanger, dangerBadgeContent).Element)
	inlineContent.AppendChild(app.CreateTextNode(" "))

	warningBadgeContent := app.CreateElement("span")
	exclamationIcon := app.Icon("exclamation-triangle-fill")
	exclamationIcon.AddClass("me-1")
	warningBadgeContent.AppendChild(exclamationIcon.Element)
	warningBadgeContent.AppendChild(app.CreateTextNode("Warning"))
	inlineContent.AppendChild(app.Badge(bs5.ColorWarning, warningBadgeContent).Element)
	inlineContent.AppendChild(app.CreateTextNode(" "))

	infoBadgeContent := app.CreateElement("span")
	infoIcon := app.Icon("info-circle-fill")
	infoIcon.AddClass("me-1")
	infoBadgeContent.AppendChild(infoIcon.Element)
	infoBadgeContent.AppendChild(app.CreateTextNode("Info"))
	inlineContent.AppendChild(app.Badge(bs5.ColorInfo, infoBadgeContent).Element)

	inlineCardBody.Element.AppendChild(inlineContent)
	container.AppendChild(inlineCard.Element)

	// Example 3: Icon-only badges
	iconOnlyCard := app.Card()
	iconOnlyCard.AddClass("mb-4")
	iconOnlyCard.Header(app.H4(app.CreateTextNode("Icon-Only Badges")).Element)

	iconOnlyCardBody := iconOnlyCard.Body()

	iconOnlyContent := app.CreateElement("div")
	iconOnlyContent.AddClass("d-flex")
	iconOnlyContent.AddClass("gap-2")
	iconOnlyContent.AddClass("align-items-center")

	iconOnlyContent.AppendChild(app.CreateTextNode("Quick actions: "))
	iconOnlyContent.AppendChild(app.Badge(bs5.ColorPrimary, app.Icon("envelope-fill").Element).Element)
	iconOnlyContent.AppendChild(app.Badge(bs5.ColorSuccess, app.Icon("bell-fill").Element).Element)
	iconOnlyContent.AppendChild(app.Badge(bs5.ColorDanger, app.Icon("heart-fill").Element).Element)
	iconOnlyContent.AppendChild(app.Badge(bs5.ColorWarning, app.Icon("gear-fill").Element).Element)
	iconOnlyContent.AppendChild(app.Badge(bs5.ColorInfo, app.Icon("calendar-fill").Element).Element)

	iconOnlyCardBody.Element.AppendChild(iconOnlyContent)
	container.AppendChild(iconOnlyCard.Element)

	// Example 4: Notification badges with icons and counts
	notificationCard := app.Card()
	notificationCard.Header(app.H4(app.CreateTextNode("Notification Badges with Icons")).Element)

	notificationCardBody := notificationCard.Body()

	notificationContent := app.CreateElement("div")
	notificationContent.AddClass("d-flex")
	notificationContent.AddClass("flex-column")
	notificationContent.AddClass("gap-3")

	messagesRow := app.CreateElement("div")
	messagesIcon := app.Icon("envelope")
	messagesIcon.SetSize("fs-5")
	messagesRow.AppendChild(messagesIcon.Element)
	messagesRow.AppendChild(app.CreateTextNode(" Messages "))
	messagesBadge := app.Badge(bs5.ColorPrimary, app.CreateTextNode("12"))
	messagesBadge.AddClass("rounded-pill")
	messagesRow.AppendChild(messagesBadge.Element)
	notificationContent.AppendChild(messagesRow)

	alertsRow := app.CreateElement("div")
	alertsIcon := app.Icon("bell")
	alertsIcon.SetSize("fs-5")
	alertsRow.AppendChild(alertsIcon.Element)
	alertsRow.AppendChild(app.CreateTextNode(" Alerts "))
	alertsBadge := app.Badge(bs5.ColorDanger, app.CreateTextNode("3"))
	alertsBadge.AddClass("rounded-pill")
	alertsRow.AppendChild(alertsBadge.Element)
	notificationContent.AppendChild(alertsRow)

	tasksRow := app.CreateElement("div")
	tasksIcon := app.Icon("check2-square")
	tasksIcon.SetSize("fs-5")
	tasksRow.AppendChild(tasksIcon.Element)
	tasksRow.AppendChild(app.CreateTextNode(" Tasks "))
	tasksBadge := app.Badge(bs5.ColorSuccess, app.CreateTextNode("8"))
	tasksBadge.AddClass("rounded-pill")
	tasksRow.AppendChild(tasksBadge.Element)
	notificationContent.AppendChild(tasksRow)

	notificationCardBody.Element.AppendChild(notificationContent)
	container.AppendChild(notificationCard.Element)

	return container.Element
}
