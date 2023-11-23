package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
)

type Todo struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

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

	return todos, nil
}

func listTodos() {
	readTodos, err := readTodosFromFile()
	if err != nil {
		fmt.Println("Error reading todos:", err)
		return
	}

	fmt.Println("Todos:")
	for _, todo := range readTodos {
		fmt.Printf("%s\n", todo.Content)
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

	// write content to json file
	existingTodos, err := readTodosFromFile()
	if err != nil {
		fmt.Println("Error reading JSON:", err)
		return
	}
	existingTodos = append(existingTodos, newTodo)

	// Encode the updated slice back to JSON
	updatedData, err := json.MarshalIndent(existingTodos, "", "    ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Write the updated JSON data back to the file
	if err := os.WriteFile(filename, updatedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Todo added successfully.")

}

func isBlank(s string) bool {
	// Trim spaces from the string and check if it's empty
	return len(strings.TrimSpace(s)) == 0
}

func deleteTodo() {
	listTodos()

}
