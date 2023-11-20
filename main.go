package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	// Create a sample todo list
	todos := []Todo{
		{1, "Learn Go"},
		{2, "Build CLI application"},
		{3, "Test the application"},
	}

	// Save todos to a file
	saveTodosToFile("todos.json", todos)

	// Read todos from the file
	readTodos, err := readTodosFromFile("todos.json")
	if err != nil {
		fmt.Println("Error reading todos:", err)
		return
	}

	// Print the todos
	fmt.Println("Todos:")
	for _, todo := range readTodos {
		fmt.Printf("ID: %d, Title: %s\n", todo.ID, todo.Title)
	}
}

func saveTodosToFile(filename string, todos []Todo) error {
	todosJSON, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, todosJSON, 0644)
	if err != nil {
		return err
	}

	fmt.Println("Todos saved to", filename)
	return nil
}

func readTodosFromFile(filename string) ([]Todo, error) {
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
