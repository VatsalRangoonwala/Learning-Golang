package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var activeOrders int64
var wg sync.WaitGroup

func createOrder() {
	defer wg.Done()
	atomic.AddInt64(&activeOrders, 1)
}

func main() {
	for range 100 {
		wg.Add(1)
		go createOrder()
	}
	wg.Wait()
	finalCount := atomic.LoadInt64(&activeOrders)
	fmt.Println("Active orders: ", finalCount)
}
