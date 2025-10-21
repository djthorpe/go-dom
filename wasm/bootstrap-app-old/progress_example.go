package main

import (
	"github.com/djthorpe/go-dom"
	"github.com/djthorpe/go-dom/pkg/bs5"
)

// AddProgressExamples adds progress bar component examples to the app
func AddProgressExamples(app *bs5.App) dom.Element {
	container := app.Container()

	// Collection to store all progress bars and track which should show labels
	var allProgressBars []*bs5.Progress
	var progressWithLabels map[*bs5.Progress]bool = make(map[*bs5.Progress]bool)

	// Control Panel Card
	controlCard := app.Card()
	controlCard.AddClass("mb-4")
	controlCard.Header(app.H4(app.CreateTextNode("Control All Progress Bars")).Element)
	controlCardBody := controlCard.Body()

	descP := app.CreateElement("p")
	descP.AppendChild(app.CreateTextNode("Click a button to set all progress bars to that value:"))
	controlCardBody.Element.AppendChild(descP)

	// Create button group - will add click handlers later after all progress bars are created
	buttonDiv := app.CreateElement("div")
	buttonDiv.AddClass("mt-2")
	controlCardBody.Element.AppendChild(buttonDiv)
	container.AppendChild(controlCard.Element)

	// Basic Progress Bars (no labels)
	basicCard := app.Card()
	basicCard.AddClass("mb-4")
	basicCard.Header(app.H4(app.CreateTextNode("Basic Progress")).Element)
	basicCardBody := basicCard.Body()

	basicProgress := app.Progress(25, "Basic progress")
	allProgressBars = append(allProgressBars, basicProgress)
	progressWithLabels[basicProgress] = false
	basicCardBody.Element.AppendChild(basicProgress.Element)
	container.AppendChild(basicCard.Element)

	// Different Heights (no labels)
	heightCard := app.Card()
	heightCard.AddClass("mb-4")
	heightCard.Header(app.H4(app.CreateTextNode("Different Heights")).Element)
	heightCardBody := heightCard.Body()

	smallProgress := app.Progress(25, "Small height").SetHeight("5px")
	allProgressBars = append(allProgressBars, smallProgress)
	progressWithLabels[smallProgress] = false
	heightCardBody.Element.AppendChild(smallProgress.Element)
	heightCardBody.Element.AppendChild(app.CreateElement("br"))

	defaultProgress := app.Progress(50, "Default height")
	allProgressBars = append(allProgressBars, defaultProgress)
	progressWithLabels[defaultProgress] = false
	heightCardBody.Element.AppendChild(defaultProgress.Element)
	heightCardBody.Element.AppendChild(app.CreateElement("br"))

	largeProgress := app.Progress(75, "Large height").SetHeight("30px")
	allProgressBars = append(allProgressBars, largeProgress)
	progressWithLabels[largeProgress] = false
	heightCardBody.Element.AppendChild(largeProgress.Element)
	container.AppendChild(heightCard.Element)

	// Colored Progress Bars (with labels)
	colorCard := app.Card()
	colorCard.AddClass("mb-4")
	colorCard.Header(app.H4(app.CreateTextNode("Colored Progress Bars")).Element)
	colorCardBody := colorCard.Body()

	successProgress := app.Progress(25, "Success").SetColor(bs5.ColorSuccess).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, successProgress)
	progressWithLabels[successProgress] = true
	colorCardBody.Element.AppendChild(successProgress.Element)
	colorCardBody.Element.AppendChild(app.CreateElement("br"))

	infoProgress := app.Progress(50, "Info").SetColor(bs5.ColorInfo).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, infoProgress)
	progressWithLabels[infoProgress] = true
	colorCardBody.Element.AppendChild(infoProgress.Element)
	colorCardBody.Element.AppendChild(app.CreateElement("br"))

	warningProgress := app.Progress(75, "Warning").SetColor(bs5.ColorWarning).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, warningProgress)
	progressWithLabels[warningProgress] = true
	colorCardBody.Element.AppendChild(warningProgress.Element)
	colorCardBody.Element.AppendChild(app.CreateElement("br"))

	dangerProgress := app.Progress(100, "Danger").SetColor(bs5.ColorDanger).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, dangerProgress)
	progressWithLabels[dangerProgress] = true
	colorCardBody.Element.AppendChild(dangerProgress.Element)
	container.AppendChild(colorCard.Element)

	// Colored with Text Background (with labels)
	textBgCard := app.Card()
	textBgCard.AddClass("mb-4")
	textBgCard.Header(app.H4(app.CreateTextNode("Colored with Text Background")).Element)
	textBgCardBody := textBgCard.Body()

	textBgSuccess := app.Progress(25, "Success").SetTextColor(bs5.ColorSuccess).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, textBgSuccess)
	progressWithLabels[textBgSuccess] = true
	textBgCardBody.Element.AppendChild(textBgSuccess.Element)
	textBgCardBody.Element.AppendChild(app.CreateElement("br"))

	textBgInfo := app.Progress(50, "Info").SetTextColor(bs5.ColorInfo).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, textBgInfo)
	progressWithLabels[textBgInfo] = true
	textBgCardBody.Element.AppendChild(textBgInfo.Element)
	textBgCardBody.Element.AppendChild(app.CreateElement("br"))

	textBgWarning := app.Progress(75, "Warning").SetTextColor(bs5.ColorWarning).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, textBgWarning)
	progressWithLabels[textBgWarning] = true
	textBgCardBody.Element.AppendChild(textBgWarning.Element)
	textBgCardBody.Element.AppendChild(app.CreateElement("br"))

	textBgDanger := app.Progress(100, "Danger").SetTextColor(bs5.ColorDanger).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, textBgDanger)
	progressWithLabels[textBgDanger] = true
	textBgCardBody.Element.AppendChild(textBgDanger.Element)
	container.AppendChild(textBgCard.Element)

	// Striped Progress Bars (with labels)
	stripedCard := app.Card()
	stripedCard.AddClass("mb-4")
	stripedCard.Header(app.H4(app.CreateTextNode("Striped Progress Bars")).Element)
	stripedCardBody := stripedCard.Body()

	stripedProgress1 := app.Progress(25, "Striped").SetStriped(true).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, stripedProgress1)
	progressWithLabels[stripedProgress1] = true
	stripedCardBody.Element.AppendChild(stripedProgress1.Element)
	stripedCardBody.Element.AppendChild(app.CreateElement("br"))

	stripedProgress2 := app.Progress(50, "Striped success").SetColor(bs5.ColorSuccess).SetStriped(true).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, stripedProgress2)
	progressWithLabels[stripedProgress2] = true
	stripedCardBody.Element.AppendChild(stripedProgress2.Element)
	stripedCardBody.Element.AppendChild(app.CreateElement("br"))

	stripedProgress3 := app.Progress(75, "Striped info").SetColor(bs5.ColorInfo).SetStriped(true).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, stripedProgress3)
	progressWithLabels[stripedProgress3] = true
	stripedCardBody.Element.AppendChild(stripedProgress3.Element)
	container.AppendChild(stripedCard.Element)

	// Animated Progress Bar (with labels)
	animatedCard := app.Card()
	animatedCard.AddClass("mb-4")
	animatedCard.Header(app.H4(app.CreateTextNode("Animated Progress Bar")).Element)
	animatedCardBody := animatedCard.Body()

	animatedProgress := app.Progress(75, "Animated progress").SetAnimated(true).ShowLabel(true, app)
	allProgressBars = append(allProgressBars, animatedProgress)
	progressWithLabels[animatedProgress] = true
	animatedCardBody.Element.AppendChild(animatedProgress.Element)
	container.AppendChild(animatedCard.Element)

	// Stacked Progress Bars (no labels)
	// Note: These are handled specially - they're not in the main progress bar list
	stackedCard := app.Card()
	stackedCard.AddClass("mb-4")
	stackedCard.Header(app.H4(app.CreateTextNode("Stacked Progress Bars")).Element)
	stackedCardBody := stackedCard.Body()

	// Create individual progress segments for the stack
	// Initial proportions: 15%, 30%, 20% = 65% total, so ratios are 15/65, 30/65, 20/65
	stackedProgress1 := app.ProgressStackedSegment(15, "Segment one")
	stackedProgress2 := app.ProgressStackedSegment(30, "Segment two").SetColor(bs5.ColorSuccess)
	stackedProgress3 := app.ProgressStackedSegment(20, "Segment three").SetColor(bs5.ColorInfo)

	// Store stacked segments separately to handle them differently
	stackedSegments := []*bs5.Progress{stackedProgress1, stackedProgress2, stackedProgress3}
	// Calculate proportions as fractions (15:30:20 ratio, total 65)
	totalProportion := 15.0 + 30.0 + 20.0
	stackedProportions := []float64{15.0 / totalProportion, 30.0 / totalProportion, 20.0 / totalProportion}

	stackedCardBody.Element.AppendChild(app.ProgressStacked(
		stackedProgress1.Element,
		stackedProgress2.Element,
		stackedProgress3.Element,
	).Element)
	container.AppendChild(stackedCard.Element)

	// Now add the control buttons at the top
	values := []int{0, 10, 25, 50, 66, 100}
	for i, val := range values {
		v := val // Capture value for closure
		var btnLabel string
		switch v {
		case 0:
			btnLabel = "0%"
		case 10:
			btnLabel = "10%"
		case 25:
			btnLabel = "25%"
		case 50:
			btnLabel = "50%"
		case 66:
			btnLabel = "66%"
		case 100:
			btnLabel = "100%"
		}

		var btnColor bs5.ColorVariant
		if v == 0 {
			btnColor = bs5.ColorSecondary
		} else if v == 100 {
			btnColor = bs5.ColorSuccess
		} else {
			btnColor = bs5.ColorPrimary
		}

		btn := app.Button(btnColor, app.CreateTextNode(btnLabel))
		btn.SetSize(bs5.ButtonSizeSmall)
		if i < len(values)-1 {
			btn.AddClass("me-2")
		}
		btn.AddEventListener("click", func(target dom.Node) {
			// Update all progress bars
			for _, pb := range allProgressBars {
				pb.SetValue(v)
				// Only show label if this progress bar should have one
				if progressWithLabels[pb] {
					pb.ShowLabel(true, app)
				} else {
					pb.ShowLabel(false, app)
				}
			}

			// Update stacked segments proportionally
			for i, segment := range stackedSegments {
				// Calculate using float then round to int
				proportionalValue := int(float64(v)*stackedProportions[i] + 0.5)
				segment.SetValue(proportionalValue)
			}
		})
		buttonDiv.AppendChild(btn.Element)
	}

	return container.Element
}
