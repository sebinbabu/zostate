package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sebinbabu/zostate"
)

const (
	Green zostate.StateType = "Green"
	Amber zostate.StateType = "Amber"
	Red   zostate.StateType = "Red"
)

const (
	TimerExpire zostate.EventType = "TimerExpire"
	Reset       zostate.EventType = "Reset"
)

func main() {
	machine, err := zostate.NewMachine(
		"trafficlight",
		Green,
		zostate.States{
			{
				Name: Green,
				Transitions: zostate.Transitions{
					{Event: TimerExpire, Dst: Amber},
					{Event: Reset, Dst: Amber},
				},
			},
			{
				Name: Amber,
				Transitions: zostate.Transitions{
					{Event: TimerExpire, Dst: Red},
					{Event: Reset, Dst: Amber},
				},
			},
			{
				Name: Red,
				Transitions: zostate.Transitions{
					{Event: TimerExpire, Dst: Green},
					{Event: Reset, Dst: Amber},
				},
			},
		},
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintln(os.Stderr, "current state: ", machine.Current())

	dot := zostate.DrawMachine(machine)
	io.Copy(os.Stdout, strings.NewReader(dot))
}
