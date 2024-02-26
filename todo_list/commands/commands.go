package commands

import (
	"fmt"

	"github.com/SornchaiTheDev/todo_list/models"
	"github.com/SornchaiTheDev/todo_list/utils"
)

func AddTodo(todos *[]models.Todo) {

	fmt.Printf("Enter task name : ")
	var taskName string

	fmt.Scanf("%s", &taskName)
	fmt.Printf("Task '%s' ? [y/n] : ", taskName)

	var confirm string
	fmt.Scanf("%s", &confirm)

	if confirm == "y" {
		*todos = append(*todos, models.Todo{Name: taskName, IsCompleted: false})
	}

}

func GetOrder() (order int) {

	fmt.Printf("Enter order of the todo : ")
	fmt.Scanf("%d", &order)

	return order
}

func getNotCompletedTodos(todos []models.Todo) []models.Todo {
	slice := utils.FilterSliceByCondition(todos,
		func(item models.Todo) bool { return !item.IsCompleted })
	return slice
}

func MarkTodo(todos *[]models.Todo) {

	order := GetOrder()

	notCompletedTodo := getNotCompletedTodos(*todos)

	completedTodo := utils.FilterSliceByCondition(*todos,
		func(item models.Todo) bool { return item.IsCompleted })

	tmpTodos := completedTodo

	todosIndex := len(notCompletedTodo) - 1

	orderIndex := order - 1

	if orderIndex < 0 || orderIndex > todosIndex {
		fmt.Println("Item not found")
		return
	}

	for index, todo := range notCompletedTodo {
		if index == orderIndex {
			updatedTodo := models.Todo{Name: todo.Name, IsCompleted: true}
			tmpTodos = append(tmpTodos, updatedTodo)
		} else {
			tmpTodos = append(tmpTodos, todo)
		}
	}

	*todos = tmpTodos
}

func getAllTodos() {
	utils.ClearScreen()
	fmt.Println("Wait to implement")
}

func ReadTodos(todos []models.Todo) {
	var actionType string
	fmt.Scanf("%s", &actionType)

	switch actionType {
	case "--all":
		{
			getAllTodos()
			break
		}
	default:
		fmt.Println("No argument " + actionType[2:] + "found!")
	}
}

func DeleteTodo(todos *[]models.Todo) {
	order := GetOrder()

	var tmpTodos []models.Todo

	notCompletedTodo := getNotCompletedTodos(*todos)

	for index, todo := range notCompletedTodo {
		if order != index+1 {
			tmpTodos = append(tmpTodos, todo)
		}
	}

	*todos = tmpTodos
}
