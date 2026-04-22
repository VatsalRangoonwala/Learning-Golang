package main

import "fmt"

func main() {
	if num := 11; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
