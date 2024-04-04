package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

const (
	COMMAND_CREATE_PROPOSITION = "create proposition"
	COMMAND_CREATE_PREMISE     = "create premise"
	COMMAND_CREATE_PREDICATE   = "create predicate"
	COMMAND_CREATE_SUBJECT     = "create subject"
	COMMAND_LIST_PROPOSITIONS  = "list propositions"
	COMMAND_LIST_PREMISES      = "list premises"
	COMMAND_LIST_PREDICATES    = "list predicates"
	COMMAND_LIST_SUBJECTS      = "list subjects"
	COMMAND_EXIT               = "exit"

	STARLINE = "**********\n"

	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"

	// ColorReset  = "\033[0m"
	ColorError       = ColorRed
	ColorSubject     = ColorGreen
	ColorPredicate   = ColorYellow
	ColorProposition = ColorBlue
	ColorPremise     = ColorPurple
	// ColorPurple = "\033[35m"
	// ColorCyan   = "\033[36m"
	// ColorWhite  = "\033[37m"
)

func initializeStacks() (subjectStack SubjectSlice, predicateStack PredicateSlice, premiseStack PremiseSlice, propositionStack PropositionSlice) {
	subjectStack = SubjectSlice{
		{Body: "the ball"},
		{Body: "the sky"},
	}
	predicateStack = PredicateSlice{
		{Body: "is red"},
		{Body: "is blue"},
	}
	premiseStack = PremiseSlice{
		{Subject: subjectStack[0], Predicate: predicateStack[0], SubjectQuantifier: ALL, PredicateQualifier: IS},
		{Subject: subjectStack[0], Predicate: predicateStack[1], SubjectQuantifier: ALL, PredicateQualifier: IS},
	}
	propositionStack = PropositionSlice{
		{Type: OR, SubPremises: &PremiseSlice{premiseStack[0], premiseStack[1]}, SubPropositions: &PropositionSlice{}},
	}
	return subjectStack, predicateStack, premiseStack, propositionStack

}

func main() {
	subjectStack, predicateStack, premiseStack, propositionStack := initializeStacks()

	for {
		templates := &promptui.SelectTemplates{
			Active:   templateGenericActive,
			Inactive: templateGenericInactive,
		}

		prompt := promptui.Select{
			Label:     "Select one of the following commands:",
			Items:     []string{COMMAND_CREATE_PROPOSITION, COMMAND_CREATE_PREMISE, COMMAND_CREATE_PREDICATE, COMMAND_CREATE_SUBJECT, COMMAND_LIST_PROPOSITIONS, COMMAND_LIST_PREMISES, COMMAND_LIST_PREDICATES, COMMAND_LIST_SUBJECTS, COMMAND_EXIT},
			Templates: templates,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case COMMAND_CREATE_PROPOSITION:
			createProposition(&premiseStack, &propositionStack)
		case COMMAND_CREATE_PREMISE:
			createPremise(&subjectStack, &predicateStack, &premiseStack)
		case COMMAND_CREATE_PREDICATE:
			createPredicate(&predicateStack)
		case COMMAND_CREATE_SUBJECT:
			createSubject(&subjectStack)
		case COMMAND_LIST_PROPOSITIONS:
			listPropositions(&propositionStack)
		case COMMAND_LIST_PREMISES:
			listPremises(&premiseStack)
		case COMMAND_LIST_PREDICATES:
			listPredicates(&predicateStack)
		case COMMAND_LIST_SUBJECTS:
			listSubjects(&subjectStack)
		case COMMAND_EXIT:
			os.Exit(0)
		default:
			fmt.Println(string(ColorReset), "Invalid command!")
		}

	}
}
