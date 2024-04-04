package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type SubjectQuantifier int
type PredicateQualifier int

const (
	ALL SubjectQuantifier = iota
	SOME
	NONE
)

func (s SubjectQuantifier) toString() string {
	switch s {
	case ALL:
		return "ALL OF"
	case SOME:
		return "SOME OF"
	case NONE:
		return "NONE OF"
	}
	return "unknown"
}

const (
	IS PredicateQualifier = iota
	IS_NOT
)

func (p PredicateQualifier) toString() string {
	switch p {
	case IS:
		return "IS"
	case IS_NOT:
		return "IS NOT"
	}
	return "unknown"
}

type Premise struct {
	Subject            Subject
	SubjectQuantifier  SubjectQuantifier
	Predicate          Predicate
	PredicateQualifier PredicateQualifier
}

type PremiseSlice []Premise

func (p Premise) toString() string {
	return fmt.Sprintf("%s %s : %s %s", p.SubjectQuantifier.toString(), p.Subject.toString(), p.PredicateQualifier.toString(), p.Predicate.toString())
}

func (ps PremiseSlice) checkIfPremiseExists(premise Premise) bool {
	for _, p := range ps {
		if p.Subject.Body == premise.Subject.Body && p.Predicate.Body == premise.Predicate.Body {
			return true
		}
	}
	return false
}

func createPremise(subjectStack *SubjectSlice, predicateStack *PredicateSlice, premiseStack *PremiseSlice) {
	fmt.Print(string(ColorPremise))
	subjectOptions := make([]string, len(*subjectStack))

	for i, subject := range *subjectStack {
		subjectOptions[i] = subject.Body
	}

	subjectTemplates := &promptui.SelectTemplates{
		Active:   templateSubjectActive,
		Inactive: templateSubjectInactive,
	}

	promptSubject := promptui.Select{
		Label:     "Select a subject",
		Items:     *subjectStack,
		Templates: subjectTemplates,
	}

	i, _, err := promptSubject.Run()

	selectedSubject := (*subjectStack)[i]

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", selectedSubject.toString())
	fmt.Printf("Premise: ? %s : ? ?\n", selectedSubject.toString())

	SubjectQuantifierOptions := []string{ALL.toString(), SOME.toString(), NONE.toString()}

	templates := &promptui.SelectTemplates{
		Active:   templateGenericActive,
		Inactive: templateGenericInactive,
	}

	promptSubjectQuantifier := promptui.Select{
		Label:     "Select a subject quantifier",
		Items:     SubjectQuantifierOptions,
		Templates: templates,
	}

	_, inputQuantifier, err := promptSubjectQuantifier.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var selectedSubjectQuantifier SubjectQuantifier

	switch inputQuantifier {
	case ALL.toString():
		selectedSubjectQuantifier = ALL
	case SOME.toString():
		selectedSubjectQuantifier = SOME
	case NONE.toString():
		selectedSubjectQuantifier = NONE
	}

	fmt.Printf("You chose %q\n", selectedSubjectQuantifier.toString())
	fmt.Printf("Premise: %s %s : ? ?\n", selectedSubjectQuantifier.toString(), selectedSubject.toString())

	predicateOptions := make([]string, len(*predicateStack))

	for i, predicate := range *predicateStack {
		predicateOptions[i] = predicate.Body
	}

	predicateTemplates := &promptui.SelectTemplates{
		Active:   templatePredicateActive,
		Inactive: templatePredicateInactive,
	}

	promptPedicate := promptui.Select{
		Label:     "Select a predicate",
		Items:     *predicateStack,
		Templates: predicateTemplates,
	}

	x, _, err := promptPedicate.Run()

	selectedPredicate := (*predicateStack)[x]

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %q\n", selectedPredicate.toString())
	fmt.Printf("Premise: %s %s : ? %s\n", selectedSubjectQuantifier.toString(), selectedSubject.toString(), selectedPredicate.toString())

	PredicateQualifierOptions := []string{IS.toString(), IS_NOT.toString()}

	promptPredicateQualifier := promptui.Select{
		Label:     "Select a predicate quantifier",
		Items:     PredicateQualifierOptions,
		Templates: templates,
	}

	_, inputQualifier, err := promptPredicateQualifier.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var selectedPredicateQualifier PredicateQualifier

	switch inputQualifier {
	case IS.toString():
		selectedPredicateQualifier = IS
	case IS_NOT.toString():
		selectedPredicateQualifier = IS_NOT
	}

	fmt.Printf("You chose %q\n", selectedPredicateQualifier.toString())
	fmt.Printf("Premise: %s %s : %s %s\n", selectedSubjectQuantifier.toString(), selectedSubject.toString(), selectedPredicateQualifier.toString(), selectedPredicate.toString())

	if premiseStack.checkIfPremiseExists(Premise{Subject: selectedSubject, SubjectQuantifier: selectedSubjectQuantifier, Predicate: selectedPredicate, PredicateQualifier: selectedPredicateQualifier}) {
		fmt.Println(string(ColorError), "Premise already exists")
		return
	}

	*premiseStack = append(*premiseStack, Premise{Subject: selectedSubject, SubjectQuantifier: selectedSubjectQuantifier, Predicate: selectedPredicate, PredicateQualifier: selectedPredicateQualifier})
}

func listPremises(premiseStack *PremiseSlice) {
	fmt.Print(string(ColorPremise), STARLINE+"Premises:\n")
	for _, premise := range *premiseStack {
		fmt.Println(string(ColorPremise), premise.toString())
	}
}
