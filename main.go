package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	COMMAND_CREATE_STATEMENT = "create-statement"
	COMMAND_CREATE_PREDICATE = "create-predicate"
	COMMAND_CREATE_SUBJECT   = "create-subject"
	COMMAND_LIST_SUBJECTS    = "list-subjects"

	STARLINE = "**********\n"

	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func initializeStacks() (subjectStack *[]Subject) {
	subjectStack = &[]Subject{}
	return subjectStack

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	subjectStack := initializeStacks()

	for {
		fmt.Print(
			string(ColorGreen),
			STARLINE+
				"Enter one of the following commands:\n"+
				COMMAND_CREATE_STATEMENT+"\n"+
				COMMAND_CREATE_PREDICATE+"\n"+
				COMMAND_CREATE_SUBJECT+"\n"+
				COMMAND_LIST_SUBJECTS+"\n"+
				"exit\n",
		)
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		switch input {
		case COMMAND_CREATE_STATEMENT:
		case COMMAND_CREATE_PREDICATE:
		case COMMAND_CREATE_SUBJECT:
			createSubject(subjectStack)
		case COMMAND_LIST_SUBJECTS:
			listSubjects(subjectStack)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println(string(ColorRed), "Invalid command!")
		}

	}
}
