package behavioral_patterns

import "fmt"

type memento struct {
	state string
}

func (m *memento) getState() string {
	return m.state
}

type originator struct {
	state string
}

func (or *originator) setState(s string) {
	or.state = s
}

func (or *originator) getState() string {
	return or.state
}

func (or *originator) createMemento() memento {
	return memento{
		state: or.state,
	}
}

func (or *originator) restore(m memento) {
	or.setState(m.getState())
}


type careTaker struct {
	states []memento
}

func (crt *careTaker) addState(m memento) {
	crt.states = append(crt.states, m)
}

func (crt *careTaker) getState(index int) memento {
	return crt.states[index]
}

func MementoPattern() {
	org := originator{
		state: "A",
	}

	crTaker := careTaker {
		states: make([]memento, 0),
	}

	fmt.Printf("Storing originator state: %v\n", org.getState())
	crTaker.addState(org.createMemento())

	org.setState("B")

	fmt.Printf("Storing originator state: %v\n", org.getState())
	crTaker.addState(org.createMemento())

	org.setState("C")

	fmt.Printf("Current state of originator is %v\n", org.getState())
	fmt.Printf("Undo originator state\n")
	
	org.restore(crTaker.getState(1))

	fmt.Printf("Current state of originator is %v\n", org.getState())
	fmt.Println("One more undo")

	org.restore(crTaker.getState(0))
	fmt.Printf("Current state of originator is %v\n", org.getState())

}
