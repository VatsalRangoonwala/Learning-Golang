package main

import (
	"fmt"
	"sync"
)

func withMutex() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func withoutMutex() {
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			counter++
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}

func main() {
	fmt.Println("==========Without Mutex==========")
	withoutMutex()
	fmt.Println("==========With Mutex==========")
	withMutex()
}
