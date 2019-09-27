package main

import "fmt"

type State struct {
	id int
}

func NewState(id int) *State {
	return &State{id: id}
}

type Transition struct {
	sourceState *State
	symbol           int32
	destinationState *State
}

func NewTransition(sourceState *State, symbol int32, destinationState *State) *Transition {
	return &Transition{sourceState: sourceState, symbol: symbol, destinationState: destinationState}
}

type DeterministicFiniteAutomaton struct {
	initialState *State
	alphabet     []int32
	currentState *State
	transitions  []*Transition
	states       []*State
	endStates    []*State
}

func NewDeterministicFiniteAutomaton(initialState *State, alphabet []int32, transitions []*Transition, states []*State, endStates []*State) *DeterministicFiniteAutomaton {
	automaton := &DeterministicFiniteAutomaton{
		initialState: initialState,
		alphabet: alphabet,
		transitions: transitions,
		states: states,
		endStates: endStates,
	}

	// Set current state to initial state
	automaton.ResetAutomaton()

	return automaton
}



func (dfa *DeterministicFiniteAutomaton) ResetAutomaton()  {

	dfa.currentState = dfa.initialState
}

func (dfa *DeterministicFiniteAutomaton) ProcessWord(word string) bool {

	dfa.ResetAutomaton()

	for _, currentRune := range word {

		for _, currentTransition := range dfa.transitions {

			if currentTransition.sourceState == dfa.currentState && currentTransition.symbol == currentRune {
				destinationState := currentTransition.destinationState
				dfa.currentState = destinationState
				break
			}
		}
	}

	for _, endStateCandidate := range dfa.endStates {
		if endStateCandidate == dfa.currentState {
			return true
		}
	}

	return false
}

func main()  {

	state1 := NewState(1)
	state2 := NewState(2)

	transition11 := NewTransition(state1, 'a', state1)
	transition12 := NewTransition(state1, 'b', state2)
	transition21 := NewTransition(state2, 'a', state1)
	transition22 := NewTransition(state2, 'b', state2)

	automaton := NewDeterministicFiniteAutomaton(state1, []int32{'a', 'b'}, []*Transition{transition11, transition12, transition21, transition22}, []*State{state1, state2}, []*State{state2})

	acceptsWord := automaton.ProcessWord("aba")
	fmt.Println("Accepts word aba: ", acceptsWord)

	acceptsWord = automaton.ProcessWord("ab")
	fmt.Println("Accepts word ab: ", acceptsWord)

	acceptsWord = automaton.ProcessWord("abaab")
	fmt.Println("Accepts word abaab: ", acceptsWord)

	acceptsWord = automaton.ProcessWord("bbbbbbbbba")
	fmt.Println("Accepts word bbbbbbbbba: ", acceptsWord)
}