package main

import (
	"LoginPortal/Login"
	"LoginPortal/SignUp"
	"LoginPortal/utils"
	"errors"
	"fmt"
	"log"
	"os"
)

//const FileName = "LoginDetails.txt"

var option int

type Person struct {
	Name string
	Pass string
}

func main() {

	var gender string
	p := Person{}
	fmt.Println("Are you a female? Say y for YES and n for NO")
	_, err := fmt.Scanln(&gender)
	if err != nil {
		return
	}
	if gender != "y" {
		err = errors.New("you are not a valid user to enter")
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Enter your name")
	_, err = fmt.Scanln(&p.Name)
	if err != nil {
		return
	}

	for !utils.IsPassCorrect(p.Pass) {
		fmt.Println("Enter a strong password")
		_, err = fmt.Scanln(&p.Pass)
		if err != nil {
			return
		}
	}

	fmt.Println("Enter 1 for SignUp and 2 for login and 3 to exit")
	_, err = fmt.Scanln(&option)
	if err != nil {
		return
	}

	//switch function
	switch option {
	case 1:
		SignUp.SignUpCheck(p.Name, p.Pass)
	case 2:
		Login.LoginCheck(p.Name, p.Pass)
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid option")
	}

}
