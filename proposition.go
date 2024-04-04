package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type PropositionType int

const (
	AND PropositionType = iota
	OR
	NAND
	NOR
	XOR
	XNOR
)

func (t PropositionType) toString() string {
	switch t {
	case AND:
		return "AND"
	case OR:
		return "OR"
	case NAND:
		return "NAND"
	case NOR:
		return "NOR"
	case XOR:
		return "XOR"
	case XNOR:
		return "XNOR"
	}
	return "unknown"
}

type Proposition struct {
	Type            PropositionType
	SubPremises     *PremiseSlice
	SubPropositions *PropositionSlice
}

type PropositionSlice []Proposition

func (p Proposition) toString() string {
	returnString := ""

	for _, premise := range *p.SubPremises {
		if len(returnString) > 0 {
			returnString += fmt.Sprintf(" %s %s", p.Type.toString(), premise.toString())
			continue
		}

		returnString += premise.toString()
	}

	for _, subProposition := range *p.SubPropositions {
		if len(returnString) > 0 {
			returnString += fmt.Sprintf(" %s %s", p.Type.toString(), subProposition.toString())
			continue
		}

		returnString += subProposition.toString()
	}

	return returnString
}

func (ps PropositionSlice) checkIfPropositionExists(proposition Proposition) bool {
	for _, p := range ps {
		if p.toString() == proposition.toString() {
			return true
		}
	}
	return false
}

func isStringInPremiseSlice(str string, list PremiseSlice) bool {
	for _, v := range list {
		if v.toString() == str {
			return true
		}
	}
	return false
}

func findPremiseInSlice(str string, premiseStack *PremiseSlice) *Premise {
	for _, v := range *premiseStack {
		if v.toString() == str {
			return &v
		}
	}
	return nil

}

func isStringInPropositionSlice(str string, list PropositionSlice) bool {
	for _, v := range list {
		if v.toString() == str {
			return true
		}
	}
	return false
}

func findPropositionInSlice(str string, propositionStack *PropositionSlice) *Proposition {
	for _, v := range *propositionStack {
		if v.toString() == str {
			return &v
		}
	}
	return nil

}

func selectPremiseForTerm(premiseStack *PremiseSlice, alreadyChosenPremises *PremiseSlice) {
	templates := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}

	premiseOptions := []string{}

	for _, premise := range *premiseStack {
		premiseString := premise.toString()
		if isStringInPremiseSlice(premiseString, *alreadyChosenPremises) {
			continue
		}

		premiseOptions = append(premiseOptions, premiseString)
	}

	promptPremise := promptui.Select{
		Label:     "Select a premise",
		Items:     premiseOptions,
		Templates: templates,
	}

	_, selectedPremise, err := promptPremise.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", selectedPremise)

	premiseToAdd := findPremiseInSlice(selectedPremise, premiseStack)

	if premiseToAdd == nil {
		fmt.Printf("Unable to find the premise %v\n", selectedPremise)
		return
	}

	*alreadyChosenPremises = append(*alreadyChosenPremises, *premiseToAdd)
}

func selectPropositionForTerm(propositionStack *PropositionSlice, alreadyChosenPropositions *PropositionSlice) {
	templates := &promptui.SelectTemplates{
		Active:   templatePropositionActive,
		Inactive: templatePropositionInactive,
	}

	propositionOptions := []string{}

	for _, proposition := range *propositionStack {
		if isStringInPropositionSlice(proposition.toString(), *alreadyChosenPropositions) {
			continue
		}

		propositionOptions = append(propositionOptions, proposition.toString())
	}

	promptProposition := promptui.Select{
		Label:     "Select a proposition",
		Items:     propositionOptions,
		Templates: templates,
	}

	_, selectedProposition, err := promptProposition.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", selectedProposition)

	propositionToAdd := findPropositionInSlice(selectedProposition, propositionStack)

	*alreadyChosenPropositions = append(*alreadyChosenPropositions, *propositionToAdd)
}

func hasPremiseToAddAsTerm(premiseStack *PremiseSlice, alreadyChosenPremises *PremiseSlice) bool {
	premiseOptions := []string{}

	for _, premise := range *premiseStack {
		if isStringInPremiseSlice(premise.toString(), *alreadyChosenPremises) {
			continue
		}

		premiseOptions = append(premiseOptions, premise.toString())
	}

	return len(premiseOptions) > 0
}

