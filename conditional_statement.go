package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type ConditionalElementType int

const (
	IF ConditionalElementType = iota
	THEN
)

func (c ConditionalElementType) toString() string {
	switch c {
	case IF:
		return "IF"
	case THEN:
		return "THEN"
	}
	return "unknown"
}

type ConditionalElement interface {
	toString() string
}

type ConditionalStatement struct {
	IfElement   ConditionalElement
	ThenElement ConditionalElement
}

type ConditionalStatementSlice []ConditionalStatement

func (c ConditionalStatement) toString() string {
	return fmt.Sprintf("IF %s THEN %s", c.IfElement.toString(), c.ThenElement.toString())
}

func (cs ConditionalStatementSlice) checkIfConditionalStatementExists(conditionalStatement ConditionalStatement) bool {
	for _, c := range cs {
		if c.IfElement.toString() == conditionalStatement.IfElement.toString() && c.ThenElement.toString() == conditionalStatement.ThenElement.toString() {
			return true
		}
	}
	return false
}

func generatePremiseOptionsForConditionalElement(premiseStack *PremiseSlice, alreadyChosenConditionalElement *ConditionalElement) (premiseOptions []string) {
	for _, premise := range *premiseStack {
		if alreadyChosenConditionalElement != nil && premise.toString() == (*alreadyChosenConditionalElement).toString() {
			continue
		}

		premiseOptions = append(premiseOptions, premise.toString())
	}

	return premiseOptions
}

func generatePropositionOptionsForConditionalElement(propositionStack *PropositionSlice, alreadyChosenConditionalElement *ConditionalElement) (propositionOptions []string) {
	for _, proposition := range *propositionStack {
		if alreadyChosenConditionalElement != nil && proposition.toString() == (*alreadyChosenConditionalElement).toString() {
			continue
		}

		propositionOptions = append(propositionOptions, proposition.toString())
	}

	return propositionOptions
}

func selectPremiseForConditionalElement(premiseStack *PremiseSlice, alreadyChosenConditionalElement *ConditionalElement) (newlySelectedConditionalElement ConditionalElement, err error) {
	premiseOptions := generatePremiseOptionsForConditionalElement(premiseStack, alreadyChosenConditionalElement)

	if len(premiseOptions) == 0 {
		err = fmt.Errorf("no premises to choose from")
		return nil, err
	}

	prompt := promptui.Select{
		Label: "Select a premise",
		Items: premiseOptions,
	}

	_, input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	for _, premise := range *premiseStack {
		if premise.toString() == input {
			return premise, nil
		}
	}

	err = fmt.Errorf("Premise not found in premise stack")

	return nil, err
}

func selectPropositionForConditionalElement(propositionStack *PropositionSlice, alreadyChosenConditionalElement *ConditionalElement) (newlySelectedConditionalElement ConditionalElement, err error) {
	propositionOptions := generatePropositionOptionsForConditionalElement(propositionStack, alreadyChosenConditionalElement)

	if len(propositionOptions) == 0 {
		err = fmt.Errorf("no propositions to choose from")
		return nil, err
	}

	prompt := promptui.Select{
		Label: "Select a proposition",
		Items: propositionOptions,
	}

	_, input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return nil, err
	}

	for _, proposition := range *propositionStack {
		if proposition.toString() == input {
			return proposition, nil
		}
	}

	err = fmt.Errorf("Proposition not found in proposition stack")

	return nil, err
}

func hasPremiseToAddAsConditionalElement(premiseStack *PremiseSlice, alreadyChosenConditionalElement *ConditionalElement) bool {
	premiseOptions := generatePremiseOptionsForConditionalElement(premiseStack, alreadyChosenConditionalElement)

	return len(premiseOptions) > 0
}

func hasPropositionToAddAsConditionalElement(propositionStack *PropositionSlice, alreadyChosenConditionalElement *ConditionalElement) bool {
	propositionOptions := generatePropositionOptionsForConditionalElement(propositionStack, alreadyChosenConditionalElement)

	return len(propositionOptions) > 0
}

func userSelectConditionalElement(conditionalElementType ConditionalElementType, premiseStack *PremiseSlice, propositionStack *PropositionSlice, alreadySelectedConditionalElement *ConditionalElement) (newlySelectedConditionalElement ConditionalElement, err error) {
	var promptOptions []string
	templateTypeOfConditionalElement := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}
	if hasPremiseToAddAsConditionalElement(premiseStack, alreadySelectedConditionalElement) {
		promptOptions = append(promptOptions, "Premise")
	}
	if hasPropositionToAddAsConditionalElement(propositionStack, alreadySelectedConditionalElement) {
		promptOptions = append(promptOptions, "Proposition")
	}
	promptString := fmt.Sprintf("What kind of conditional statement would you like to add for the %v statement?", conditionalElementType.toString())
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
		newlySelectedConditionalElement, err = selectPremiseForConditionalElement(premiseStack, alreadySelectedConditionalElement)
	case "Proposition":
		newlySelectedConditionalElement, err = selectPropositionForConditionalElement(propositionStack, alreadySelectedConditionalElement)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil, err
	}

	return newlySelectedConditionalElement, nil
}

func createConditionalStatement(conditionalStatementStack *ConditionalStatementSlice, premiseStack *PremiseSlice, propositionStack *PropositionSlice) {
	hasAtLeastTwoConditionalElements := len(*premiseStack)+len(*propositionStack) >= 2

	if !hasAtLeastTwoConditionalElements {
		fmt.Println(string(ColorReset), "You need at least two conditional elements (e.g premises and propositions) to create a conditional statement!")
		return
	}

	ifStatement, err := userSelectConditionalElement(IF, premiseStack, propositionStack, nil)

	if err != nil {
		fmt.Println(string(ColorError), "Error: ", err)
		return
	}

	fmt.Printf("Selected conditional statements:\n")
	fmt.Printf("%v\n", ifStatement.toString())

	thenStatement, err := userSelectConditionalElement(THEN, premiseStack, propositionStack, &ifStatement)

	if err != nil {
		fmt.Println(string(ColorError), "Error: ", err)
		return
	}

	fmt.Printf("Selected conditional statements:\n")
	fmt.Printf("%v\n", ifStatement.toString())
	fmt.Printf("%v\n", thenStatement.toString())

	newConditionalStatement := ConditionalStatement{
		IfElement:   ifStatement,
		ThenElement: thenStatement,
	}

	if conditionalStatementStack.checkIfConditionalStatementExists(newConditionalStatement) {
		fmt.Println(string(ColorError), "This conditional statement already exists!")
		return
	}

	*conditionalStatementStack = append(*conditionalStatementStack, newConditionalStatement)
}

func listConditionalStatements(conditionalStatementStack *ConditionalStatementSlice) {
	if len(*conditionalStatementStack) == 0 {
		fmt.Println(string(ColorConditional), "There are no conditional statements to list!")
		return
	}

	fmt.Println(string(ColorConditional), "Conditional Statements:")
	for _, conditionalStatement := range *conditionalStatementStack {
		fmt.Printf("%v\n", conditionalStatement.toString())
	}
}
