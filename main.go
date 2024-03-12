package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.Replace(input, "\n", "", -1)

		if input == "exit" {
			break
		}
	}
}
