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

type SubjectSlice []Subject

func (s Subject) toString() string {
	if s.Body == "" {
		return "unknown"

	}
	return s.Body
}

func (ss SubjectSlice) checkIfSubjectExists(subject Subject) bool {
	for _, s := range ss {
		if s.Body == subject.Body {
			return true
		}
	}
	return false

}

func createSubject(subjectStack *SubjectSlice) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorSubject), STARLINE+"Enter the Body for the new subject:\n")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	if subjectStack.checkIfSubjectExists(Subject{Body: input}) {
		fmt.Println(string(ColorError), "Subject already exists")
		return
	}

	*subjectStack = append(*subjectStack, Subject{Body: input})
}

func listSubjects(subjectStack *SubjectSlice) {
	fmt.Print(string(ColorSubject), STARLINE+"Subjects:\n")
	for _, subject := range *subjectStack {
		fmt.Println(string(ColorSubject), subject.Body)
	}
}
