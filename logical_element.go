package main

type LogicalElement interface {
	toString() string
}

// TODO fix, these don't seem to work
func isPremise(e LogicalElement) bool {
	_, ok := e.(*Premise)
	return ok
}

func isProposition(e LogicalElement) bool {
	_, ok := e.(*Proposition)
	return ok
}

func isConditionalStatement(e LogicalElement) bool {
	_, ok := e.(*ConditionalStatement)
	return ok
}
