package main

import (
	"errors"
	"fmt"
	"strings"
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

func sendData(ch chan string) {
	ch <- "Hello from goroutine"
}
func sendTotal(totalCh chan float64, products []Product) {
	var total float64
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}
	totalCh <- total
}

func sendEmail(contentCh chan string, content string) {
	contentCh <- content
}

func createOrder(products []Product, user User) (Order, error) {
	ch := make(chan float64)
	
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

	go sendTotal(ch, products)
	
	result := Order{
		ID:        uuid.New().String(),
		Products:  products,
		User:      user,
		CreatedAt: time.Now(),
		Total:     <-ch,
		Status:    "pending",
	}
	return result, nil
}

func main() {
	p1 := Product{"Shirt", 149.50, 2}
	p2 := Product{"Pant", 179.50, 2}
	p3 := Product{"shoes", 179.50, 2}
	p4 := Product{"belt", 179.50, 2}
	p5 := Product{"perfume", 179.50, 2}
	u1 := User{"Vandan", "vandan@boghara.com"}
	u2 := User{"Vatsal", "vatsal@wala.com"}
	// u3 := User{"Ahemad", "ahemad@pappu.com"}
	taskCh := make(chan string)
	errorCh := make(chan error)
	bufferCh := make(chan int, 3)

	bufferCh<-10
	bufferCh<-20
	bufferCh<-30

	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)
	fmt.Println(<-bufferCh)

	close(bufferCh)


	go func() {
		sendData(taskCh)
	}()

	go func() {
		o1, err := createOrder([]Product{p1, p2}, u1)
		if err != nil {
			errorCh <- err
			return
		}
		fmt.Println(o1)
		taskCh <- "Task completed"
	}()
	go func() {
		o2, err := createOrder([]Product{p3, p4, p5}, u2)
		if err != nil {
			errorCh <- err
			return
		}
		fmt.Println(o2)
		taskCh <- "Task completed"
	}()
	go func() {
		sendEmail(taskCh, "welcome")
	}()

	for i := 1; i <= 4; i++ {
		select {
		case err := <-errorCh:
			// If ANY goroutine sends an error, this block runs.
			fmt.Println("Task failed with error:", err)
			return // Exits the program immediately

		case msg := <-taskCh:
			// If a goroutine succeeds, this block runs.
			fmt.Println(msg)
		}
	}
	close(errorCh)
	close(taskCh)
}
