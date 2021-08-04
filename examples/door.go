package main

import (
	"fmt"
	"io/ioutil"

	"github.com/sebinbabu/zostate"
)

const (
	OPEN   zostate.StateType = "Open"
	CLOSED zostate.StateType = "Closed"
)

const (
	TOGGLE zostate.TransitionType = "Toggle"
)

func main() {
	machine, err := zostate.NewMachine(
		"door",
		CLOSED,
		zostate.States{
			{
				Name: OPEN,
				Transitions: zostate.Transitions{
					{Name: TOGGLE, Dst: CLOSED},
				},
			},
			{
				Name: CLOSED,
				Transitions: zostate.Transitions{
					{Name: TOGGLE, Dst: OPEN},
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("current", machine.Current())

	current, err := machine.Transition(TOGGLE)
	fmt.Println("current", current, err)

	current, err = machine.Transition(TOGGLE)
	fmt.Println("current", current, err)

	current, err = machine.Transition("vivi")
	fmt.Println("current", current, err)

	dot := zostate.Visualize(machine)
	err = ioutil.WriteFile("./door.dot", []byte(dot), 0644)
	if err != nil {
		fmt.Println(err)
	}
}
