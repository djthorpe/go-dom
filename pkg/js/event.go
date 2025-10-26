//go:build js && wasm

package js

import (
	"syscall/js"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

// EventTarget wraps a JavaScript EventTarget object
type EventTarget struct {
	value Value
}

// Event wraps a JavaScript Event object
type Event struct {
	value Value
}

// EventListener is a function that handles events
type EventListener func(event *Event)

// EventListenerOptions represents options for addEventListener
type EventListenerOptions struct {
	Capture bool  // Event will be dispatched to this listener before any EventTarget beneath it in the DOM tree
	Once    bool  // Listener will be removed after being invoked once
	Passive bool  // Listener will never call preventDefault()
	Signal  Value // AbortSignal to remove the listener
}

///////////////////////////////////////////////////////////////////////////////
// CONSTRUCTORS

// NewEventTarget creates a new EventTarget object
func NewEventTarget() *EventTarget {
	return &EventTarget{
		value: Global().Get("EventTarget").New(),
	}
}

// NewEventTargetFromValue creates an EventTarget wrapper from a Value
func NewEventTargetFromValue(value Value) *EventTarget {
	return &EventTarget{value: value}
}

// NewEvent creates a new Event object with the specified type and options
func NewEvent(eventType string, options ...map[string]interface{}) *Event {
	var event Value
	if len(options) > 0 && options[0] != nil {
		opts := NewObject()
		for key, val := range options[0] {
			opts.Set(key, ValueOf(val))
		}
		event = Global().Get("Event").New(eventType, opts)
	} else {
		event = Global().Get("Event").New(eventType)
	}
	return &Event{value: event}
}

// NewEventFromValue creates an Event wrapper from a Value
func NewEventFromValue(value Value) *Event {
	return &Event{value: value}
}

// NewCustomEvent creates a new CustomEvent with the specified type and detail
func NewCustomEvent(eventType string, detail interface{}, options ...map[string]interface{}) *Event {
	opts := NewObject()
	opts.Set("detail", Unwrap(detail))

	if len(options) > 0 && options[0] != nil {
		for key, val := range options[0] {
			opts.Set(key, ValueOf(val))
		}
	}

	event := Global().Get("CustomEvent").New(eventType, opts)
	return &Event{value: event}
}

///////////////////////////////////////////////////////////////////////////////
// EVENT TARGET METHODS

// AddEventListener registers an event listener on the EventTarget
func (et *EventTarget) AddEventListener(eventType string, listener EventListener, options ...*EventListenerOptions) js.Func {
	callback := js.FuncOf(func(this Value, args []Value) interface{} {
		if len(args) > 0 {
			listener(NewEventFromValue(args[0]))
		}
		return nil
	})

	if len(options) > 0 && options[0] != nil {
		opts := buildEventListenerOptions(options[0])
		et.value.Call("addEventListener", eventType, callback, opts)
	} else {
		et.value.Call("addEventListener", eventType, callback)
	}

	return callback
}

// RemoveEventListener removes an event listener from the EventTarget
func (et *EventTarget) RemoveEventListener(eventType string, callback js.Func, options ...*EventListenerOptions) {
	if len(options) > 0 && options[0] != nil {
		opts := buildEventListenerOptions(options[0])
		et.value.Call("removeEventListener", eventType, callback, opts)
	} else {
		et.value.Call("removeEventListener", eventType, callback)
	}
}

// DispatchEvent dispatches an event to this EventTarget
// Returns false if the event is cancelable and at least one of the event handlers
// which received the event called Event.preventDefault(). Otherwise it returns true.
func (et *EventTarget) DispatchEvent(event *Event) bool {
	return et.value.Call("dispatchEvent", event.value).Bool()
}

// JSValue returns the underlying Value
func (et *EventTarget) JSValue() Value {
	return et.value
}

///////////////////////////////////////////////////////////////////////////////
// EVENT METHODS

// Type returns the type of event
func (e *Event) Type() string {
	return e.value.Get("type").String()
}

// Target returns the EventTarget to which the event was originally dispatched
func (e *Event) Target() *EventTarget {
	target := e.value.Get("target")
	if target.IsNull() || target.IsUndefined() {
		return nil
	}
	return NewEventTargetFromValue(target)
}

// CurrentTarget returns the EventTarget whose event listeners are currently being processed
func (e *Event) CurrentTarget() *EventTarget {
	target := e.value.Get("currentTarget")
	if target.IsNull() || target.IsUndefined() {
		return nil
	}
	return NewEventTargetFromValue(target)
}

// EventPhase returns the phase of the event flow being processed
// 0 = NONE, 1 = CAPTURING_PHASE, 2 = AT_TARGET, 3 = BUBBLING_PHASE
func (e *Event) EventPhase() int {
	return e.value.Get("eventPhase").Int()
}

// Bubbles returns whether the event bubbles up through the DOM tree
func (e *Event) Bubbles() bool {
	return e.value.Get("bubbles").Bool()
}

// Cancelable returns whether the event is cancelable
func (e *Event) Cancelable() bool {
	return e.value.Get("cancelable").Bool()
}

// DefaultPrevented returns whether preventDefault() has been called on this event
func (e *Event) DefaultPrevented() bool {
	return e.value.Get("defaultPrevented").Bool()
}

// Composed returns whether the event will propagate across the shadow DOM boundary
func (e *Event) Composed() bool {
	return e.value.Get("composed").Bool()
}

// IsTrusted returns whether the event was initiated by the browser or by script
func (e *Event) IsTrusted() bool {
	return e.value.Get("isTrusted").Bool()
}

// TimeStamp returns the time at which the event was created (in milliseconds)
func (e *Event) TimeStamp() float64 {
	return e.value.Get("timeStamp").Float()
}

// PreventDefault cancels the event if it is cancelable
func (e *Event) PreventDefault() {
	e.value.Call("preventDefault")
}

// StopPropagation prevents further propagation of the current event
func (e *Event) StopPropagation() {
	e.value.Call("stopPropagation")
}

// StopImmediatePropagation prevents other listeners of the same event from being called
func (e *Event) StopImmediatePropagation() {
	e.value.Call("stopImmediatePropagation")
}

// Detail returns the detail property for CustomEvent
func (e *Event) Detail() Value {
	return e.value.Get("detail")
}

// JSValue returns the underlying Value
func (e *Event) JSValue() Value {
	return e.value
}

///////////////////////////////////////////////////////////////////////////////
// HELPER FUNCTIONS

// buildEventListenerOptions converts EventListenerOptions to a JavaScript object
func buildEventListenerOptions(options *EventListenerOptions) Value {
	opts := NewObject()

	if options.Capture {
		opts.Set("capture", true)
	}

	if options.Once {
		opts.Set("once", true)
	}

	if options.Passive {
		opts.Set("passive", true)
	}

	if !options.Signal.IsUndefined() && !options.Signal.IsNull() {
		opts.Set("signal", options.Signal)
	}

	return opts
}

// EventPhase constants
const (
	EventPhaseNone      = 0
	EventPhaseCapturing = 1
	EventPhaseAtTarget  = 2
	EventPhaseBubbling  = 3
)
