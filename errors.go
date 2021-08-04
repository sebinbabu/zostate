package zostate

import "errors"

var ErrMachineCreationFailed = errors.New("machine creation failed")
var ErrTransitionFailed = errors.New("transition failed")
