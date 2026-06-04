package main

import (
	"fmt"
	"sync"
	"time"
)

type Product struct {
	Name  string
	Price float64
}

var wg sync.WaitGroup
var mu sync.Mutex
var rmu sync.RWMutex
var products = map[string]Product{
	"shirt": {Name: "Shirt", Price: 299},
	"shoes": {Name: "Shoes", Price: 999},
}

func withMutex() {
	var counter int

	for range 1000 {

		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
		})
	}

	wg.Wait()

	fmt.Println(counter)
}

func withoutMutex() {
	var counter int

	for range 1000 {

		wg.Go(func() {

			counter++
		})
	}

	wg.Wait()

	fmt.Println(counter)
}
func stockCheck() {
	var stock int = 100

	for range 100 {

		wg.Go(func() {
			mu.Lock()
			defer mu.Unlock()
			stock--
		})
	}
	wg.Wait()

	fmt.Println(stock)
}

func getProduct(name string) {
	defer wg.Done()
	rmu.RLock()
	defer rmu.RUnlock()

	fmt.Println("Reader started")
	time.Sleep(2 * time.Second)
	data, ok := products[name]
	if !ok {
		fmt.Println(name, " is not in system!")
	} else {
		fmt.Println("Reader finished:", data)
	}
}

func updatePrice(name string, price float64) {
	defer wg.Done()
	fmt.Println("Writer waiting...")
	rmu.Lock()
	defer rmu.Unlock()

	fmt.Println("Writer acquired lock")
	p, ok := products[name]
	if !ok {
		fmt.Println(name, " is not in system!")
	} else {
		p.Price = price
		products[name] = p
		fmt.Println("Writer updated price: ", p)
	}
}

func multiReader() {
	for range 5 {
		wg.Add(1)
		go getProduct("shirt")
	}
	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	go updatePrice("shirt", 399)
	wg.Wait()
	fmt.Println("Done")
}

func main() {
	fmt.Println("==========Without Mutex==========")
	withoutMutex()
	fmt.Println("==========With Mutex==========")
	withMutex()
	fmt.Println("==========Stock Check==========")
	stockCheck()
	fmt.Println("==========RWMutex==========")
	multiReader()
}
