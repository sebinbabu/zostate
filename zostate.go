// Package zostate implements a simple finite state machine library, which helps you declaratively model FSMs.
package zostate

// EventType represents all events that exist in the state machine.
type EventType string

// StateType represents all states that exist in the state machine.
type StateType string

// state represents the configuration of a state in the machine, one of which is the transitions.
type state struct {
	// transitions map an event with a destination state.
	transitions map[EventType]StateType
}

// Machine represents a finite state machine.
type Machine struct {
	name    string              // Name of the machine
	current StateType           // Current state of the machine
	initial StateType           // Initial state of the machine
	states  map[StateType]state // states maps the name of a state with the state configuration
}

// Transition maps an event with a destination state.
type Transition struct {
	Event EventType // Name of the event that triggers the transition
	Dst   StateType // Destination state of the transition
}

// Transitions is a slice of related transitions.
type Transitions []Transition

// StateDesc is the description of a state.
type StateDesc struct {
	Name        StateType   // Name holds the name of state
	Transitions Transitions // Transitions holds all valid transitions originating from the state
}

// States is an slice of related state descriptions (StateDesc).
type States []StateDesc

// NewMachine returns an instance of the state machine (Machine) after it accepts
// the name of the details of the machine, such as initial state & state descriptions.
// It returns an error if the parameters are invalid.
func NewMachine(name string, initial StateType, states States) (*Machine, error) {
	mStates := make(map[StateType]state)

	for _, s := range states {
		state := state{
			transitions: make(map[EventType]StateType),
		}

		for _, t := range s.Transitions {
			state.transitions[t.Event] = t.Dst
		}

		mStates[s.Name] = state
	}

	if _, ok := mStates[initial]; !ok {
		return nil, ErrMachineCreationFailed
	}

	machine := &Machine{
		name:    name,
		states:  mStates,
		initial: initial,
	}

	return machine, nil
}

// Returns the current state that the machine is in.
func (m *Machine) Current() StateType {
	if m.current == "" {
		return m.initial
	}
	return m.current
}

// getNextState returns the next state for a given event after considering
// the current state and it's transitions. If a valid transition doesn't exist
// it returns an error.
func (m *Machine) getNextState(event EventType) (StateType, error) {
	current := m.Current()
	next, ok := m.states[current].transitions[event]

	if !ok {
		return "", ErrEventDeclined
	}

	return next, nil
}

// Transition returns the final state after accepting a valid event, and transitioning
// to the next state after considering the current state and it's transitions.
// If a valid transition doesn't exist it returns an error.
func (m *Machine) Transition(event EventType) (StateType, error) {
	next, err := m.getNextState(event)
	if err != nil {
		return m.current, ErrTransitionFailed
	}

	m.current = next
	return next, nil
}

// States returns a slice of all state descriptions (StateDesc) of the state machine.
func (m *Machine) States() States {
	states := make(States, 0)

	for name, state := range m.states {
		s := StateDesc{
			Name:        name,
			Transitions: make(Transitions, 0),
		}

		for tname, dst := range state.transitions {
			s.Transitions = append(s.Transitions, Transition{
				Event: tname,
				Dst:   dst,
			})
		}

		states = append(states, s)
	}

	return states
}
