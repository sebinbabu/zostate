package zostate

import (
	"bytes"
	"fmt"
	"sort"
)

// DrawMachine accepts a Machine and "draws" it and returns a string
// in the Graphviz DOT format. This string can be parsed by Graphviz
// to be converted into a diagram.
func DrawMachine(machine *Machine) string {
	var buf bytes.Buffer
	states := machine.States()

	sortStates(states)

	writeHeader(&buf, machine.name)
	writeTransitions(&buf, machine.initial, states)
	writeStates(&buf, states)
	writeStartPoint(&buf, machine.initial)
	writeFooter(&buf)

	return buf.String()
}

// Writes the header to buffer in DOT format.
func writeHeader(buf *bytes.Buffer, name string) {
	buf.WriteString(fmt.Sprintf(`digraph %s {`, name))
	buf.WriteString("\n")
}

// Writes all state transitions to the buffer in DOT format.
func writeTransitions(buf *bytes.Buffer, current StateType, states States) {
	for _, k := range states {
		if k.Name == current {
			for _, t := range k.Transitions {
				buf.WriteString(fmt.Sprintf(`    "%s" -> "%s" [ label = "%s" ];`, k.Name, t.Dst, t.Event))
				buf.WriteString("\n")
			}
		}
	}

	for _, k := range states {
		if k.Name != current {
			for _, t := range k.Transitions {
				buf.WriteString(fmt.Sprintf(`    "%s" -> "%s" [ label = "%s" ];`, k.Name, t.Dst, t.Event))
				buf.WriteString("\n")
			}
		}
	}

	buf.WriteString("\n")
}

// Writes all states to buffer in DOT format.
func writeStates(buf *bytes.Buffer, states States) {
	for _, k := range states {
		buf.WriteString(fmt.Sprintf(`    "%s";`, k.Name))
		buf.WriteString("\n")
	}
}

// Writes the pointer to the starting node to the buffer in DOT format.
func writeStartPoint(buf *bytes.Buffer, initial StateType) {
	buf.WriteString(fmt.Sprintln("    node [width=0.3 shape=point style=filled];"))
	buf.WriteString(fmt.Sprintf(`    "" -> "%s";`, initial))
	buf.WriteString("\n")
}

// Writes the footer to buffer in DOT format.
func writeFooter(buf *bytes.Buffer) {
	buf.WriteString(fmt.Sprintln("}"))
}

// Sorts the state descriptions and transitions in ascending order
// so that the diagram doesn't change after each run.
func sortStates(states States) {
	sort.Slice(states, func(i, j int) bool {
		return states[i].Name < states[j].Name
	})

	for _, state := range states {
		sort.Slice(state.Transitions, func(i, j int) bool {
			return state.Transitions[i].Event < state.Transitions[j].Event
		})
	}
}
