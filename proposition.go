package main

import (
	"fmt"
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
	SubPremises     *[]Premise
	SubPropositions *[]Proposition
}

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

func listPropositions(propositionStack *[]Proposition) {
	fmt.Print(string(ColorProposition), STARLINE+"Propositions:\n")
	for _, proposition := range *propositionStack {
		fmt.Printf("%s\n", proposition.toString())
	}
}
