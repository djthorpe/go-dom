# Bootstrap 5 Demo Application

This directory contains a comprehensive demonstration of all Bootstrap 5 components available in the `pkg/bs5` package. Each component example is organized in cards with icons for better visual presentation.

## Structure

The application is split into separate example files for better organization and maintainability:

- **`main.go`** - Main entry point with tabbed interface orchestrating all component examples
- **`nav_example.go`** - Navigation bars with dark, light, and primary color schemes with icons
- **`alert_example.go`** - Alert components with different color variants and dismissible options
- **`badge_example.go`** - Badge components with icons, various colors, sizes, and pill styles
- **`button_example.go`** - Button components with icons, solid/outline variants, sizes, and icon-only buttons
- **`modal_example.go`** - Modal dialogs including login form with input groups and icons
- **`accordion_example.go`** - Accordion component with multiple collapsible items
- **`breadcrumb_example.go`** - Breadcrumb navigation with custom dividers in card layouts
- **`card_example.go`** - Card components with headers, footers, images, and various layouts
- **`pagination_example.go`** - Pagination with arrow icons, disabled states, and different alignments
- **`table_example.go`** - Tables with striped, hoverable, bordered, and colored row variations
- **`tabs_example.go`** - Tabs and pills navigation with content panes
- **`progress_example.go`** - Progress bars with colors, stripes, animation, and stacked variants
- **`toast_example.go`** - Toast notifications with icons, colored backgrounds, and auto-hide options
- **`offcanvas_example.go`** - Offcanvas panels with different placements and backdrop options
- **`grid_example.go`** - Responsive grid layouts with Bootstrap's column system
- **`form_example.go`** - Form controls including inputs, checkboxes, radios, and validation
- **`icon_example.go`** - Bootstrap Icons showcase with input groups

## Building

To build the WASM application:

```bash
GOOS=js GOARCH=wasm go build -o build/bootstrap-app.wasm ./cmd/wasm/bootstrap-app
```

## Running

Serve the application using the wasmserver:

```bash
./build/wasmserver serve --watch ./cmd/wasm/bootstrap-app --bs-5
```

Then open your browser to the displayed URL (typically `http://localhost:9090`).

## Component Coverage

This demo application showcases:

### Layout & Navigation

- ✅ **Navigation** - Navbars with dark, light, and primary color schemes, brand, items, dropdowns, and spacers
- ✅ **Breadcrumbs** - Navigation trails with custom dividers and click handlers
- ✅ **Grid System** - Responsive layouts with Row and Col components supporting all breakpoints
- ✅ **Cards** - Card layouts with headers, footers, images, and various styling options
- ✅ **Tabs** - Tab navigation with content panes, pills style, and justified layouts

### Components

- ✅ **Alerts** - Contextual alerts with all color variants and dismissible functionality
- ✅ **Badges** - Badges with icons, all color variants, pill shapes, and notification styles
- ✅ **Buttons** - Solid, outline, and size variants with icons, icon-only buttons, and action buttons
- ✅ **Modal** - Modal dialogs with login form, input groups with icons, and various sizes
- ✅ **Accordion** - Collapsible accordion items with event handling
- ✅ **Pagination** - Pagination with arrow icons, disabled states, sizes, and alignment options
- ✅ **Progress Bars** - Progress indicators with colors, heights, stripes, animation, and stacked variants
- ✅ **Toast** - Toast notifications with contextual icons, colored backgrounds, and auto-hide
- ✅ **Offcanvas** - Slide-out panels with different placements, backdrop options, and dark theme
- ✅ **Tables** - Data tables with striped rows, hover effects, borders, and colored rows

### Forms & Input

- ✅ **Form Controls** - Text inputs, textareas, selects, checkboxes, radios, and switches
- ✅ **Input Groups** - Input groups with prepended/appended icons and text
- ✅ **Form Validation** - Client-side validation with valid/invalid states

### Visual Enhancements

- ✅ **Bootstrap Icons** - Comprehensive icon integration throughout components
- ✅ **Card-Based Layout** - All examples organized in cards with headers for consistent presentation
- ✅ **Responsive Design** - Grid system supporting sm, md, lg, xl, xxl breakpoints

## Features

- **Icon Integration** - Bootstrap Icons used throughout for visual communication (envelope for email, lock for password, check/x for success/error, arrows for navigation)
- **Input Groups** - Enhanced forms with prepended/appended icons and text
- **Responsive Grid** - Full Bootstrap grid system with Row and Col components
- **Card Organization** - All component examples wrapped in cards with headers and descriptions
- **Interactive Examples** - Event handlers demonstrating click, dismiss, and toggle interactions
- **Color Schemes** - Full support for Bootstrap's contextual colors (primary, secondary, success, danger, warning, info, light, dark)

Each example demonstrates the full API of the component including styling options, event handling, and best practices.
