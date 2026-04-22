package main

import (
	"encoding/json"
	"errors"
)

func jsonToStruct(jsonData string) (User, error) {
	var user User
	if jsonData == "" {
		return user, errors.New("JSON required")
	}
	err := json.Unmarshal([]byte(jsonData), &user)
	return user, err
}
