package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// Custom String() method
// When a ServerState value (e.g., StateIdle) is printed using fmt.Println(), Go implicitly calls its String() method.
// The method looks up the corresponding string in the stateName map using the enum value as the key.
// Example: StateIdle.String() returns "idle".
func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns) // ns is of type ServerState
	// fmt.Println() checks whether ServerState implements String().
	// Since it does, it calls ns.String(),
	// which retrieves the corresponding string from stateName

	ns2 := transition(ns)
	fmt.Println(ns2)

	/*
		First transition:
			transition(StateIdle) → StateConnected
			fmt.Println(ns) prints "connected" (retrieved via stateName[StateConnected]).
		Second transition:
			transition(StateConnected) → StateIdle
			fmt.Println(ns2) prints "idle" (retrieved via stateName[StateIdle]).
	*/
}

// transition emulates a state transition for a server;
// it takes the existing state and returns a new state.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))

	}
}
