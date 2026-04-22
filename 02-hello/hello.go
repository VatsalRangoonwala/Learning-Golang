package main

import (
	"fmt"
	"log"

	"greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("Vatsal")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
