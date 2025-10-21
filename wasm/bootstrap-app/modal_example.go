package main

import (
	"fmt"

	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddModalExample adds modal component examples to the app and returns the container
func AddModalExample(app *bs5.App) dom.Element {
	container := app.Container()

	// Example 1: Standard centered modal
	card1 := app.Card()
	card1.AddClass("mb-4")
	card1.Header(app.H4(app.CreateTextNode("Centered Modal")).Element)
	card1Body := card1.Body()

	descP1 := app.CreateElement("p")
	descP1.AppendChild(app.CreateTextNode("A modal that is vertically centered in the viewport."))
	card1Body.Element.AppendChild(descP1)

	btn1 := app.Button(bs5.ColorPrimary, app.CreateTextNode("Launch Centered Modal"))
	btn1.Element.SetAttribute("data-bs-toggle", "modal")
	btn1.Element.SetAttribute("data-bs-target", "#centeredModal")
	card1Body.Element.AppendChild(btn1.Element)
	container.AppendChild(card1.Element)

	modal1 := app.Modal("centeredModal")
	modal1.SetCentered(true)
	modal1.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal1.Header().AddTitle("Centered Modal").AddCloseButton()
	modal1.Body(app.CreateTextNode("This modal is vertically centered in the viewport."))
	closeBtn1 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Close"))
	closeBtn1.Element.SetAttribute("data-bs-dismiss", "modal")
	modal1.Footer(closeBtn1.Element)
	app.Document.Body().AppendChild(modal1.Element)

	// Example 2: Login Form Modal
	card2 := app.Card()
	card2.AddClass("mb-4")
	card2.Header(app.H4(app.CreateTextNode("Login Form Modal")).Element)
	card2Body := card2.Body()

	descP2 := app.CreateElement("p")
	descP2.AppendChild(app.CreateTextNode("A modal containing a login form with email and password fields, using input groups with icons."))
	card2Body.Element.AppendChild(descP2)

	btn2 := app.Button(bs5.ColorSuccess, app.CreateTextNode("Launch Login Modal"))
	btn2.Element.SetAttribute("data-bs-toggle", "modal")
	btn2.Element.SetAttribute("data-bs-target", "#loginModal")
	card2Body.Element.AppendChild(btn2.Element)
	container.AppendChild(card2.Element)

	modal2 := app.Modal("loginModal")
	modal2.SetCentered(true)
	modal2.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal2.Header().AddTitle("Login").AddCloseButton()

	// Create login form
	loginForm := app.CreateElement("form")

	// Email field with envelope icon
	emailGroup := app.CreateElement("div")
	emailGroup.AddClass("mb-3")
	emailLabel := app.FormLabel("loginEmail", app.CreateTextNode("Email"))
	emailGroup.AppendChild(emailLabel.Element)

	emailInputGroup := app.InputGroup()
	emailIcon := app.Icon("envelope-fill")
	emailIconText := app.InputGroupText(emailIcon.Element)
	emailInputGroup.Prepend(emailIconText.Element)

	emailInput := app.FormInput("loginEmail", "email")
	emailInput.SetPlaceholder("Enter your email")
	emailInput.SetRequired(true)
	emailInputGroup.AppendChild(emailInput.Element)

	emailGroup.AppendChild(emailInputGroup.Element)
	loginForm.AppendChild(emailGroup)

	// Password field with lock icon
	passwordGroup := app.CreateElement("div")
	passwordGroup.AddClass("mb-3")
	passwordLabel := app.FormLabel("loginPassword", app.CreateTextNode("Password"))
	passwordGroup.AppendChild(passwordLabel.Element)

	passwordInputGroup := app.InputGroup()
	lockIcon := app.Icon("lock-fill")
	lockIconText := app.InputGroupText(lockIcon.Element)
	passwordInputGroup.Prepend(lockIconText.Element)

	passwordInput := app.FormInput("loginPassword", "password")
	passwordInput.SetPlaceholder("Enter your password")
	passwordInput.SetRequired(true)
	passwordInputGroup.AppendChild(passwordInput.Element)

	passwordGroup.AppendChild(passwordInputGroup.Element)
	loginForm.AppendChild(passwordGroup)

	// Remember me checkbox
	rememberGroup := app.CreateElement("div")
	rememberGroup.AddClass("mb-3")
	rememberCheck := app.FormCheck("loginRemember", "checkbox", "Remember me")
	rememberGroup.AppendChild(rememberCheck.Element)
	loginForm.AppendChild(rememberGroup)

	modal2.Body(loginForm)

	cancelBtn2 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Cancel"))
	cancelBtn2.Element.SetAttribute("data-bs-dismiss", "modal")
	loginBtn2 := app.Button(bs5.ColorPrimary, app.CreateTextNode("Login")).
		AddEventListener("click", func(target dom.Node) {
			fmt.Println("Login button clicked!")
			// In a real app, you would validate and submit the form here
		})
	modal2.Footer(cancelBtn2.Element, app.CreateTextNode(" "), loginBtn2.Element)
	app.Document.Body().AppendChild(modal2.Element)

	// Example 3: Scrollable modal with long content
	card3 := app.Card()
	card3.AddClass("mb-4")
	card3.Header(app.H4(app.CreateTextNode("Scrollable Modal")).Element)
	card3Body := card3.Body()

	descP3 := app.CreateElement("p")
	descP3.AppendChild(app.CreateTextNode("A modal with scrollable content when it exceeds the viewport height."))
	card3Body.Element.AppendChild(descP3)

	btn3 := app.Button(bs5.ColorInfo, app.CreateTextNode("Launch Scrollable Modal"))
	btn3.Element.SetAttribute("data-bs-toggle", "modal")
	btn3.Element.SetAttribute("data-bs-target", "#scrollableModal")
	card3Body.Element.AppendChild(btn3.Element)
	container.AppendChild(card3.Element)

	modal3 := app.Modal("scrollableModal")
	modal3.SetScrollable(true)
	modal3.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal3.Header().AddTitle("Scrollable Modal").AddCloseButton()

	// Create long content
	bodyContent := app.CreateElement("div")
	for i := 1; i <= 20; i++ {
		p := app.CreateElement("p")
		p.AppendChild(app.CreateTextNode(fmt.Sprintf("Paragraph %d: This is a scrollable modal with long content. The modal body will scroll independently when content exceeds the viewport height.", i)))
		bodyContent.AppendChild(p)
	}
	modal3.Body(bodyContent)

	closeBtn3 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Close"))
	closeBtn3.Element.SetAttribute("data-bs-dismiss", "modal")
	modal3.Footer(closeBtn3.Element)
	app.Document.Body().AppendChild(modal3.Element)

	// Example 4: Large modal
	card4 := app.Card()
	card4.AddClass("mb-4")
	card4.Header(app.H4(app.CreateTextNode("Large Modal")).Element)
	card4Body := card4.Body()

	descP4 := app.CreateElement("p")
	descP4.AppendChild(app.CreateTextNode("A larger modal providing more space for content."))
	card4Body.Element.AppendChild(descP4)

	btn4 := app.Button(bs5.ColorWarning, app.CreateTextNode("Launch Large Modal"))
	btn4.Element.SetAttribute("data-bs-toggle", "modal")
	btn4.Element.SetAttribute("data-bs-target", "#largeModal")
	card4Body.Element.AppendChild(btn4.Element)
	container.AppendChild(card4.Element)

	modal4 := app.Modal("largeModal")
	modal4.SetSize(bs5.ModalSizeLarge)
	modal4.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal4.Header().AddTitle("Large Modal").AddCloseButton()
	modal4.Body(app.CreateTextNode("This modal is larger than the default size, providing more space for content."))
	closeBtn4 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Close"))
	closeBtn4.Element.SetAttribute("data-bs-dismiss", "modal")
	modal4.Footer(closeBtn4.Element)
	app.Document.Body().AppendChild(modal4.Element)

	// Example 5: Small modal
	card5 := app.Card()
	card5.AddClass("mb-4")
	card5.Header(app.H4(app.CreateTextNode("Small Modal")).Element)
	card5Body := card5.Body()

	descP5 := app.CreateElement("p")
	descP5.AppendChild(app.CreateTextNode("A smaller modal perfect for simple confirmations."))
	card5Body.Element.AppendChild(descP5)

	btn5 := app.Button(bs5.ColorDanger, app.CreateTextNode("Launch Small Modal"))
	btn5.Element.SetAttribute("data-bs-toggle", "modal")
	btn5.Element.SetAttribute("data-bs-target", "#smallModal")
	card5Body.Element.AppendChild(btn5.Element)
	container.AppendChild(card5.Element)

	modal5 := app.Modal("smallModal")
	modal5.SetSize(bs5.ModalSizeSmall)
	modal5.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal5.Header().AddTitle("Small Modal").AddCloseButton()
	modal5.Body(app.CreateTextNode("This is a small modal, perfect for simple confirmations."))
	closeBtn5 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Close"))
	closeBtn5.Element.SetAttribute("data-bs-dismiss", "modal")
	saveBtn5 := app.Button(bs5.ColorPrimary, app.CreateTextNode("Confirm")).
		AddEventListener("click", func(target dom.Node) {
			fmt.Println("Confirmed!")
		})
	modal5.Footer(closeBtn5.Element, app.CreateTextNode(" "), saveBtn5.Element)
	app.Document.Body().AppendChild(modal5.Element)

	// Example 6: Fullscreen modal
	card6 := app.Card()
	card6.AddClass("mb-4")
	card6.Header(app.H4(app.CreateTextNode("Fullscreen Modal")).Element)
	card6Body := card6.Body()

	descP6 := app.CreateElement("p")
	descP6.AppendChild(app.CreateTextNode("A modal that takes up the entire viewport."))
	card6Body.Element.AppendChild(descP6)

	btn6 := app.Button(bs5.ColorPrimary, app.CreateTextNode("Launch Fullscreen Modal"))
	btn6.Element.SetAttribute("data-bs-toggle", "modal")
	btn6.Element.SetAttribute("data-bs-target", "#fullscreenModal")
	card6Body.Element.AppendChild(btn6.Element)
	container.AppendChild(card6.Element)

	modal6 := app.Modal("fullscreenModal")
	modal6.SetFullscreen(true)
	modal6.AddEventListener("hide.bs.modal", func(target dom.Node) {
		if activeEl := app.Document.ActiveElement(); activeEl != nil {
			activeEl.Blur()
		}
	})
	modal6.Header().AddTitle("Fullscreen Modal").AddCloseButton()
	modal6.Body(app.CreateTextNode("This modal takes up the entire viewport. Perfect for displaying full-page content or forms."))
	closeBtn6 := app.Button(bs5.ColorSecondary, app.CreateTextNode("Close"))
	closeBtn6.Element.SetAttribute("data-bs-dismiss", "modal")
	modal6.Footer(closeBtn6.Element)
	app.Document.Body().AppendChild(modal6.Element)

	return container.Element
}
