package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type Argument struct {
	Title      string
	Elements   []LogicalElement
	Conclusion LogicalElement
}

type ArgumentSlice []Argument

func (a Argument) toString() string {
	var elementsString string
	for i, e := range a.Elements {

		elementsString += fmt.Sprintf("%d. %s\n", i+1, e.toString())
	}

	return fmt.Sprintf("%sConclusion: %s\n", elementsString, a.Conclusion.toString())
}

func (as ArgumentSlice) checkIfArgumentExists(argument Argument) bool {
	for _, a := range as {
		if a.toString() == argument.toString() {
			return true
		}
	}
	return false
}

func generatePremiseOptionsForLogicalElement(premiseStack *PremiseSlice, alreadyChosenLogicalElements *[]LogicalElement) (premiseOptions []string) {
	for _, premise := range *premiseStack {
		addPremise := true

		if alreadyChosenLogicalElements != nil {
			for _, chosenElement := range *alreadyChosenLogicalElements {
				if premise.toString() == chosenElement.toString() {
					addPremise = false
					break
				}
			}
		}

		if addPremise {
			premiseOptions = append(premiseOptions, premise.toString())
		}
	}

	return premiseOptions
}

func generatePropositionOptionsForLogicalElement(propositionStack *PropositionSlice, alreadyChosenLogicalElements *[]LogicalElement) (propositionOptions []string) {
	for _, proposition := range *propositionStack {
		addProposition := true

		if alreadyChosenLogicalElements != nil {
			for _, chosenElement := range *alreadyChosenLogicalElements {
				if proposition.toString() == chosenElement.toString() {
					addProposition = false
					break
				}
			}
		}

		if addProposition {
			propositionOptions = append(propositionOptions, proposition.toString())
		}
	}

	return propositionOptions
}

func generateConditionalStatementOptionsForLogicalElement(conditionalStack *ConditionalStatementSlice, alreadyChosenLogicalElements *[]LogicalElement) (conditionalOptions []string) {
	for _, conditionalStatement := range *conditionalStack {
		addConditionalStatement := true

		if alreadyChosenLogicalElements != nil {
			for _, chosenElement := range *alreadyChosenLogicalElements {
				if conditionalStatement.toString() == chosenElement.toString() {
					addConditionalStatement = false
					break
				}
			}
		}

		if addConditionalStatement {
			conditionalOptions = append(conditionalOptions, conditionalStatement.toString())
		}
	}

	return conditionalOptions
}

func hasPremiseToAddAsLogicalElement(premiseStack *PremiseSlice, alreadyChosenLogicalElements *[]LogicalElement) bool {
	premiseOptions := generatePremiseOptionsForLogicalElement(premiseStack, alreadyChosenLogicalElements)

	return len(premiseOptions) > 0
}

func hasPropositionToAddAsLogicalElement(propositionStack *PropositionSlice, alreadyChosenLogicalElements *[]LogicalElement) bool {
	propositionOptions := generatePropositionOptionsForLogicalElement(propositionStack, alreadyChosenLogicalElements)

	return len(propositionOptions) > 0
}

func hasConditionalStatementToAddAsLogicalElement(conditionalStack *ConditionalStatementSlice, alreadyChosenLogicalElements *[]LogicalElement) bool {
	conditionalOptions := generateConditionalStatementOptionsForLogicalElement(conditionalStack, alreadyChosenLogicalElements)

	return len(conditionalOptions) > 0
}

