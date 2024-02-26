package main

import (
	"fmt"
	"strings"

	"github.com/SornchaiTheDev/todo_list/commands"
	"github.com/SornchaiTheDev/todo_list/models"
)

func greeting(todos []models.Todo) {

	todoLength := len(todos)

	fmt.Printf("----- Todo (%d) -----\n", todoLength)
	fmt.Println("Commands: ")
	fmt.Println("1. CREATE to add new todo")
	fmt.Println("3. READ to get all todos")
	fmt.Println("1. MARK to mark todo as complete")
	fmt.Println("2. DELETE to delete todo from the list")
	fmt.Println("----------------")
	fmt.Printf("Enter Command : ")

}

func main() {

	var todos []models.Todo

	var command string

	for {
		greeting(todos)

		fmt.Scanf("%s", &command)

		lowerCmd := strings.ToLower(command)

		if lowerCmd == "q" {
			fmt.Println("Bye Bye")
			break
		} else if lowerCmd == "create" {
			commands.AddTodo(&todos)
		} else if lowerCmd == "read" {
			commands.ReadTodos(todos)
		} else if lowerCmd == "mark" {
			commands.MarkTodo(&todos)
		} else if lowerCmd == "delete" {
			commands.DeleteTodo(&todos)
		} else {
			fmt.Println("Unknown Command!")
		}
	}

}
