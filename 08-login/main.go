package main

import "fmt"

func main() {
	username := "vatsal"
	password := "123"
	storedUsername := "admin"
	storedPassword := "1234"

	if username == storedUsername && password == storedPassword {
		fmt.Println("Login Successful")
	} else {
		fmt.Println("Invalid Credentials")
	}
}
