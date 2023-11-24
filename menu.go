package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var menuItems = []string{
	"(L)ist Todos",
	"(C)reate a Todo",
	"(D)elete a Todo",
	"(Q)uit",
}

func displayMenuItems() {
	for _, item := range menuItems {
		fmt.Printf("%s, ", item)
	}
	fmt.Printf("Your choice: ")
}

func readChoice() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	return strings.ToUpper(string(input[0])), nil
}
