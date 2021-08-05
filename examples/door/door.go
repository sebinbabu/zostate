package main

import (
	"fmt"
	"io"
	"os"
	"strings"

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

	fmt.Fprintln(os.Stderr, "current", machine.Current())

	current, err := machine.Transition(TOGGLE)
	fmt.Fprintln(os.Stderr, "current", current, err)

	current, err = machine.Transition(TOGGLE)
	fmt.Fprintln(os.Stderr, "current", current, err)

	current, err = machine.Transition("vivi")
	fmt.Fprintln(os.Stderr, "current", current, err)

	dot := zostate.DrawMachine(machine)
	io.Copy(os.Stdout, strings.NewReader(dot))
}
