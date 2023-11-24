package main

import (
	"fmt"
)

func main() {

	for {
		displayMenuItems()

		menuChoice, err := readChoice()

		if err != nil {
			fmt.Println("Error reading menu selection:", err)
			return
		}

		switch menuChoice {

		case "L":
			listTodos()
			continue
		case "C":
			createTodo()
			continue
		case "D":
			deleteTodo()
			continue
		case "Q":
			fmt.Println("Quitting...")
		default:
			continue
		}

		break
	}
}
