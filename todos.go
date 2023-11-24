package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Todo struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var numberOfTodos int

var filename string = "todos.json"

func readTodosFromFile() ([]Todo, error) {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	err = json.Unmarshal(fileContent, &todos)
	if err != nil {
		return nil, err
	}

	numberOfTodos = len(todos)

	return todos, nil
}

func listTodos() {
	readTodos, err := readTodosFromFile()
	if err != nil {
		fmt.Println("Error reading todos:", err)
		return
	}

	fmt.Println("Todos:")
	for index, todo := range readTodos {
		fmt.Printf("%d. %s\n", index+1, todo.Content)
	}

}

func promptForContent() (string, error) {
	fmt.Printf("Enter todo content: ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	// Trim trailing newline character
	input = strings.TrimRight(input, "\n")

	if err != nil {
		return "", err
	}

	return input, nil
}

func createTodo() {
	content, err := promptForContent()

	if err != nil {
		fmt.Println("Error creating todo:", err)
		return
	}

	if isBlank(content) {
		fmt.Println("No content.")
		return
	}

	newTodo := Todo{
		ID:      uuid.New().String(),
		Content: content,
	}

	// get existing todos
	existingTodos, err := readTodosFromFile()
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}

	// append new todo
	existingTodos = append(existingTodos, newTodo)

	writeTodosToFile(existingTodos)

	fmt.Println("Todo added successfully.")
}

func isBlank(s string) bool {
	// Trim spaces from the string and check if it's empty
	return len(strings.TrimSpace(s)) == 0
}

func deleteTodo() {
	listTodos()

	var index int

	for {
		fmt.Printf("Delete number: ")

		todoChoice, err := readChoice()
		if err != nil {
			fmt.Println("Error reading selection:", err)
			return
		}

		index, err = strconv.Atoi(todoChoice)
		if err != nil {
			continue
		}

		if index < 1 || index > numberOfTodos {
			continue
		}
		break
	}

	existingTodos, err := readTodosFromFile()
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}

	// Remove the element at the specified index
	existingTodos = append(existingTodos[:index-1], existingTodos[index:]...)

	fmt.Println("Removing todo " + strconv.Itoa(index))

	writeTodosToFile(existingTodos)
}

func writeTodosToFile(todos []Todo) {
	// Encode the slice back to JSON
	updatedData, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Write the JSON data to the file
	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
