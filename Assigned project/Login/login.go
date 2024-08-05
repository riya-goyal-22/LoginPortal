package Login

import (
	"LoginPortal/ToDo"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FileName = "LoginDetails.txt"

func find(name string) bool {
	f, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), name) {
			return true
		}
	}
	return false

}

func findWithPass(name string, pass string) bool {
	f, err := os.Open(FileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Splits on newlines by default.
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), name) {
			if strings.Contains(scanner.Text(), pass) {
				return true
			}

		}
	}
	return false

}

func Login(name string) bool {
	if find(name) == true {
		return true
	} else {
		return false
	}
}

func LoginCheck(name string, pass string) {
	if findWithPass(name, pass) {
		fmt.Println("Successful Login")
		Todo()

	} else if find(name) {
		fmt.Println("Wrong Password")
	} else {
		fmt.Println("SignUp First")
	}

}

func Todo() {

	fmt.Println("Welcome to the To-Do CLI")

	for {
		fmt.Println("\n1. Add a task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Delete a task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		reader := bufio.NewReader(os.Stdin)
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			ToDo.AddTodo()
		case "2":
			ToDo.ListTodos()
		case "3":
			ToDo.DeleteTodo()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}

}