func selectPremiseForLogicalElement(premiseStack *PremiseSlice, alreadyChosenLogicalElements *[]LogicalElement) (newlySelectedLogicalElement LogicalElement, err error) {
	premiseOptions := generatePremiseOptionsForLogicalElement(premiseStack, alreadyChosenLogicalElements)

	if len(premiseOptions) == 0 {
		return nil, fmt.Errorf("no premises to choose from")
	}

	prompt := promptui.Select{
		Label: "Select a premise",
		Items: premiseOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	for _, premise := range *premiseStack {
		if premise.toString() == result {
			return premise, nil
		}
	}

	return nil, fmt.Errorf("premise not found")

}

func selectPropositionForLogicalElement(propositionStack *PropositionSlice, alreadyChosenLogicalElements *[]LogicalElement) (newlySelectedLogicalElement LogicalElement, err error) {
	propositionOptions := generatePropositionOptionsForLogicalElement(propositionStack, alreadyChosenLogicalElements)

	if len(propositionOptions) == 0 {
		return nil, fmt.Errorf("no propositions to choose from")
	}

	prompt := promptui.Select{
		Label: "Select a proposition",
		Items: propositionOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	for _, proposition := range *propositionStack {
		if proposition.toString() == result {
			return proposition, nil
		}
	}

	return nil, fmt.Errorf("proposition not found")
}

func selectConditionalStatementForLogicalElement(conditionalStack *ConditionalStatementSlice, alreadyChosenLogicalElements *[]LogicalElement) (newlySelectedLogicalElement LogicalElement, err error) {
	conditionalOptions := generateConditionalStatementOptionsForLogicalElement(conditionalStack, alreadyChosenLogicalElements)

	if len(conditionalOptions) == 0 {
		return nil, fmt.Errorf("no conditional statements to choose from")
	}

	prompt := promptui.Select{
		Label: "Select a conditional statement",
		Items: conditionalOptions,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	for _, conditionalStatement := range *conditionalStack {
		if conditionalStatement.toString() == result {
			return conditionalStatement, nil
		}
	}

	return nil, fmt.Errorf("conditional statement not found")
}

func userSelectLogicalElement(statementNumber int, premiseStack *PremiseSlice, propositionStack *PropositionSlice, conditionalStack *ConditionalStatementSlice, alreadyChosenLogicalElements *[]LogicalElement) (newlySelectedLogicalElement LogicalElement, err error) {
	var promptOptions []string
	templateTypeOfConditionalElement := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}
	if hasPremiseToAddAsLogicalElement(premiseStack, alreadyChosenLogicalElements) {
		promptOptions = append(promptOptions, "Premise")
	}
	if hasPropositionToAddAsLogicalElement(propositionStack, alreadyChosenLogicalElements) {
		promptOptions = append(promptOptions, "Proposition")
	}
	if hasConditionalStatementToAddAsLogicalElement(conditionalStack, alreadyChosenLogicalElements) {
		promptOptions = append(promptOptions, "Conditional Statement")
	}

	promptString := fmt.Sprintf("What kind of logic statement would you like to add for statement %d?", statementNumber)
	prompt := promptui.Select{
		Label:     promptString,
		Items:     promptOptions,
		Templates: templateTypeOfConditionalElement,
	}
	_, input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	switch input {
	case "Premise":
		newlySelectedLogicalElement, err = selectPremiseForLogicalElement(premiseStack, alreadyChosenLogicalElements)
	case "Proposition":
		newlySelectedLogicalElement, err = selectPropositionForLogicalElement(propositionStack, alreadyChosenLogicalElements)
	case "Conditional Statement":
		newlySelectedLogicalElement, err = selectConditionalStatementForLogicalElement(conditionalStack, alreadyChosenLogicalElements)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return newlySelectedLogicalElement, nil
}

func userSelectConclusion(premiseStack *PremiseSlice, propositionStack *PropositionSlice, conditionalStack *ConditionalStatementSlice) (conclusion LogicalElement, err error) {
	var promptOptions []string
	templateTypeOfConditionalElement := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}
	if hasPremiseToAddAsLogicalElement(premiseStack, nil) {
		promptOptions = append(promptOptions, "Premise")
	}
	if hasPropositionToAddAsLogicalElement(propositionStack, nil) {
		promptOptions = append(promptOptions, "Proposition")
	}
	if hasConditionalStatementToAddAsLogicalElement(conditionalStack, nil) {
		promptOptions = append(promptOptions, "Conditional Statement")
	}

	promptString := "What kind of logic statement would you like for the conclusion?"
	prompt := promptui.Select{
		Label:     promptString,
		Items:     promptOptions,
		Templates: templateTypeOfConditionalElement,
	}
	_, input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	switch input {
	case "Premise":
		conclusion, err = selectPremiseForLogicalElement(premiseStack, nil)
	case "Proposition":
		conclusion, err = selectPropositionForLogicalElement(propositionStack, nil)
	case "Conditional Statement":
		conclusion, err = selectConditionalStatementForLogicalElement(conditionalStack, nil)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return conclusion, nil
}

func createArgument(argumentStack *ArgumentSlice, premiseStack *PremiseSlice, propositionStack *PropositionSlice, conditionalStack *ConditionalStatementSlice) {
	hasAtLeastTwoLogicalElements := len(*premiseStack)+len(*propositionStack)+len(*conditionalStack) >= 2

	if !hasAtLeastTwoLogicalElements {
		fmt.Println(string(ColorReset), "You need at least two conditional elements (e.g premises, propositions, and conditional statements) to create an argument!")
		return
	}

	userIsDone := false

	selectedLogicalElements := []LogicalElement{}

	for !userIsDone && (hasPremiseToAddAsLogicalElement(premiseStack, &selectedLogicalElements) || hasPropositionToAddAsLogicalElement(propositionStack, &selectedLogicalElements) || hasConditionalStatementToAddAsLogicalElement(conditionalStack, &selectedLogicalElements)) {
		newLogicalElement, err := userSelectLogicalElement(len(selectedLogicalElements)+1, premiseStack, propositionStack, conditionalStack, &selectedLogicalElements)

		if err != nil {
			fmt.Println(string(ColorError), "Error: ", err)
			return
		}

		selectedLogicalElements = append(selectedLogicalElements, newLogicalElement)

		templates := &promptui.SelectTemplates{
			Active:   templateGenericActive,
			Inactive: templateGenericInactive,
		}

		prompt := promptui.Select{
			Label:     "Are you done adding terms to the proposition?",
			Items:     []string{"Yes", "No"},
			Templates: templates,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		hasAddedAllPossibleElements := len(*premiseStack)+len(*propositionStack)+len(*conditionalStack) == len(selectedLogicalElements)

		if result == "Yes" || hasAddedAllPossibleElements {
			userIsDone = true
			break
		}

	}

	conclusion, err := userSelectConclusion(premiseStack, propositionStack, conditionalStack)

	if err != nil {
		fmt.Println(string(ColorError), "Error: ", err)
		return
	}

	newArgument := Argument{Elements: selectedLogicalElements, Conclusion: conclusion}

	fmt.Printf("Argument to add:\n%v\n", newArgument.toString())

	if argumentStack.checkIfArgumentExists(newArgument) {
		fmt.Println(string(ColorError), "Argument already exists")
		return
	}

	*argumentStack = append(*argumentStack, newArgument)
}

func listArguments(argumentStack *ArgumentSlice) {
	for i, argument := range *argumentStack {
		fmt.Printf("%d. %s\n", i+1, argument.toString())
	}
}
