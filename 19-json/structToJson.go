package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{Name: "Vatsal", Email: "vatsal@gmail.com"}

	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	result, err := jsonToStruct(`{"name":"Vatsal","email":"test@gmail.com"}`)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	p1 := Product{"Shirt", 199.50}
	p2 := Product{"Shoes", 349.80}

	res, err := productApiResponse(p1, p2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	o1 := Order{p1, user}

	orderRes, err := orderToJson(o1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	finalRes, err := createResponse(true, "Order Created", o1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
	fmt.Println(result)
	fmt.Println(string(res))
	fmt.Println(string(orderRes))
	fmt.Println(string(finalRes))
}
