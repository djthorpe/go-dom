package main

import (
	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"

	// Namespace import for interfaces
	. "github.com/djthorpe/go-wasmbuild"
)

// FormExamples returns a container with various form examples
func FormExamples() Component {
	container := bs.Container(
		bs.WithBreakpoint(bs.BreakpointLarge),
		bs.WithMargin(bs.TOP, 4),
	)

	// Section heading
	container.Append(
		bs.Heading(2, bs.WithMargin(bs.BOTTOM, 4)).Append("Form Examples"),
	)

	// Basic form
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3)).Append("Basic Form"),
	)

	basicForm := bs.Form(
		bs.WithAction("/submit"),
		bs.WithMethod("POST"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Form submitted!")
	})

	basicForm.Append(
		bs.Para().Append("This is a basic form with default Bootstrap validation."),
	)

	container.Append(basicForm)

	// Form with custom validation
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with Custom Validation"),
	)

	validationForm := bs.Form(
		bs.WithAction("/api/login"),
		bs.WithMethod("POST"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		// Custom validation logic would go here
		println("Validation form submitted!")
	})

	validationForm.Append(
		bs.Para().Append("Forms have novalidate attribute and needs-validation class by default."),
		bs.Para(bs.WithClass("text-muted", "small")).Append(
			"Use OnSubmit to validate and add 'was-validated' class to show feedback.",
		),
	)

	container.Append(validationForm)

	// Form with browser validation
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with Browser Validation"),
	)

	browserForm := bs.Form(
		bs.WithAction("/submit"),
		bs.WithMethod("POST"),
		bs.WithoutValidation(), // Enables browser HTML5 validation
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Browser validation form submitted!")
	})

	browserForm.Append(
		bs.Para().Append("This form uses browser's native HTML5 validation."),
		bs.Para(bs.WithClass("text-muted", "small")).Append(
			"WithoutValidation() option removes the novalidate attribute.",
		),
	)

	container.Append(browserForm)

	// Form with file upload
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with File Upload"),
	)

	fileForm := bs.Form(
		bs.WithAction("/upload"),
		bs.WithMethod("POST"),
		bs.WithEnctype("multipart/form-data"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("File upload form submitted!")
	})

	fileForm.Append(
		bs.Para().Append("This form is configured for file uploads with multipart/form-data encoding."),
	)

	container.Append(fileForm)

	// Inline form layout
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Inline Form Layout"),
	)

	inlineForm := bs.Form(
		bs.WithAction("/search"),
		bs.WithMethod("GET"),
		bs.WithClass("row", "g-3", "align-items-center"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Inline form submitted!")
	})

	inlineForm.Append(
		bs.Para(bs.WithClass("col-auto")).Append("Search:"),
		bs.Para(bs.WithClass("col-auto", "text-muted", "small")).Append(
			"(Inline forms use Bootstrap grid classes)",
		),
	)

	container.Append(inlineForm)

	// Horizontal form layout
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Horizontal Form Layout"),
	)

	horizontalForm := bs.Form(
		bs.WithAction("/profile"),
		bs.WithMethod("POST"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Horizontal form submitted!")
	})

	horizontalForm.Append(
		bs.Container(bs.WithClass("row", "mb-3")).Append(
			bs.Para(bs.WithClass("col-sm-2", "col-form-label")).Append("Email:"),
			bs.Para(bs.WithClass("col-sm-10", "text-muted", "small")).Append(
				"(Use Bootstrap row/col classes for horizontal layout)",
			),
		),
	)

	container.Append(horizontalForm)

	// Form with colored theme
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with Theme"),
	)

	themedForm := bs.Form(
		bs.WithAction("/submit"),
		bs.WithMethod("POST"),
		bs.WithTheme(bs.DARK),
		bs.WithBackground(bs.DARK),
		bs.WithClass("p-4", "rounded"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Themed form submitted!")
	})

	themedForm.Append(
		bs.Para(bs.WithClass("text-light")).Append("This form uses a dark theme."),
	)

	container.Append(themedForm)

	// Form with custom classes
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with Custom Styling"),
	)

	customForm := bs.Form(
		bs.WithAction("/submit"),
		bs.WithMethod("POST"),
		bs.WithClass("border", "p-4", "rounded", "shadow-sm"),
		bs.WithBorder(bs.BorderAll, bs.PRIMARY),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Custom styled form submitted!")
	})

	customForm.Append(
		bs.Para().Append("This form has custom border, padding, and shadow classes."),
	)

	container.Append(customForm)

	// Form with grid layout
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form with Grid Layout"),
	)

	gridForm := bs.Form(
		bs.WithAction("/submit"),
		bs.WithMethod("POST"),
		bs.WithClass("p-3", "bg-light", "rounded"),
	).OnSubmit(func(e Event) {
		e.PreventDefault()
		println("Grid form submitted!")
	})

	gridForm.Append(
		bs.Container(bs.WithClass("row", "g-3")).Append(
			bs.Para(bs.WithClass("col-md-6")).Append("First column (50% width on medium screens)"),
			bs.Para(bs.WithClass("col-md-6")).Append("Second column (50% width on medium screens)"),
			bs.Para(bs.WithClass("col-12")).Append("Full width column"),
		),
	)

	container.Append(gridForm)

	// Select examples
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Select Dropdown Examples"),
	)

	// Basic select
	basicSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Basic Select"),
		bs.Select(bs.WithName("country")).Append(
			bs.Option().Append("Choose a country..."),
			bs.Option(bs.WithValue("us")).Append("United States"),
			bs.Option(bs.WithValue("uk")).Append("United Kingdom"),
			bs.Option(bs.WithValue("ca")).Append("Canada"),
			bs.Option(bs.WithValue("au")).Append("Australia"),
			bs.Option(bs.WithValue("de")).Append("Germany"),
		),
	)
	container.Append(basicSelect)

	// Select with selected option
	selectedSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Select with Default Selection"),
		bs.Select(bs.WithName("size")).Append(
			bs.Option(bs.WithValue("sm")).Append("Small"),
			bs.Option(bs.WithValue("md"), bs.WithSelected()).Append("Medium (Selected)"),
			bs.Option(bs.WithValue("lg")).Append("Large"),
			bs.Option(bs.WithValue("xl")).Append("Extra Large"),
		),
	)
	container.Append(selectedSelect)

	// Select with required
	requiredSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Required Select"),
		bs.Select(bs.WithName("department"), bs.WithRequired()).Append(
			bs.Option(bs.WithValue("")).Append("Select department..."),
			bs.Option(bs.WithValue("sales")).Append("Sales"),
			bs.Option(bs.WithValue("marketing")).Append("Marketing"),
			bs.Option(bs.WithValue("engineering")).Append("Engineering"),
			bs.Option(bs.WithValue("support")).Append("Support"),
		),
		bs.Para(bs.WithClass("form-text", "text-muted", "small")).Append("This field is required."),
	)
	container.Append(requiredSelect)

	// Select with disabled
	disabledSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Disabled Select"),
		bs.Select(bs.WithName("disabled"), bs.WithDisabled()).Append(
			bs.Option().Append("This select is disabled"),
			bs.Option(bs.WithValue("1")).Append("Option 1"),
			bs.Option(bs.WithValue("2")).Append("Option 2"),
		),
	)
	container.Append(disabledSelect)

	// Multiple select
	multipleSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Multiple Select"),
		bs.Select(bs.WithName("colors"), bs.WithMultiple(), bs.WithClass("form-select")).Append(
			bs.Option(bs.WithValue("red")).Append("Red"),
			bs.Option(bs.WithValue("green")).Append("Green"),
			bs.Option(bs.WithValue("blue")).Append("Blue"),
			bs.Option(bs.WithValue("yellow")).Append("Yellow"),
			bs.Option(bs.WithValue("purple")).Append("Purple"),
			bs.Option(bs.WithValue("orange")).Append("Orange"),
		),
		bs.Para(bs.WithClass("form-text", "text-muted", "small")).Append("Hold Ctrl/Cmd to select multiple options."),
	)
	container.Append(multipleSelect)

	// Select with onChange
	changeSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Select with Change Handler"),
		bs.Select(bs.WithName("action")).OnChange(func(e Event) {
			println("Selection changed!")
		}).Append(
			bs.Option().Append("Select an action..."),
			bs.Option(bs.WithValue("view")).Append("View"),
			bs.Option(bs.WithValue("edit")).Append("Edit"),
			bs.Option(bs.WithValue("delete")).Append("Delete"),
		),
		bs.Para(bs.WithClass("form-text", "text-muted", "small")).Append("Check console for change events."),
	)
	container.Append(changeSelect)

	// Select sizes
	selectSizes := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Select Sizes"),
		bs.Container(bs.WithClass("mb-3")).Append(
			bs.Para(bs.WithClass("small", "text-muted", "mb-1")).Append("Large:"),
			bs.Select(bs.WithClass("form-select-lg")).Append(
				bs.Option().Append("Large select"),
				bs.Option(bs.WithValue("1")).Append("Option 1"),
			),
		),
		bs.Container(bs.WithClass("mb-3")).Append(
			bs.Para(bs.WithClass("small", "text-muted", "mb-1")).Append("Default:"),
			bs.Select().Append(
				bs.Option().Append("Default select"),
				bs.Option(bs.WithValue("1")).Append("Option 1"),
			),
		),
		bs.Container().Append(
			bs.Para(bs.WithClass("small", "text-muted", "mb-1")).Append("Small:"),
			bs.Select(bs.WithClass("form-select-sm")).Append(
				bs.Option().Append("Small select"),
				bs.Option(bs.WithValue("1")).Append("Option 1"),
			),
		),
	)
	container.Append(selectSizes)

	// Select in form example
	formWithSelect := bs.Container(bs.WithClass("mb-4")).Append(
		bs.Heading(6, bs.WithMargin(bs.BOTTOM, 2)).Append("Select in Form"),
		bs.Form(
			bs.WithAction("/submit"),
			bs.WithMethod("POST"),
			bs.WithClass("border", "p-3", "rounded", "bg-light"),
		).OnSubmit(func(e Event) {
			e.PreventDefault()
			println("Form with select submitted!")
		}).Append(
			bs.Container(bs.WithClass("mb-3")).Append(
				bs.Para(bs.WithClass("form-label", "mb-1")).Append("Select your plan:"),
				bs.Select(bs.WithName("plan"), bs.WithRequired()).Append(
					bs.Option(bs.WithValue("")).Append("Choose a plan..."),
					bs.Option(bs.WithValue("free")).Append("Free - $0/month"),
					bs.Option(bs.WithValue("pro")).Append("Pro - $29/month"),
					bs.Option(bs.WithValue("enterprise")).Append("Enterprise - Custom pricing"),
				),
			),
			bs.Button(bs.PRIMARY, bs.WithAttribute("type", "submit")).Append("Subscribe"),
		),
	)
	container.Append(formWithSelect)

	// Form attributes example
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Form Attributes"),
	)

	attrContainer := bs.Container(bs.WithClass("list-group", "mb-4")).Append(
		bs.Container(bs.WithClass("list-group-item")).Append(
			bs.Heading(6, bs.WithClass("mb-1")).Append(bs.Span(bs.WithClass("font-monospace")).Append("WithAction(url)")),
			bs.Para(bs.WithClass("mb-0", "text-muted")).Append("Sets the URL to submit form data to"),
		),
		bs.Container(bs.WithClass("list-group-item")).Append(
			bs.Heading(6, bs.WithClass("mb-1")).Append(bs.Span(bs.WithClass("font-monospace")).Append("WithMethod(method)")),
			bs.Para(bs.WithClass("mb-0", "text-muted")).Append("Sets HTTP method (GET or POST)"),
		),
		bs.Container(bs.WithClass("list-group-item")).Append(
			bs.Heading(6, bs.WithClass("mb-1")).Append(bs.Span(bs.WithClass("font-monospace")).Append("WithEnctype(enctype)")),
			bs.Para(bs.WithClass("mb-0", "text-muted")).Append("Sets encoding type for file uploads (multipart/form-data)"),
		),
		bs.Container(bs.WithClass("list-group-item")).Append(
			bs.Heading(6, bs.WithClass("mb-1")).Append(bs.Span(bs.WithClass("font-monospace")).Append("WithoutValidation()")),
			bs.Para(bs.WithClass("mb-0", "text-muted")).Append("Enables browser HTML5 validation (removes novalidate)"),
		),
		bs.Container(bs.WithClass("list-group-item")).Append(
			bs.Heading(6, bs.WithClass("mb-1")).Append(bs.Span(bs.WithClass("font-monospace")).Append("OnSubmit(callback)")),
			bs.Para(bs.WithClass("mb-0", "text-muted")).Append("Adds submit event listener (chainable method)"),
		),
	)

	container.Append(attrContainer)

	// Form validation classes
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Validation Classes"),
	)

	validationInfo := bs.Container(bs.WithClass("alert", "alert-info")).Append(
		bs.Heading(6).Append("Default Validation Behavior:"),
		bs.Para(bs.WithClass("mb-1")).Append(
			"• Forms include ", bs.Span(bs.WithClass("font-monospace")).Append("novalidate"),
			" attribute by default",
		),
		bs.Para(bs.WithClass("mb-1")).Append(
			"• Forms include ", bs.Span(bs.WithClass("font-monospace")).Append("needs-validation"),
			" class by default",
		),
		bs.Para(bs.WithClass("mb-1")).Append(
			"• Validation feedback is shown only after submit attempt",
		),
		bs.Para(bs.WithClass("mb-0")).Append(
			"• Add ", bs.Span(bs.WithClass("font-monospace")).Append("was-validated"),
			" class in OnSubmit handler to display feedback",
		),
	)

	container.Append(validationInfo)

	// Usage example code
	container.Append(
		bs.Heading(4, bs.WithMargin(bs.BOTTOM, 3), bs.WithMargin(bs.TOP, 4)).Append("Usage Example"),
	)

	usageExample := bs.Container(bs.WithClass("bg-dark", "text-light", "p-3", "rounded")).Append(
		bs.Para(bs.WithClass("font-monospace", "mb-1")).Append(
			`Form(`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-1", "ms-3")).Append(
			`WithAction("/api/login"),`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-1", "ms-3")).Append(
			`WithMethod("POST"),`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-1")).Append(
			`).OnSubmit(func(e Event) {`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-1", "ms-3")).Append(
			`e.PreventDefault()`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-1", "ms-3")).Append(
			`// Validate form and handle submission`,
		),
		bs.Para(bs.WithClass("font-monospace", "mb-0")).Append(
			`})`,
		),
	)

	container.Append(usageExample)

	return container
}
