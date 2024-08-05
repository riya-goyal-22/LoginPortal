package ToDo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ToDoItem struct {
	Task string
}

var todoList = []ToDoItem{}

func AddTodo() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("")
	fmt.Print("Enter a new task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	todoList = append(todoList, ToDoItem{Task: task})
	fmt.Println("Task added.")
}

func ListTodos() {
	if len(todoList) == 0 {
		fmt.Println("")
		fmt.Println("No tasks in the to-do list.")
		return
	}
	fmt.Println("")
	fmt.Println("To-Do List:")
	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item.Task)
	}
}

func DeleteTodo() {
	ListTodos()
	fmt.Print("Enter the task number to delete: ")
	var index int
	_, err := fmt.Scanln(&index)
	if err != nil || index < 1 || index > len(todoList) {
		fmt.Println("Invalid task number.")
		return
	}
	todoList = append(todoList[:index-1], todoList[index:]...)
	fmt.Println("Task deleted.")
}
