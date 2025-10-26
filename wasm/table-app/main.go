package main

import (
	"fmt"
	"syscall/js"

	// Packages
	bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
	dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
	jsutil "github.com/djthorpe/go-wasmbuild/pkg/js"

	// Namespace imports
	. "github.com/djthorpe/go-wasmbuild"
)

// formatSalary converts a numeric salary to a formatted string with $ and commas
// e.g., 95000 -> "$95,000"
func formatSalary(salary float64) string {
	// Convert to integer
	salaryInt := int(salary)

	// Convert to string
	s := fmt.Sprintf("%d", salaryInt)

	// Add commas
	n := len(s)
	if n <= 3 {
		return "$" + s
	}

	// Build result with commas from right to left
	result := ""
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 {
			result += ","
		}
		result += string(s[i])
	}

	return "$" + result
}

// parseSalary converts a formatted salary string back to a plain number string
// e.g., "$95,000" -> "95000"
func parseSalary(formattedSalary string) string {
	result := ""
	for _, ch := range formattedSalary {
		if ch >= '0' && ch <= '9' {
			result += string(ch)
		}
	}
	return result
}

func main() {
	offcanvas := offcanvasComponent()
	table, tableContainer := tableComponent(offcanvas)

	// Create color dropdown items with individual click handlers
	colorItems := []Component{
		bs.NavItem("#", bs.WithID("default"), "Default"),
		bs.NavItem("#", bs.WithID("primary"), bs.PillBadge(bs.WithBackground(bs.PRIMARY), bs.WithClass("border")).Append(" "), " Primary"),
		bs.NavItem("#", bs.WithID("secondary"), bs.PillBadge(bs.WithBackground(bs.SECONDARY), bs.WithClass("border")).Append(" "), " Secondary"),
		bs.NavItem("#", bs.WithID("success"), bs.PillBadge(bs.WithBackground(bs.SUCCESS), bs.WithClass("border")).Append(" "), " Success"),
		bs.NavItem("#", bs.WithID("danger"), bs.PillBadge(bs.WithBackground(bs.DANGER), bs.WithClass("border")).Append(" "), " Danger"),
		bs.NavItem("#", bs.WithID("warning"), bs.PillBadge(bs.WithBackground(bs.WARNING), bs.WithClass("border")).Append(" "), " Warning"),
		bs.NavItem("#", bs.WithID("info"), bs.PillBadge(bs.WithBackground(bs.INFO), bs.WithClass("border")).Append(" "), " Info"),
		bs.NavItem("#", bs.WithID("light"), bs.PillBadge(bs.WithBackground(bs.LIGHT), bs.WithClass("border")).Append(" "), " Light"),
		bs.NavItem("#", bs.WithID("dark"), bs.PillBadge(bs.WithBackground(bs.DARK), bs.WithClass("border")).Append(" "), " Dark"),
	}

	// Helper function to update color ticks
	updateColorTicks := func(selectedID string) {
		for _, item := range colorItems {
			if comp := item.(Component); comp != nil {
				itemID := comp.ID()
				if itemID == "" {
					continue
				}

				elem := comp.Element()

				// Find if this is the selected item
				isSelected := itemID == selectedID

				// Check if tick already exists (look for span with tick)
				var tickSpan Element
				for child := elem.FirstChild(); child != nil; child = child.NextSibling() {
					if span, ok := child.(Element); ok && span.TagName() == "SPAN" {
						if span.TextContent() == "✓" {
							tickSpan = span
							break
						}
					}
				}

				if isSelected && tickSpan == nil {
					// Add tick
					newTick := dom.GetWindow().Document().CreateElement("SPAN")
					newTick.SetAttribute("style", "margin-left: auto; padding-left: 1rem;")
					newTick.AppendChild(dom.GetWindow().Document().CreateTextNode("✓"))
					elem.AppendChild(newTick)
					elem.SetAttribute("style", "display: flex; align-items: center;")
				} else if !isSelected && tickSpan != nil {
					// Remove tick
					elem.RemoveChild(tickSpan)
					elem.SetAttribute("style", "")
				} else if isSelected && tickSpan != nil {
					// Already has tick, ensure flex styling
					elem.SetAttribute("style", "display: flex; align-items: center;")
				}
			}
		}
	}

	// Initialize default color tick
	updateColorTicks("default")

	// Create navbar with brand header and color dropdown
	navbar := bs.NavBar(
		bs.WithClass("navbar-expand-lg"),
		bs.WithTheme(bs.DARK),
		bs.WithClass("bg-primary"),
	).Header(
		bs.Icon("people-fill"),
		" Employee Data Example",
	).Append(bs.NavDropdown("Color", colorItems...).AddEventListener("click", func(target Node) {
		if component := target.Component(); component != nil {
			colorID := component.ID()
			if colorID == "" {
				return
			}
			switch colorID {
			case "default":
				table.Apply(bs.WithColor(bs.TRANSPARENT))
				updateColorTicks(colorID)
			default:
				table.Apply(bs.WithColor(bs.Color(colorID)))
				updateColorTicks(colorID)
			}
		}
	}))

	// Add click handler to each color item
	/*
		for _, item := range colorItems {
			item.AddEventListener("click", func(target Node) {
				if component := target.Component(); component != nil {
					colorID := component.ID()
					if colorID == "" {
						return
					}
					switch colorID {
					case "default":
						table.Apply(bs.WithColor(bs.TRANSPARENT))
						updateColorTicks(colorID)
					default:
						table.Apply(bs.WithColor(bs.Color(colorID)))
						updateColorTicks(colorID)
					}
				}
			})
		}
	*/

	// Create styling dropdown with organized sections
	stylingItems := []Component{
		// Border section
		bs.NavDropdownHeader("Border"),
		bs.NavItem("#", bs.WithID("border-default"), "Default"),
		bs.NavItem("#", bs.WithID("border-bordered"), "Bordered"),
		bs.NavItem("#", bs.WithID("border-borderless"), "Borderless"),
		bs.NavDivider(false),

		// Size section
		bs.NavDropdownHeader("Size"),
		bs.NavItem("#", bs.WithID("size-default"), "Default"),
		bs.NavItem("#", bs.WithID("size-small"), "Small"),
		bs.NavDivider(false),

		// Group Divider section
		bs.NavDropdownHeader("Group Divider"),
		bs.NavItem("#", bs.WithID("divider-off"), "Off"),
		bs.NavItem("#", bs.WithID("divider-on"), "On"),
		bs.NavDivider(false),

		// Stripe section
		bs.NavDropdownHeader("Stripe"),
		bs.NavItem("#", bs.WithID("stripe-none"), "None"),
		bs.NavItem("#", bs.WithID("stripe-row"), "Row"),
		bs.NavItem("#", bs.WithID("stripe-column"), "Column"),
		bs.NavDivider(false),

		// Hover section
		bs.NavDropdownHeader("Hover"),
		bs.NavItem("#", bs.WithID("hover-off"), "Off"),
		bs.NavItem("#", bs.WithID("hover-on"), "On"),
	}

	// Helper function to update tick marks
	updateTicks := func(selectedID string, groupIDs []string) {
		for _, item := range stylingItems {
			if comp := item.(Component); comp != nil {
				itemID := comp.ID()
				if itemID == "" {
					continue
				}

				// Check if this item is in the current group
				isInGroup := false
				for _, gid := range groupIDs {
					if itemID == gid {
						isInGroup = true
						break
					}
				}

				if !isInGroup {
					continue
				}

				// Get the element and its text content
				elem := comp.Element()

				// Clear existing content and rebuild with or without tick
				for child := elem.FirstChild(); child != nil; {
					next := child.NextSibling()
					elem.RemoveChild(child)
					child = next
				}

				// Get the label text based on ID
				labelText := ""
				switch itemID {
				case "border-default", "size-default":
					labelText = "Default"
				case "border-bordered":
					labelText = "Bordered"
				case "border-borderless":
					labelText = "Borderless"
				case "size-small":
					labelText = "Small"
				case "divider-off", "hover-off":
					labelText = "Off"
				case "divider-on", "hover-on":
					labelText = "On"
				case "stripe-none":
					labelText = "None"
				case "stripe-row":
					labelText = "Row"
				case "stripe-column":
					labelText = "Column"
				}

				if itemID == selectedID {
					// Add text with tick on the right using flexbox
					elem.AppendChild(dom.GetWindow().Document().CreateTextNode(labelText))
					tickSpan := dom.GetWindow().Document().CreateElement("SPAN")
					tickSpan.SetAttribute("style", "margin-left: auto; padding-left: 1rem;")
					tickSpan.AppendChild(dom.GetWindow().Document().CreateTextNode("✓"))
					elem.AppendChild(tickSpan)
					elem.SetAttribute("style", "display: flex; align-items: center;")
				} else {
					elem.AppendChild(dom.GetWindow().Document().CreateTextNode(labelText))
					elem.SetAttribute("style", "")
				}
			}
		}
	}

	// Initialize default selections with ticks
	updateTicks("border-default", []string{"border-default", "border-bordered", "border-borderless"})
	updateTicks("size-default", []string{"size-default", "size-small"})
	updateTicks("divider-off", []string{"divider-off", "divider-on"})
	updateTicks("stripe-none", []string{"stripe-none", "stripe-row", "stripe-column"})
	updateTicks("hover-on", []string{"hover-off", "hover-on"})

	// Add click handler to each styling item
	for _, item := range stylingItems {
		item.AddEventListener("click", func(target Node) {
			if component := target.Component(); component != nil {
				itemID := component.ID()
				if itemID == "" {
					return
				}

				switch itemID {
				// Border options
				case "border-default":
					table.Element().ClassList().Remove("table-bordered", "table-borderless")
					updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})
				case "border-bordered":
					table.Element().ClassList().Remove("table-borderless")
					table.Element().ClassList().Add("table-bordered")
					updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})
				case "border-borderless":
					table.Element().ClassList().Remove("table-bordered")
					table.Element().ClassList().Add("table-borderless")
					updateTicks(itemID, []string{"border-default", "border-bordered", "border-borderless"})

				// Size options
				case "size-default":
					table.Element().ClassList().Remove("table-sm")
					updateTicks(itemID, []string{"size-default", "size-small"})
				case "size-small":
					table.Element().ClassList().Add("table-sm")
					updateTicks(itemID, []string{"size-default", "size-small"})

				// Group Divider options
				case "divider-off":
					table.Element().ClassList().Remove("table-group-divider")
					updateTicks(itemID, []string{"divider-off", "divider-on"})
				case "divider-on":
					table.Element().ClassList().Add("table-group-divider")
					updateTicks(itemID, []string{"divider-off", "divider-on"})

				// Stripe options
				case "stripe-none":
					table.Element().ClassList().Remove("table-striped", "table-striped-columns")
					updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})
				case "stripe-row":
					table.Element().ClassList().Remove("table-striped-columns")
					table.Element().ClassList().Add("table-striped")
					updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})
				case "stripe-column":
					table.Element().ClassList().Remove("table-striped")
					table.Element().ClassList().Add("table-striped-columns")
					updateTicks(itemID, []string{"stripe-none", "stripe-row", "stripe-column"})

				// Hover options
				case "hover-off":
					table.Element().ClassList().Remove("table-hover")
					updateTicks(itemID, []string{"hover-off", "hover-on"})
				case "hover-on":
					table.Element().ClassList().Add("table-hover")
					updateTicks(itemID, []string{"hover-off", "hover-on"})
				}
			}
		})
	}

	navbar.Append(bs.NavDropdown("Styling", stylingItems...))

	// Create the bootstrap app
	bs.New().Insert(
		navbar,
		bs.Container(
			bs.WithBreakpoint(bs.BreakpointLarge),
			bs.WithMargin(bs.TOP|bs.BOTTOM, 5),
		).Insert(
			tableContainer,
		),
		offcanvas,
	)

	// Run the application
	select {}
}

