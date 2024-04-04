package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Predicate struct {
	Body string
}

type PredicateSlice []Predicate

func (p Predicate) toString() string {
	if p.Body == "" {
		return "unknown"

	}
	return p.Body
}

func (ps PredicateSlice) checkIfPredicateExists(predicate Predicate) bool {
	for _, p := range ps {
		if p.Body == predicate.Body {
			return true
		}
	}
	return false

}

func createPredicate(predicateStack *PredicateSlice) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorPredicate), STARLINE+"Enter the Body for the new Predicate:\n")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	if predicateStack.checkIfPredicateExists(Predicate{Body: input}) {
		fmt.Println(string(ColorError), "Predicate already exists")
		return
	}

	*predicateStack = append(*predicateStack, Predicate{Body: input})
}

func listPredicates(predicateStack *PredicateSlice) {
	fmt.Print(string(ColorPredicate), STARLINE+"Predicates:\n")
	for _, Predicate := range *predicateStack {
		fmt.Println(string(ColorPredicate), Predicate.Body)
	}
}
