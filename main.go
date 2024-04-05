package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

const (
	COMMAND_CREATE_ARGUMENT              = "create argument"
	COMMAND_CREATE_CONDITIONAL_STATEMENT = "create conditional statement"
	COMMAND_CREATE_PROPOSITION           = "create proposition"
	COMMAND_CREATE_PREMISE               = "create premise"
	COMMAND_CREATE_PREDICATE             = "create predicate"
	COMMAND_CREATE_SUBJECT               = "create subject"
	COMMAND_LIST_ARGUMENTS               = "list arguments"
	COMMAND_LIST_CONDITIONAL_STATEMENTS  = "list conditional statements"
	COMMAND_LIST_PROPOSITIONS            = "list propositions"
	COMMAND_LIST_PREMISES                = "list premises"
	COMMAND_LIST_PREDICATES              = "list predicates"
	COMMAND_LIST_SUBJECTS                = "list subjects"
	COMMAND_EXIT                         = "exit"

	STARLINE = "**********\n"

	ColorReset         = "\033[0m"
	ColorBlack         = "\033[30m"
	ColorRed           = "\033[31m"
	ColorGreen         = "\033[32m"
	ColorYellow        = "\033[33m"
	ColorBlue          = "\033[34m"
	ColorPurple        = "\033[35m"
	ColorCyan          = "\033[36m"
	ColorWhite         = "\033[37m"
	ColorBrightBlack   = "\033[90m"
	ColorBrightRed     = "\033[91m"
	ColorBrightGreen   = "\033[92m"
	ColorBrightYellow  = "\033[93m"
	ColorBrightBlue    = "\033[94m"
	ColorBrightMagenta = "\033[95m"
	ColorBrightCyan    = "\033[96m"
	ColorBrightWhite   = "\033[97m"

	// ColorReset  = "\033[0m"
	ColorError       = ColorRed
	ColorSubject     = ColorGreen
	ColorPredicate   = ColorBrightGreen
	ColorProposition = ColorBlue
	ColorPremise     = ColorPurple
	ColorConditional = ColorCyan
	ColorArgument    = ColorYellow
)

func initializeStacks() (subjectStack SubjectSlice, predicateStack PredicateSlice, premiseStack PremiseSlice, propositionStack PropositionSlice, conditionalStatementStack ConditionalStatementSlice, argumentStack ArgumentSlice) {
	subjectStack = SubjectSlice{
		{Body: "the ball"},
		{Body: "the sky"},
		{Body: "the time to play"},
	}
	predicateStack = PredicateSlice{
		{Body: "red"},
		{Body: "blue"},
		{Body: "now"},
	}
	premiseStack = PremiseSlice{
		{Subject: subjectStack[0], Predicate: predicateStack[0], SubjectQuantifier: ALL, PredicateQualifier: IS},
		{Subject: subjectStack[0], Predicate: predicateStack[1], SubjectQuantifier: ALL, PredicateQualifier: IS},
		{Subject: subjectStack[2], Predicate: predicateStack[2], SubjectQuantifier: ALL, PredicateQualifier: IS},
		{Subject: subjectStack[2], Predicate: predicateStack[2], SubjectQuantifier: ALL, PredicateQualifier: IS_NOT},
	}
	propositionStack = PropositionSlice{
		{Type: OR, SubPremises: &PremiseSlice{premiseStack[0], premiseStack[1]}, SubPropositions: &PropositionSlice{}},
		{Type: OR, SubPremises: &PremiseSlice{premiseStack[2], premiseStack[3]}, SubPropositions: &PropositionSlice{}},
	}
	conditionalStatementStack = ConditionalStatementSlice{
		{IfElement: premiseStack[0], ThenElement: premiseStack[2]},
		{IfElement: premiseStack[1], ThenElement: premiseStack[3]},
	}
	argumentStack = ArgumentSlice{
		{Title: "Is it time to play?", Elements: []LogicalElement{conditionalStatementStack[0], premiseStack[0]}, Conclusion: premiseStack[2]},
	}
	return subjectStack, predicateStack, premiseStack, propositionStack, conditionalStatementStack, argumentStack

}

func main() {
	subjectStack, predicateStack, premiseStack, propositionStack, conditionalStatementStack, argumentStack := initializeStacks()

	for {
		templates := &promptui.SelectTemplates{
			Active:   templateGenericActive,
			Inactive: templateGenericInactive,
		}

		prompt := promptui.Select{
			Label:     "Select one of the following commands:",
			Items:     []string{COMMAND_CREATE_ARGUMENT, COMMAND_CREATE_PROPOSITION, COMMAND_CREATE_PREMISE, COMMAND_CREATE_PREDICATE, COMMAND_CREATE_SUBJECT, COMMAND_LIST_ARGUMENTS, COMMAND_LIST_PROPOSITIONS, COMMAND_LIST_PREMISES, COMMAND_LIST_PREDICATES, COMMAND_LIST_SUBJECTS, COMMAND_EXIT},
			Templates: templates,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		switch result {
		case COMMAND_CREATE_ARGUMENT:
			createArgument(&argumentStack, &premiseStack, &propositionStack, &conditionalStatementStack)
		case COMMAND_CREATE_CONDITIONAL_STATEMENT:
			createConditionalStatement(&conditionalStatementStack, &premiseStack, &propositionStack)
		case COMMAND_CREATE_PROPOSITION:
			createProposition(&premiseStack, &propositionStack)
		case COMMAND_CREATE_PREMISE:
			createPremise(&subjectStack, &predicateStack, &premiseStack)
		case COMMAND_CREATE_PREDICATE:
			createPredicate(&predicateStack)
		case COMMAND_CREATE_SUBJECT:
			createSubject(&subjectStack)
		case COMMAND_LIST_ARGUMENTS:
			listArguments(&argumentStack)
		case COMMAND_LIST_CONDITIONAL_STATEMENTS:
			listConditionalStatements(&conditionalStatementStack)
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
