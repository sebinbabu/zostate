package main

import (
	"fmt"

	"github.com/sebinbabu/zostate"
)

const (
	ON  zostate.StateType = "On"
	OFF zostate.StateType = "Off"
)

const (
	TOGGLE zostate.EventType = "Toggle"
)

func main() {
	machine := zostate.Machine{
		Initial: OFF,
		States: zostate.States{
			OFF: zostate.State{
				On: zostate.Transitions{
					TOGGLE: ON,
				},
			},
			ON: zostate.State{
				On: zostate.Transitions{
					TOGGLE: OFF,
				},
			},
		},
	}

	fmt.Println("current", machine.Current())

	current, err := machine.Transition(TOGGLE)
	fmt.Println("current", current, err)

	current, err = machine.Transition("vivi")
	fmt.Println("current", current, err)
}
