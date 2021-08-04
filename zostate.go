package zostate

type EventType string
type StateType string

type Transitions map[EventType]StateType

type State struct {
	On Transitions
}

type States map[StateType]State

type Machine struct {
	current StateType

	Initial StateType
	States  States
}

func (m *Machine) Current() StateType {
	if m.current == "" {
		return m.Initial
	}
	return m.current
}

func (m *Machine) Transition(event EventType) (StateType, error) {
	current := m.Current()
	next, ok := m.States[current].On[event]

	if !ok {
		return current, ErrTransitionFailed
	}

	m.current = next
	return next, nil
}
