package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Order represents a basic job payload
type Order struct {
	ID            int
	EstimatedTime time.Duration // How long this order actually takes to process
}

// 1. WORKER POOL: This function runs as multiple concurrent goroutines
func orderWorker(id int, ctx context.Context, orders <-chan Order, wg *sync.WaitGroup) {
	defer wg.Done() // 3. GRACEFUL SHUTDOWN: Notify WaitGroup when this worker completely exits

	for {
		select {
		// 2. CONTEXT CANCELLATION: The global kill switch
		case <-ctx.Done():
			fmt.Printf("Worker %d: Received shutdown signal. Stopping.\n", id)
			return

		// Pulling an order from the shared queue
		case order := <-orders:
			fmt.Printf("Worker %d: Picked up Order %d\n", id, order.ID)

			// Process the order with timeout awareness
			processOrder(ctx, id, order)
		}
	}
}

// 4. TIMEOUT-AWARE PROCESSING: Handles a single order
func processOrder(parentCtx context.Context, workerID int, order Order) {
	// Simulate doing the actual work for this order
	select {
	case <-time.After(order.EstimatedTime):
		// The work finished before the 2-second timeout
		fmt.Printf("   -> Worker %d: Successfully finished Order %d\n", workerID, order.ID)

	case <-parentCtx.Done():
		// The 2-second timeout hit OR the parent context was cancelled!
		fmt.Printf("   -> [ERROR] Worker %d: Order %d FAILED (%v)\n", workerID, order.ID, parentCtx.Err())
	}
}

func process(ctx context.Context) {

	select {

	case <-time.After(3 * time.Second):
		fmt.Println("Process finished")

	case <-ctx.Done():
		fmt.Println("Process cancelled", ctx.Err())
	}
}

func worker(ctx context.Context, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		select {

		case <-ctx.Done():
			fmt.Println("Worker shutting down", ctx.Err())
			return

		case job := <-jobs:
			fmt.Println("Processing", job)
		}
	}
}

func fetchUserFromDB(ctx context.Context, userID int) (string, error) {
	fmt.Println("   [Repository] Executing DB query...")

	// Simulate a slow database query that takes 2 seconds
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("   [Repository] DB query successful!")
		return fmt.Sprintf("User_%d_Data", userID), nil

	case <-ctx.Done():
		// If the context is cancelled before the 2 seconds are up, abort!
		fmt.Println("   [Repository] DB query ABORTED due to context cancellation!")
		return "", ctx.Err()
	}
}

// --- 2. SERVICE LAYER ---
// Handles business logic
func getUserService(ctx context.Context, userID int) (string, error) {
	fmt.Println("  [Service] Processing business logic...")

	// The service just passes the context straight down to the repository
	data, err := fetchUserFromDB(ctx, userID)
	if err != nil {
		return "", err
	}

	return data, nil
}

// --- 1. HANDLER LAYER ---
// Entry point for the request
func handleUserRequest(ctx context.Context, userID int) {
	fmt.Println("[Handler] Request received for User ID:", userID)

	// The handler passes the context down to the service
	result, err := getUserService(ctx, userID)
	if err != nil {
		fmt.Println("[Handler] Error returning response to user:", err)
		return
	}

	fmt.Println("[Handler] Successfully returned to user:", result)
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	jobs := make(chan int, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(ctx, jobs, &wg)
	go process(ctx)

	// Send some dummy jobs to verify the worker is listening
	jobs <- 101
	jobs <- 102
	jobs <- 103

	fmt.Println("Main: Waiting 3 seconds before cancelling...")
	time.Sleep(3 * time.Second)

	fmt.Println("Main: Cancelling context manually now!")
	cancel()
	// Wait for the worker to finish shutting down before exiting main

	fmt.Print("\n===========================\n\n")
	webCtx, webCancel := context.WithTimeout(context.Background(), 5*time.Second)
	time.Sleep(3 * time.Second)
	webCancel()
	handleUserRequest(webCtx, 420)

	fmt.Print("\n===========================\n\n")
	orderCtx, cancelOrder := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelOrder()
	orders := make(chan Order, 10)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go orderWorker(i, orderCtx, orders, &wg)
	}

	orders <- Order{ID: 101, EstimatedTime: 1 * time.Second}
	orders <- Order{ID: 102, EstimatedTime: 3 * time.Second}
	orders <- Order{ID: 103, EstimatedTime: 1 * time.Second}
	orders <- Order{ID: 104, EstimatedTime: 1 * time.Second}

	wg.Wait()
	fmt.Println("Main: Program exited cleanly.")
}
