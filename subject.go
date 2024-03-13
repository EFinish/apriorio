package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Subject struct {
	Body string
}

func (s Subject) toString() string {
	if s.Body == "" {
		return "unknown"

	}
	return s.Body
}

func createSubject(subjectStack *[]Subject) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorSubject), STARLINE+"Enter the Body for the new subject:\n")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	*subjectStack = append(*subjectStack, Subject{Body: input})
}

func listSubjects(subjectStack *[]Subject) {
	fmt.Print(string(ColorSubject), STARLINE+"Subjects:\n")
	for _, subject := range *subjectStack {
		fmt.Println(string(ColorSubject), subject.Body)
	}
}
