package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        string    `json:"id"`
	Products  []Product `json:"products"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func logOrder(o Order) {
	jsonData, err := json.Marshal(o)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	time.Sleep(time.Second)
	fmt.Println(string(jsonData))
}

func createOrder(products []Product, user User) (Order, error) {
	if len(products) == 0 {
		return Order{}, errors.New("no products provided")
	}

	for _, p := range products {
		if len(strings.TrimSpace(p.Name)) < 1 || p.Price <= 0 || p.Quantity <= 0 {
			return Order{}, errors.New("Invalid product data")
		}
	}

	if len(strings.TrimSpace(user.Name)) < 1 || len(strings.TrimSpace(user.Email)) < 1 {
		return Order{}, errors.New("Invalid User data")
	}

	var total float64
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}

	result := Order{
		ID:        uuid.New().String(),
		Products:  products,
		User:      user,
		CreatedAt: time.Now(),
		Total:     total,
		Status:    "pending",
	}
	time.Sleep(time.Second)
	return result, nil
}

func printTask(name string) {
	for i := 1; i <= 2; i++ {
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
}

func sendEmail(content string) {
	time.Sleep(time.Second)
	fmt.Println(content)
}

func measureConcurrent() {
	fmt.Println("--- Starting Concurrent Processing ---")
	var wg sync.WaitGroup
	p1 := Product{"Shirt", 149.50, 2}
	p2 := Product{"Pant", 179.50, 2}
	p3 := Product{"shoes", 179.50, 2}
	p4 := Product{"belt", 179.50, 2}
	p5 := Product{"perfume", 179.50, 2}
	u1 := User{"Vandan", "vandan@boghara.com"}
	u2 := User{"Vatsal", "vatsal@wala.com"}
	u3 := User{"Ahemad", "ahemad@pappu.com"}

	start := time.Now()
	wg.Add(8)
	go func() {
		defer wg.Done()
		o1, err := createOrder([]Product{p1, p2}, u1)
		if err != nil {
			fmt.Println("order creation failed:", err)
			return
		}
		logOrder(o1)
	}()
	go func() {
		defer wg.Done()
		o2, err := createOrder([]Product{p3, p4}, u2)
		if err != nil {
			fmt.Println("order creation failed:", err)
			return
		}
		logOrder(o2)
	}()
	go func() {
		defer wg.Done()
		o3, err := createOrder([]Product{p5}, u3)
		if err != nil {
			fmt.Println("order creation failed:", err)
			return
		}
		logOrder(o3)
	}()
	go func() {
		defer wg.Done()
		printTask("Task A")
	}()
	go func() {
		defer wg.Done()
		printTask("Task B")
	}()
	go func() {
		defer wg.Done()
		sendEmail("Welcome")
	}()
	go func() {
		defer wg.Done()
		sendEmail("invoice email")
	}()
	go func() {
		defer wg.Done()
		sendEmail("notification email")
	}()

	wg.Wait()

	duration := time.Since(start)

	fmt.Println("Concurrent execution finished")
	fmt.Printf("Execution time: %v\n", duration)
}

func measureSequential() {
	fmt.Println("--- Starting Sequential Processing ---")
	p1 := Product{"Shirt", 149.50, 2}
	p2 := Product{"Pant", 179.50, 2}
	p3 := Product{"shoes", 179.50, 2}
	p4 := Product{"belt", 179.50, 2}
	p5 := Product{"perfume", 179.50, 2}
	u1 := User{"Vandan", "vandan@boghara.com"}
	u2 := User{"Vatsal", "vatsal@wala.com"}
	u3 := User{"Ahemad", "ahemad@pappu.com"}

	start := time.Now()

	o1, err := createOrder([]Product{p1, p2}, u1)
	if err != nil {
		fmt.Println("order creation failed:", err)
		return
	}
	logOrder(o1)

	o2, err := createOrder([]Product{p3, p4}, u2)
	if err != nil {
		fmt.Println("order creation failed:", err)
		return
	}
	logOrder(o2)

	o3, err := createOrder([]Product{p5}, u3)
	if err != nil {
		fmt.Println("order creation failed:", err)
		return
	}
	logOrder(o3)

	printTask("Task A")
	printTask("Task B")
	sendEmail("Welcome")
	sendEmail("invoice email")
	sendEmail("notification email")
	duration := time.Since(start)
	fmt.Println("Sequential execution finished")
	fmt.Printf("Execution time: %v\n", duration)
}

func main() {
	measureConcurrent()
	measureSequential()
}
