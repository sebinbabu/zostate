# zostate
Simple & declarative finite state machines in Go

Visit https://pkg.go.dev/github.com/sebinbabu/zostate for API documentation.

## Quick start

```go
package main

import (
	"fmt"

	"github.com/sebinbabu/zostate"
)

const (
	On  zostate.StateType = "On"
	Off zostate.StateType = "Off"
)

const (
	Toggle zostate.EventType = "Toggle"
)

func main() {
	machine, err := zostate.NewMachine(
		"LightBulb",
		Off,
		zostate.States{
			{
				Name: Off,
				Transitions: zostate.Transitions{
					{Event: Toggle, Dst: On},
				},
			},
			{
				Name: On,
				Transitions: zostate.Transitions{
					{Event: Toggle, Dst: Off},
				},
			},
		},
	)

	fmt.Println("current state: ", machine.Current()) // current state: Off
	machine.Transition(Toggle)
	fmt.Println("current state: ", machine.Current()) // current state: On
}
```

## Todo

1. Add tests
2. Add support for callbacks for events, states & transitions
3. Implement [statecharts](https://statecharts.dev/)

## References

* [xstate](https://xstate.js.org/): State machines and statecharts for the modern web
* [looplab/fsm](https://github.com/looplab/fsm): Finite State Machine for Go