func offcanvasComponent() Component {
	// Create offcanvas for displaying employee details
	offcanvas := bs.Offcanvas(
		bs.WithID("employeeDetails"),
		bs.WithPosition(bs.END),
	)
	offcanvas.Header(
		bs.Heading(5).Append("Employee Details"),
		bs.CloseButton(bs.WithAttribute("data-bs-dismiss", "offcanvas")),
	)
	return offcanvas
}

func tableComponent(offcanvas Component) (Component, Component) {
	container := bs.Container()

	// Create employee model
	model := NewEmployeeModel()

	// Create a table with striped rows, hover, and animations
	table := bs.Table(
		bs.WithHover(), bs.WithAnimation(), bs.WithCursor(bs.CursorPointer),
	).Header("Name", "Position", "Salary", "Location")

	tableElem := table.Element()

	// Helper function to show employee details in offcanvas
	showEmployeeDetails := func(employee *Employee) {
		if employee == nil {
			return
		}

		// Find the offcanvas body element and clear it
		offcanvasElem := offcanvas.Element()

		// Find the body div (it's the second child after the header)
		var bodyElem Element
		for child := offcanvasElem.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok {
				if elem.ClassList().Contains("offcanvas-body") {
					bodyElem = elem
					break
				}
			}
		}

		if bodyElem != nil {
			// Clear existing content
			for child := bodyElem.FirstChild(); child != nil; {
				next := child.NextSibling()
				bodyElem.RemoveChild(child)
				child = next
			}

			// Create a card with employee details
			card := bs.Card()

			// Card header with employee name
			card.Header(employee.Name) // Create a form with employee details for the card body
			form := bs.Form()

			// Helper to create a form field with label and input
			createField := func(label, name, value string) Component {
				field := bs.Container(bs.WithClass("mb-3")).Append(
					bs.Label(label),
					bs.Input(name, bs.WithValue(value)),
				)
				return field
			}

			// Create form fields
			posDiv := createField("Position", "position", employee.Position)
			// Salary field - use NumberInput for numeric input
			// Use the actual numeric salary value from the model
			salaryValue := fmt.Sprintf("%.0f", employee.Salary)
			fmt.Printf("Creating salary input with value: %s (from %.0f)\n", salaryValue, employee.Salary)
			salaryInput := bs.NumberInput("salary")
			salaryInput.Value(salaryValue)
			salDiv := bs.Container(bs.WithClass("mb-3")).Append(
				bs.Label("Salary"),
				salaryInput,
			)
			locDiv := createField("Location", "location", employee.Location) // Append all fields to form
			form.Append(posDiv, salDiv, locDiv)                              // Add submit event listener to form
			form.AddEventListener("submit", func(node Node) {
				// Get the form element as js.Value and create jsutil.Form wrapper
				formElem := form.Element()
				formJS := formElem.(interface{ JSValue() js.Value }).JSValue()
				jsForm := jsutil.NewForm(formJS)

				// Get form data as a map
				formData := jsForm.FormData().Entries()

				fmt.Printf("Form submitted with data:\n")
				fmt.Printf("  Name: %s\n", employee.Name)
				fmt.Printf("  Position: %s\n", formData["position"])
				fmt.Printf("  Salary: %s\n", formData["salary"])
				fmt.Printf("  Location: %s\n", formData["location"])

				// Hide the offcanvas
				offcanvasJSValue := js.Global().Get("document").Call("getElementById", "employeeDetails")
				bsOffcanvas := js.Global().Get("bootstrap").Get("Offcanvas")
				instance := bsOffcanvas.Call("getOrCreateInstance", offcanvasJSValue)
				instance.Call("hide")
			}) // Add form to card body
			card.Append(form)

			// Add Save button to card footer that triggers form submit
			saveBtn := bs.Button(bs.PRIMARY).Append("Save")
			saveBtn.AddEventListener("click", func(node Node) {
				// Trigger form submit event
				formElem := form.Element()
				// Create and dispatch a submit event
				submitEvent := js.Global().Get("Event").New("submit")
				formElem.(interface {
					Call(string, ...interface{}) js.Value
				}).Call("dispatchEvent", submitEvent)
			})
			card.Footer(saveBtn)

			bodyElem.AppendChild(card.Element())
		}

		// Use Bootstrap's JavaScript API to properly show the offcanvas with backdrop
		// Get the raw JavaScript value of the element
		offcanvasJSValue := js.Global().Get("document").Call("getElementById", "employeeDetails")
		bsOffcanvas := js.Global().Get("bootstrap").Get("Offcanvas")
		instance := bsOffcanvas.Call("getOrCreateInstance", offcanvasJSValue)
		instance.Call("show")
	}

	// Action toolbar
	toolbar := bs.ButtonToolbar(bs.WithMargin(bs.END, 3), bs.WithMargin(bs.BOTTOM, 3))
	addBtn := bs.Button(bs.SUCCESS, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("plus-circle"),
		" Add",
	)
	removeBtn := bs.Button(bs.DANGER, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("dash-circle"),
		" Remove",
	)
	updateBtn := bs.Button(bs.WARNING, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("arrow-repeat"),
		" Update",
	)
	clearBtn := bs.Button(bs.SECONDARY, bs.WithSize(bs.SizeSmall)).Append(
		bs.Icon("x-circle"),
		" Clear",
	)

	// Helper function to update button states based on active rows
	updateButtonStates := func() {
		activeRows := table.Active()
		hasActiveRows := len(activeRows) > 0

		if hasActiveRows {
			// Enable all buttons - we have 4 buttons (0=Add, 1=Remove, 2=Update, 3=Clear)
			// When there are active rows, keep only Add enabled would be: toolbar.Disabled(1,2,3)
			// But we want ALL enabled, so disable NONE - pass empty/out of range indices
			// Actually, let's be explicit: we'll manually iterate or use a workaround
			// For now, let's disable button at index -1 (doesn't exist), forcing all to be enabled
			toolbar.Disabled(-1)
		} else {
			toolbar.Disabled(1, 2, 3) // Disable Remove, Update, Clear (indices 1, 2, 3)
		}
	}

	// Add mutation observer to table to automatically update button states
	// whenever row active states change
	observer := dom.GetWindow().NewMutationObserver(func() {
		updateButtonStates()
	})
	observer.Observe(tableElem, map[string]interface{}{
		"attributes":      true,
		"attributeFilter": []interface{}{"class"},
		"subtree":         true,
	})

	// Helper function to add click handler to a row
	addRowClickHandlerWithUpdate := func(row Component) {
		row.AddEventListener("click", func(node Node) {
			// node.Component() will walk up the DOM tree to find the row component
			// even if the click target is a TD element
			if rowComp := node.Component(); rowComp != nil {
				rowElem := rowComp.Element()

				// Find the row index by iterating through tbody children
				parent := rowElem.ParentNode()
				if parent == nil {
					return
				}

				rowIndex := -1
				currentIndex := 0
				for child := parent.FirstChild(); child != nil; child = child.NextSibling() {
					if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
						if elem == rowElem {
							rowIndex = currentIndex
							break
						}
						currentIndex++
					}
				}

				if rowIndex >= 0 {
					// Get the employee from the model using the row index
					employee := model.GetByIndex(rowIndex)
					if employee != nil {
						// Debug: print employee data
						fmt.Printf("Row clicked: index=%d, name=%s, salary=%.0f\n", rowIndex, employee.Name, employee.Salary)
						// Show employee details in offcanvas
						showEmployeeDetails(employee)
					} else {
						fmt.Printf("Row clicked: index=%d, employee is nil\n", rowIndex)
					}
				}

				// Toggle the active class directly on the element
				classList := rowElem.ClassList()
				if classList.Contains("table-active") {
					classList.Remove("table-active")
				} else {
					classList.Add("table-active")
				}
				// updateButtonStates() will be called automatically by mutation observer
			}
		})
	}

	addBtn.AddEventListener("click", func(node Node) {
		if model.Count() > 0 {
			// Get a random employee from the data
			randomEmployee := model.GetRandom()

			newRow := bs.Row(
				randomEmployee.Name,
				randomEmployee.Position,
				formatSalary(randomEmployee.Salary),
				randomEmployee.Location,
			)
			addRowClickHandlerWithUpdate(newRow)

			// Get active rows and insert before the first active row, or at the beginning
			activeRows := table.Active()
			if len(activeRows) > 0 {
				// Insert before the first active row
				table.InsertBefore(activeRows[0], newRow)
			} else {
				// Insert at the beginning (before index 0)
				table.InsertBefore(0, newRow)
			}
		}
	})
	removeBtn.AddEventListener("click", func(node Node) {
		activeRows := table.Active()
		if len(activeRows) > 0 {
			// Delete all active rows, starting from the highest index to avoid index shifting issues
			for i := len(activeRows) - 1; i >= 0; i-- {
				table.Delete(activeRows[i])
			}
			// updateButtonStates() will be called automatically by mutation observer
		}
	})
	updateBtn.AddEventListener("click", func(node Node) {
		activeRows := table.Active()
		if len(activeRows) > 0 && model.Count() > 0 {
			// Update all active rows with random employee data
			for _, rowIndex := range activeRows {
				// Get a random employee from the data
				randomEmployee := model.GetRandom()

				newRow := bs.Row(
					randomEmployee.Name,
					randomEmployee.Position,
					formatSalary(randomEmployee.Salary),
					randomEmployee.Location,
				).WithActive()
				addRowClickHandlerWithUpdate(newRow)
				table.Replace(rowIndex, newRow)
			}
		}
	})
	clearBtn.AddEventListener("click", func(node Node) {
		table.Active(-1)
		// updateButtonStates() will be called automatically by mutation observer
	})

	toolbar.Append(addBtn, removeBtn, updateBtn, clearBtn)

	// Initialize button states after buttons are appended
	updateButtonStates()

	container.Append(toolbar)

	// Create a loading message
	loadingDiv := bs.Container().Insert(
		"Loading employee data...",
	)
	container.Append(loadingDiv)

	// Function to calculate and update the total salary
	updateTotalSalary := func() {
		tbody := tableElem.FirstChild().(Element)
		totalSalary := 0.0

		// Iterate through all rows in tbody
		for child := tbody.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok && elem.TagName() == "TR" {
				// Get the third cell (salary column, index 2)
				cells := elem.Children()
				if len(cells) >= 3 {
					salaryCell := cells[2]
					salaryText := salaryCell.TextContent()

					// Parse salary (remove $ and commas)
					cleanSalary := ""
					for _, ch := range salaryText {
						if ch != '$' && ch != ',' {
							cleanSalary += string(ch)
						}
					}

					// Convert to float
					salaryVal := js.Global().Get("parseFloat").Invoke(cleanSalary).Float()
					if !js.Global().Get("isNaN").Invoke(salaryVal).Bool() {
						totalSalary += salaryVal
					}
				}
			}
		}

		// Update footer with new total
		totalFormatted := js.Global().Get("Intl").Get("NumberFormat").New("en-US", map[string]interface{}{
			"style":    "currency",
			"currency": "USD",
		}).Call("format", totalSalary).String()

		// Find and update the footer
		tfoot := tableElem.LastChild()
		if tfootElem, ok := tfoot.(Element); ok && tfootElem.TagName() == "TFOOT" {
			tfootRow := tfootElem.FirstChild()
			if tfootRowElem, ok := tfootRow.(Element); ok {
				cells := tfootRowElem.Children()
				if len(cells) >= 4 {
					lastCell := cells[3]
					for child := lastCell.FirstChild(); child != nil; {
						next := child.NextSibling()
						lastCell.RemoveChild(child)
						child = next
					}
					lastCell.AppendChild(dom.GetWindow().Document().CreateTextNode(totalFormatted))
				}
			}
		}
	}

	// Add mutation observer to watch for table changes
	tableElem.AddEventListener("DOMSubtreeModified", func(node Node) {
		updateTotalSalary()
	})

	// Create pagination component
	pagination := bs.Pagination("Employee table navigation")

	// Declare updatePagination first so it can be called recursively
	var updatePagination func()

	// Function to refresh the table based on current model offset/limit
	refreshTable := func() {
		// Clear existing table rows (keep header)
		tbody := tableElem.FirstChild()
		if tbodyElem, ok := tbody.(Element); ok && tbodyElem.TagName() == "TBODY" {
			for child := tbodyElem.FirstChild(); child != nil; {
				next := child.NextSibling()
				tbodyElem.RemoveChild(child)
				child = next
			}
		}

		// Populate table with current page of employee data
		for _, employee := range model.GetAll() {
			row := bs.Row(employee.Name, employee.Position, formatSalary(employee.Salary), employee.Location)
			addRowClickHandlerWithUpdate(row)
			table.Append(row)
		}

		updateTotalSalary()
	}

	// Function to update pagination based on model state
	updatePagination = func() {
		// Clear existing pagination items by removing and recreating
		paginationElem := pagination.Element()

		// Find and remove the UL element
		for child := paginationElem.FirstChild(); child != nil; {
			next := child.NextSibling()
			paginationElem.RemoveChild(child)
			child = next
		}

		// Recreate UL element with centered alignment
		doc := dom.GetWindow().Document()
		list := doc.CreateElement("UL")
		list.SetAttribute("class", "pagination justify-content-center")
		paginationElem.AppendChild(list)

		totalPages := (model.Count() + model.Limit() - 1) / model.Limit()
		if totalPages == 0 {
			totalPages = 1
		}
		currentPage := (model.Offset() / model.Limit()) + 1

		// Helper to create a page link element
		createPageLink := func(text string, pageNum int, disabled bool, active bool) Element {
			doc := dom.GetWindow().Document()
			link := doc.CreateElement("A")
			link.SetAttribute("class", "page-link")
			link.SetAttribute("href", "javascript:void(0)")
			if disabled {
				link.SetAttribute("tabindex", "-1")
				link.SetAttribute("aria-disabled", "true")
			}
			if active {
				link.SetAttribute("aria-current", "page")
			}
			link.AppendChild(doc.CreateTextNode(text))

			if !disabled {
				link.AddEventListener("click", func(node Node) {
					model.SetOffsetLimit((pageNum-1)*model.Limit(), model.Limit())
					refreshTable()
					updatePagination()
				})
			}

			// Create LI wrapper
			li := doc.CreateElement("LI")
			li.SetAttribute("class", "page-item")
			if disabled {
				li.ClassList().Add("disabled")
			}
			if active {
				li.ClassList().Add("active")
			}
			li.AppendChild(link)

			return li
		}

		// Add Previous button
		list.AppendChild(createPageLink("Previous", currentPage-1, currentPage == 1, false))

		// Add page numbers (show max 5 pages)
		startPage := currentPage - 2
		if startPage < 1 {
			startPage = 1
		}
		endPage := startPage + 4
		if endPage > totalPages {
			endPage = totalPages
			startPage = endPage - 4
			if startPage < 1 {
				startPage = 1
			}
		}

		for page := startPage; page <= endPage; page++ {
			pageNum := page
			pageText := ""
			// Convert int to string
			pageText = js.Global().Get("String").Invoke(pageNum).String()
			list.AppendChild(createPageLink(pageText, pageNum, false, pageNum == currentPage))
		}

		// Add Next button
		list.AppendChild(createPageLink("Next", currentPage+1, currentPage == totalPages, false))
	}

	// Create an interval to check for data
	model.WaitForData(func() {
		container.Element().RemoveChild(loadingDiv.Element())

		// Populate table with employee data
		for _, employee := range model.GetAll() {
			row := bs.Row(employee.Name, employee.Position, formatSalary(employee.Salary), employee.Location)
			addRowClickHandlerWithUpdate(row)
			table.Append(row)
		}

		table.Footer("", "", "Total:", "$0.00")
		updateTotalSalary()
		container.Append(table)

		// Add pagination below the table
		updatePagination()
		container.Append(pagination)
	})

	return table, container
}
