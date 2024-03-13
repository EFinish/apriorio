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

func (p Predicate) toString() string {
	if p.Body == "" {
		return "unknown"

	}
	return p.Body
}

func createPredicate(predicateStack *[]Predicate) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorPredicate), STARLINE+"Enter the Body for the new Predicate:\n")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	*predicateStack = append(*predicateStack, Predicate{Body: input})
}

func listPredicates(predicateStack *[]Predicate) {
	fmt.Print(string(ColorPredicate), STARLINE+"Predicates:\n")
	for _, Predicate := range *predicateStack {
		fmt.Println(string(ColorPurple), Predicate.Body)
	}
}
