package main

import "fmt"

func main() {
	age := 20
	hasID := true

	if age >= 18 && hasID {
		fmt.Println("Access Granted")
	} else {
		fmt.Println("Access Denied")
	}
}
