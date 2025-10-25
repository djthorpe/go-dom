package bootstrap

///////////////////////////////////////////////////////////////////////////////
// OFFCANVAS OPTIONS

// WithScroll enables body scrolling when offcanvas is open.
// By default, body scrolling is disabled when offcanvas is shown.
func WithScroll() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-scroll"] = "true"
		return nil
	}
}

// WithBackdrop configures the backdrop behavior for offcanvas.
// Use "false" for no backdrop, "static" for a backdrop that doesn't close on click,
// or omit this option for the default backdrop behavior.
func WithBackdrop(backdrop string) Opt {
	return func(o *opts) error {
		if backdrop == "false" || backdrop == "static" {
			o.attributes["data-bs-backdrop"] = backdrop
		}
		return nil
	}
}

// WithoutKeyboard disables closing the offcanvas with the escape key.
// By default, pressing escape closes the offcanvas.
func WithoutKeyboard() Opt {
	return func(o *opts) error {
		o.attributes["data-bs-keyboard"] = "false"
		return nil
	}
}
