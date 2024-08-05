package SignUp

import (
	"LoginPortal/Login"
	"fmt"
	"os"
)

const FileName = "LoginDetails.txt"

func SignUpCheck(name string, pass string) {

	if Login.Login(name) {
		fmt.Println("You are already a user,no need to SignUp just Login")

	} else {
		SignUp(name, pass)
	}

}

func SignUp(name string, pass string) {
	f, err := os.OpenFile(FileName, os.O_RDWR|os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	content := fmt.Sprintf("%s %s\n", name, pass)
	_, err = f.Write([]byte(content))
	if err != nil {
		panic(err)
	}
	fmt.Println("Successful signUp")
}
