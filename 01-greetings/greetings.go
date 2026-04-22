package greetings

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
)

func Hello(name string) (string, error) {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
