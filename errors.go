package zostate

import "errors"

// ErrMachineCreationFailed is the error returned when creation fails.
var ErrMachineCreationFailed = errors.New("machine creation failed")

// ErrTransitionFailed is the error returned when a transition fails.
// This can be due to an invalid event.
var ErrTransitionFailed = errors.New("transition failed")

// ErrEventDeclined is the error returned when an event cannot be
// processed by the state machine, due to it's current state.
var ErrEventDeclined = errors.New("event declined")
