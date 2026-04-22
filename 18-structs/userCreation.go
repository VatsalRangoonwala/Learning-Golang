package main

import "errors"

func createUser(name, email string, age int) (User, error) {
	if email == "" {
		return User{}, errors.New("Email is required")
	}
	if age <= 0 {
		return User{}, errors.New("Invalid Age")
	}
	return User{Name: name, Email: email, Age: age}, nil
}
