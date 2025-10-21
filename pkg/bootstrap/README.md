# Bootstrap Package

A Go package for building Bootstrap 5.3-compliant web applications using WebAssembly (WASM). This package provides a type-safe, composable API for creating Bootstrap components with proper CSS classes and attributes.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Core Concepts](#core-concepts)
- [Components](#components)
- [Configuration Options](#configuration-options)
- [Creating Custom Components](#creating-custom-components)
- [Examples](#examples)

## Installation

```bash
go get github.com/djthorpe/go-wasmbuild/pkg/bootstrap
```

## Quick Start

```go
package main

import (
    bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
)

func main() {
    // Create a new Bootstrap application
    app := bs.New()
    
    // Add components to the app
    app.Append(
        bs.Container().Append(
            bs.Heading(1, bs.WithClass("text-center")).Append("Welcome"),
            bs.Alert(bs.WithColor(bs.PRIMARY)).Append("Hello, Bootstrap!"),
            bs.Button(bs.PRIMARY, bs.WithSize(bs.SizeLarge)).Append("Click Me"),
        ),
    )
    
    // Application runs indefinitely
    select {}
}
```

## Core Concepts

### Component Interface

All Bootstrap components implement the `Component` interface:

```go
type Component interface {
    Element() Element          // Returns the underlying DOM element
    Append(children ...any) Component  // Appends children and returns self for chaining
}
```

### Options System

Components are configured using functional options:

```go
// WithClass adds CSS classes
bs.Badge(bs.WithClass("ms-2", "rounded-pill"))

// WithColor sets component color
bs.Button(bs.PRIMARY, bs.WithColor(bs.DANGER))

// WithSize sets component size
bs.Button(bs.PRIMARY, bs.WithSize(bs.SizeLarge))

// WithAttribute sets custom HTML attributes
bs.Container(bs.WithAttribute("id", "main-content"))
```

### Method Chaining

All components support method chaining for fluent API design:

```go
container := bs.Container(bs.WithClass("text-center")).
    Append("Welcome to our site").
    Append(bs.Button(bs.PRIMARY).Append("Get Started"))
```

## Components

### Alert

Create Bootstrap alert components for displaying messages.

```go
// Basic alert
alert := bs.Alert(bs.WithColor(bs.PRIMARY)).Append("Primary alert message")

// Alert with multiple colors
successAlert := bs.Alert(bs.WithColor(bs.SUCCESS)).Append("Operation successful!")
dangerAlert := bs.Alert(bs.WithColor(bs.DANGER)).Append("Error occurred!")
```

**HTML Output:**

```html
<div class="alert alert-primary" role="alert">Primary alert message</div>
```

### Badge

Create Bootstrap badge components for labels and notifications.

```go
// Basic badge
badge := bs.Badge(bs.WithColor(bs.SECONDARY)).Append("New")

// Pill badge (rounded)
pillBadge := bs.PillBadge(bs.WithColor(bs.PRIMARY)).Append("Pro")
```

**HTML Output:**

```html
<span class="badge text-bg-secondary">New</span>
<span class="badge text-bg-primary rounded-pill">Pro</span>
```

### Button

Create Bootstrap button components with various styles and sizes.

```go
// Basic button
btn := bs.Button(bs.PRIMARY).Append("Primary Button")

// Outline button
outlineBtn := bs.OutlineButton(bs.SECONDARY).Append("Outline Button")

// Button with size
largeBnt := bs.Button(bs.SUCCESS, bs.WithSize(bs.SizeLarge)).Append("Large Button")
```

**HTML Output:**

```html
<button class="btn btn-primary" type="button">Primary Button</button>
<button class="btn btn-outline-secondary" type="button">Outline Button</button>
```

### Card

Create Bootstrap card components for flexible content containers.

```go
// Basic card
card := bs.Card().Append(
    bs.CardHeader().Append("Card Header"),
    bs.CardBody().Append(
        bs.CardTitle().Append("Card Title"),
        bs.CardText().Append("Card content goes here."),
        bs.Button(bs.PRIMARY).Append("Go somewhere"),
    ),
)
```

**HTML Output:**

```html
<div class="card">
    <div class="card-header">Card Header</div>
    <div class="card-body">
        <h5 class="card-title">Card Title</h5>
        <p class="card-text">Card content goes here.</p>
        <button class="btn btn-primary" type="button">Go somewhere</button>
    </div>
</div>
```

### Container

Create Bootstrap container components for responsive layouts.

```go
// Default container
container := bs.Container().Append("Default container content")

// Fluid container
fluidContainer := bs.Container(bs.WithBreakpoint(bs.BreakpointFluid)).Append(
    "Full-width container content",
)

// Responsive containers
smContainer := bs.Container(bs.WithBreakpoint(bs.BreakpointSmall))
```

**HTML Output:**

```html
<div class="container">Default container content</div>
<div class="container-fluid">Full-width container content</div>
<div class="container-sm">Small container content</div>
```

### Heading

Create Bootstrap heading components (H1-H6) with proper styling.

```go
// Basic headings
h1 := bs.Heading(1).Append("Main Title")
h2 := bs.Heading(2).Append("Subtitle")
h3 := bs.Heading(3, bs.WithClass("text-muted")).Append("Section Title")

// Display headings
displayHeading := bs.Heading(1, bs.WithClass("display-4")).Append("Large Display")
```

**HTML Output:**

```html
<h1>Main Title</h1>
<h2>Subtitle</h2>
<h3 class="text-muted">Section Title</h3>
<h1 class="display-4">Large Display</h1>
```

### Icon

Create Bootstrap Icons for visual elements.

```go
// Basic icon
icon := bs.Icon("bi-heart").Append()

// Icon with color
coloredIcon := bs.Icon("bi-star", bs.WithColor(bs.WARNING)).Append()

// Icon as button content
iconButton := bs.Button(bs.PRIMARY).Append(
    bs.Icon("bi-download").Append(),
    " Download",
)
```

**HTML Output:**

```html
<i class="bi bi-heart"></i>
<i class="bi bi-star text-warning"></i>
<button class="btn btn-primary" type="button">
    <i class="bi bi-download"></i> Download
</button>
```

### Image

Create Bootstrap image components with responsive behavior.

```go
// Basic image
img := bs.Image("photo.jpg", bs.WithClass("img-fluid"))

// Responsive image
responsiveImg := bs.Image("hero.jpg", 
    bs.WithClass("img-fluid", "rounded"),
    bs.WithAttribute("alt", "Hero image"),
)
```

**HTML Output:**

```html
<img src="photo.jpg" class="img-fluid">
<img src="hero.jpg" class="img-fluid rounded" alt="Hero image">
```

### Link

Create Bootstrap link components with proper styling.

```go
// Basic link
link := bs.Link("/about", bs.WithColor(bs.PRIMARY)).Append("About Us")

// External link
externalLink := bs.Link("https://example.com",
    bs.WithAttribute("target", "_blank"),
    bs.WithAttribute("rel", "noopener"),
).Append("External Site")
```

**HTML Output:**

```html
<a href="/about" class="link-primary">About Us</a>
<a href="https://example.com" target="_blank" rel="noopener">External Site</a>
```

### Nav

Create Bootstrap navigation components.

```go
// Basic nav
nav := bs.Nav(bs.WithClass("nav-pills")).Append(
    bs.Link("#home", bs.WithClass("nav-link", "active")).Append("Home"),
    bs.Link("#about", bs.WithClass("nav-link")).Append("About"),
    bs.Link("#contact", bs.WithClass("nav-link")).Append("Contact"),
)
```

**HTML Output:**

```html
<nav class="nav nav-pills">
    <a href="#home" class="nav-link active">Home</a>
    <a href="#about" class="nav-link">About</a>
    <a href="#contact" class="nav-link">Contact</a>
</nav>
```

### Para (Paragraph)

Create Bootstrap paragraph components.

```go
// Basic paragraph
para := bs.Para().Append("This is a paragraph of text.")

// Styled paragraph
styledPara := bs.Para(
    bs.WithClass("lead", "text-muted"),
).Append("This is a lead paragraph with muted text.")
```

**HTML Output:**

```html
<p>This is a paragraph of text.</p>
<p class="lead text-muted">This is a lead paragraph with muted text.</p>
```

### Rule

Create Bootstrap horizontal and vertical rules for visual separation.

```go
// Horizontal rule
hr := bs.Rule()

// Vertical rule (for use in flex containers)
vr := bs.VerticalRule()
```

**HTML Output:**

```html
<hr>
<div class="vr"></div>
```

### Span

Create Bootstrap span components for inline elements.

```go
// Basic span
span := bs.Span().Append("Inline text")

// Styled span
styledSpan := bs.Span(
    bs.WithClass("badge", "bg-secondary"),
).Append("Badge Text")
```

**HTML Output:**

```html
<span>Inline text</span>
<span class="badge bg-secondary">Badge Text</span>
```

## Configuration Options

The Bootstrap package uses a functional options pattern for configuring components:

### Color Options

```go
// Available colors
bs.PRIMARY, bs.SECONDARY, bs.SUCCESS, bs.DANGER, bs.WARNING, 
bs.INFO, bs.LIGHT, bs.DARK, bs.WHITE, bs.BLACK

// Subtle color variants
bs.PRIMARY_SUBTLE, bs.SECONDARY_SUBTLE, bs.SUCCESS_SUBTLE,
bs.DANGER_SUBTLE, bs.WARNING_SUBTLE, bs.INFO_SUBTLE,
bs.LIGHT_SUBTLE, bs.DARK_SUBTLE

// Usage
bs.Button(bs.PRIMARY)  // btn-primary
bs.Alert(bs.WithColor(bs.DANGER))  // alert-danger
bs.Badge(bs.WithColor(bs.SUCCESS))  // text-bg-success
```

### Size Options

```go
// Available sizes
bs.SizeSmall    // sm
bs.SizeDefault  // (no class added)
bs.SizeLarge    // lg

// Usage
bs.Button(bs.PRIMARY, bs.WithSize(bs.SizeLarge))  // btn btn-primary btn-lg
```

### Position Options

```go
// Available positions (can be combined with |)
bs.TOP, bs.BOTTOM, bs.START, bs.END, bs.ALL
bs.CENTER, bs.MIDDLE  // For flex alignment

// Usage for margins and padding
bs.WithMargin(bs.TOP|bs.BOTTOM, 3)    // my-3
bs.WithPadding(bs.ALL, 2)             // p-2

// Usage for borders
bs.WithBorder(bs.TOP, bs.PRIMARY)     // border-top border-primary

// Usage for flex alignment
bs.WithFlex(bs.CENTER)                // d-flex align-items-center
```

### Custom Options

```go
// Add custom CSS classes
bs.WithClass("custom-class", "another-class")

// Set HTML attributes
bs.WithAttribute("id", "unique-id")
bs.WithAttribute("data-bs-toggle", "modal")

// Set aria-label for accessibility
bs.WithAriaLabel("Close button")
```

## Creating Custom Components

You can create your own Bootstrap-compliant components by following these patterns:

### Basic Component Structure

```go
package main

import (
    "strings"
    
    dom "github.com/djthorpe/go-wasmbuild/pkg/dom"
    bs "github.com/djthorpe/go-wasmbuild/pkg/bootstrap"
    . "github.com/djthorpe/go-wasmbuild"
)

// CustomComponent represents a custom Bootstrap component
type CustomComponent struct {
    component
}

// NewCustomComponent creates a new custom component
func NewCustomComponent(opt ...bs.Opt) *CustomComponent {
    // Create the root element
    root := dom.GetWindow().Document().CreateElement("DIV")
    
    // Apply options with default classes
    if opts, err := bs.NewOpts("custom", bs.WithClass("custom-component")); err != nil {
        panic(err)
    } else if err := opts.apply(opt...); err != nil {
        panic(err)
    } else {
        // Set class list and attributes
        classes := opts.classList.Values()
        if len(classes) > 0 {
            root.SetAttribute("class", strings.Join(classes, " "))
        }
    }
    
    return &CustomComponent{
        component: component{
            name: "custom",
            root: root,
        },
    }
}
```

## Examples

### Complete Page Layout

```go
func CreatePage() Component {
    return bs.Container().Append(
        // Header
        bs.Container(bs.WithClass("bg-primary", "text-white", "p-4", "mb-4")).Append(
            bs.Heading(1, bs.WithClass("mb-0")).Append("My Bootstrap App"),
        ),
        
        // Main content
        bs.Container(bs.WithClass("row")).Append(
            // Sidebar
            bs.Container(bs.WithClass("col-md-3", "mb-4")).Append(
                bs.Card().Append(
                    bs.CardHeader().Append("Navigation"),
                    bs.CardBody().Append(
                        bs.Nav(bs.WithClass("nav-pills", "flex-column")).Append(
                            bs.Link("#home", bs.WithClass("nav-link", "active")).Append("Home"),
                            bs.Link("#about", bs.WithClass("nav-link")).Append("About"),
                        ),
                    ),
                ),
            ),
            
            // Main content area
            bs.Container(bs.WithClass("col-md-9")).Append(
                bs.Alert(bs.WithColor(bs.SUCCESS)).Append("Welcome!"),
                bs.Card().Append(
                    bs.CardHeader().Append("Dashboard"),
                    bs.CardBody().Append(
                        bs.Heading(3).Append("Statistics"),
                        bs.Para().Append("Here are some key metrics:"),
                        bs.ButtonGroup().Append(
                            bs.Button(bs.PRIMARY).Append("View All"),
                            bs.Button(bs.SECONDARY).Append("Filter"),
                        ),
                    ),
                ),
            ),
        ),
        
        // Footer
        bs.Rule(bs.WithClass("my-4")),
        bs.Container(bs.WithClass("text-center", "text-muted")).Append(
            bs.Para().Append("© 2025 My Bootstrap App. Built with Go and WASM."),
        ),
    )
}
```

## Best Practices

1. **Use Method Chaining**: Components support fluent method chaining for cleaner code
2. **Combine Options**: Use multiple options to configure components precisely
3. **Semantic HTML**: Choose appropriate components for semantic meaning
4. **Accessibility**: Use `WithAriaLabel` and proper attributes for screen readers
5. **Responsive Design**: Utilize Bootstrap's grid system and responsive classes
6. **Color Consistency**: Stick to Bootstrap's color palette for consistent theming
7. **Performance**: Components are lightweight and generate minimal DOM overhead

## Browser Compatibility

This package generates Bootstrap 5.3-compliant HTML and works in all modern browsers that support WebAssembly:

- Chrome 57+
- Firefox 52+
- Safari 11+
- Edge 16+

For complete Bootstrap functionality, include the Bootstrap CSS and JavaScript files in your HTML template.

## Contributing

Contributions are welcome! Please ensure that new components:

1. Follow the existing component patterns
2. Include comprehensive tests
3. Support the full options system
4. Generate valid Bootstrap 5.3 HTML
5. Include documentation and examples

## License

This package is part of the go-wasmbuild project. See the main project for license information.
