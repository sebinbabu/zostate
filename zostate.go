// Package zostate implements a finite state machine.
package zostate

type TransitionType string
type StateType string

type machineState struct {
	Transitions map[TransitionType]StateType
}

type Machine struct {
	name    string
	current StateType
	states  map[StateType]machineState
	initial StateType
}

type Transition struct {
	Name TransitionType
	Dst  StateType
}

type Transitions []Transition

type State struct {
	Name        StateType
	Transitions Transitions
}

type States []State

func NewMachine(name string, initial StateType, states States) (*Machine, error) {
	machineStates := make(map[StateType]machineState)

	for _, s := range states {
		state := machineState{
			Transitions: make(map[TransitionType]StateType),
		}

		for _, t := range s.Transitions {
			state.Transitions[t.Name] = t.Dst
		}

		machineStates[s.Name] = state
	}

	if _, ok := machineStates[initial]; !ok {
		return nil, ErrMachineCreationFailed
	}

	machine := &Machine{
		name:    name,
		states:  machineStates,
		initial: initial,
	}

	return machine, nil
}

func (m *Machine) Current() StateType {
	if m.current == "" {
		return m.initial
	}
	return m.current
}

func (m *Machine) Transition(event TransitionType) (StateType, error) {
	current := m.Current()
	next, ok := m.states[current].Transitions[event]

	if !ok {
		return current, ErrTransitionFailed
	}

	m.current = next
	return next, nil
}

func (m *Machine) States() States {
	states := make(States, 0)

	for name, state := range m.states {
		s := State{
			Name:        name,
			Transitions: make(Transitions, 0),
		}

		for tname, dst := range state.Transitions {
			s.Transitions = append(s.Transitions, Transition{
				Name: tname,
				Dst:  dst,
			})
		}

		states = append(states, s)
	}

	return states
}
