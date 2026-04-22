package main

import "fmt"

type User struct {
	Name  string
	Email string
	Age   int
}

func main() {

	firstUser := User{
		Name:  "Vatsal",
		Email: "vatsal@gmail.com",
		Age:   20,
	}

	secondUser := User{
		Name:  "Vandan",
		Email: "vandan@gmail.com",
		Age:   21,
	}

	p1 := Product{
		Name:     "Shirt",
		Price:    249.50,
		Quantity: 5,
	}
	O1 := Order{
		Product: p1,
		User:    firstUser,
	}

	total, err := getTotal(p1)
	if err != nil {
		fmt.Println(err)
	}

	fp, err := p1.finalPrice(10)
	if err != nil {
		fmt.Println(err)
	}

	user, err := createUser("Momin", "momin@gmail.com", 20)
	if err != nil {
		fmt.Println(err)
	}

	summary, err := O1.summary()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(firstUser)
	fmt.Println(secondUser)
	fmt.Println(total)
	fmt.Println(fp)
	fmt.Println(user)
	fmt.Println(summary)
}