func hasPropositionToAddAsTerm(propositionStack *PropositionSlice, alreadyChosenPropositions *PropositionSlice) bool {
	propositionOptions := []string{}

	for _, proposition := range *propositionStack {
		if isStringInPropositionSlice(proposition.toString(), *alreadyChosenPropositions) {
			continue
		}

		propositionOptions = append(propositionOptions, proposition.toString())
	}

	return len(propositionOptions) > 0
}

func userSelectTerm(premiseStack *PremiseSlice, propositionStack *PropositionSlice, selectedPremises *PremiseSlice, selectedPropositions *PropositionSlice, selectedPropositionType *PropositionType) {
	// TODO return errors too and handle errors in parent function
	var promptOptions []string

	templateTypeOfTerm := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}

	if hasPremiseToAddAsTerm(premiseStack, selectedPremises) {
		promptOptions = append(promptOptions, "Premise")
	}
	if hasPropositionToAddAsTerm(propositionStack, selectedPropositions) {
		promptOptions = append(promptOptions, "Proposition")
	}

	promptFirstTermType := promptui.Select{
		Label:     "What kind of term would you like to add first?",
		Items:     promptOptions,
		Templates: templateTypeOfTerm,
	}

	_, input, err := promptFirstTermType.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch input {
	case "Premise":
		selectPremiseForTerm(premiseStack, selectedPremises)
	case "Proposition":
		selectPropositionForTerm(propositionStack, selectedPropositions)
	}

	fmt.Printf("Selected Proposition Type: %v\n", selectedPropositionType.toString())
	fmt.Printf("Selected premises:\n")
	for _, premise := range *selectedPremises {
		fmt.Printf("%v\n", premise.toString())
	}
	fmt.Printf("Selected propositions:\n")
	for _, proposition := range *selectedPropositions {
		fmt.Printf("%v\n", proposition.toString())
	}
}

func createProposition(premiseStack *PremiseSlice, propositionStack *PropositionSlice) {
	fmt.Print(string(ColorProposition))
	selectedPremises := PremiseSlice{}
	selectedPropositions := PropositionSlice{}
	hasAtLeastTwoTerms := len(*premiseStack)+len(*propositionStack) >= 2

	if (!hasPremiseToAddAsTerm(premiseStack, &selectedPremises) && !hasPropositionToAddAsTerm(propositionStack, &selectedPropositions)) || !hasAtLeastTwoTerms {
		fmt.Printf("There are not enough possible terms (premises or other propositions) in order to create a proposition. Create more premises or propositions and try again.\n")
		return
	}

	PropositionTypeOptions := []string{AND.toString(), OR.toString(), NAND.toString(), NOR.toString(), XOR.toString(), XNOR.toString()}

	templates := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}

	promptPropositionType := promptui.Select{
		Label:     "Select a proposition type",
		Items:     PropositionTypeOptions,
		Templates: templates,
	}

	_, inputType, err := promptPropositionType.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var selectedPropositionType PropositionType

	switch inputType {
	case AND.toString():
		selectedPropositionType = AND
	case OR.toString():
		selectedPropositionType = OR
	case NAND.toString():
		selectedPropositionType = NAND
	case NOR.toString():
		selectedPropositionType = NOR
	case XOR.toString():
		selectedPropositionType = XOR
	case XNOR.toString():
		selectedPropositionType = XNOR
	}

	fmt.Printf("You chose %q\n", selectedPropositionType.toString())

	for i := 0; i < 2; i++ {
		userSelectTerm(premiseStack, propositionStack, &selectedPremises, &selectedPropositions, &selectedPropositionType)
	}

	userIsDone := false

	for !userIsDone && (hasPremiseToAddAsTerm(premiseStack, &selectedPremises) || hasPropositionToAddAsTerm(propositionStack, &selectedPropositions)) {
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

		if result == "Yes" {
			userIsDone = true
			break
		}

		userSelectTerm(premiseStack, propositionStack, &selectedPremises, &selectedPropositions, &selectedPropositionType)
	}

	propositionToAdd := Proposition{Type: selectedPropositionType, SubPremises: &selectedPremises, SubPropositions: &selectedPropositions}
	fmt.Printf("Proposition to add: %v\n", propositionToAdd.toString())

	if propositionStack.checkIfPropositionExists(propositionToAdd) {
		fmt.Println(string(ColorError), "Proposition already exists")
		return
	}

	*propositionStack = append(*propositionStack, propositionToAdd)
}

func listPropositions(propositionStack *PropositionSlice) {
	fmt.Print(string(ColorProposition), STARLINE+"Propositions:\n")
	for _, proposition := range *propositionStack {
		fmt.Println(string(ColorProposition), proposition.toString())
	}
}
