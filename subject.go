package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Subject struct {
	Value string
}

func createSubject(subjectStack *[]Subject) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(string(ColorPurple), STARLINE+"Enter the value for the new subject:\n")
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	*subjectStack = append(*subjectStack, Subject{Value: input})
}

func listSubjects(subjectStack *[]Subject) {
	fmt.Println(string(ColorPurple), STARLINE+"Subjects:")
	for _, subject := range *subjectStack {
		fmt.Println(string(ColorPurple), subject.Value)
	}
}
