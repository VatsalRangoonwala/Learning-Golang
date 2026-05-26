package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

// --- Structs ---

type Order struct {
	ID        string    `json:"id"`
	Products  []Product `json:"products"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
}

type NewOrder struct {
	Products []Product `json:"products"`
	User     User      `json:"user"`
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

type Email struct {
	To      string
	Subject string
}

// --- Core Processing Logic ---

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
	return result, nil
}

// Helper to simulate processing a generic job
func processJob(id int, job int) {
	// fmt.Printf("Worker %d processing job %d\n", id, job)
	time.Sleep(10* time.Millisecond) // 1 second delay
}

// Helper to simulate sending an email
func processEmail(id int, email Email) {
	// fmt.Printf("[EmailWorker %d] Preparing to send to %s...\n", id, email.To)
	time.Sleep(20 * time.Millisecond) // 1 second delay
	// fmt.Printf("[EmailWorker %d] ✅ Successfully sent: '%s' to %s\n", id, email.Subject, email.To)
}

// Helper to simulate processing and saving an order
func processOrder(order NewOrder) {
	_, err := createOrder(order.Products, order.User)
	if err != nil {
		fmt.Println("Order Creation Error", err)
		return
	}
	time.Sleep(50 * time.Millisecond) // 2 second DB save simulation
	// fmt.Println("Successfully processed order ID: ", result.ID)
}

// --- Concurrent Worker Functions ---

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		processJob(id, job)
	}
}

func emailWorker(id int, emails <-chan Email, wg *sync.WaitGroup) {
	defer wg.Done()
	for email := range emails {
		processEmail(id, email)
	}
}

func orderWorker(orders <-chan NewOrder, wg *sync.WaitGroup) {
	defer wg.Done()
	for order := range orders {
		processOrder(order)
	}
}

// --- Execution Runners ---

// 1. Sequential Execution
func runSequential(jobsData []int, emailsData []Email, ordersData []NewOrder) time.Duration {
	start := time.Now()

	// Process Jobs (simulate running sequentially on main thread)
	for _, j := range jobsData {
		processJob(0, j)
	}

	// Process Emails
	for _, e := range emailsData {
		processEmail(0, e)
	}

	// Process Orders
	for _, o := range ordersData {
		processOrder(o)
	}

	return time.Since(start)
}

// 2. Concurrent Execution
func runConcurrent(jobsData []int, emailsData []Email, ordersData []NewOrder) time.Duration {
	start := time.Now()

	// Create channels with buffers large enough for all data
	jobs := make(chan int, len(jobsData))
	emails := make(chan Email, len(emailsData))
	results := make(chan NewOrder, len(ordersData))

	var wg sync.WaitGroup

	// Spawn Workers (2 Job, 2 Email, 5 Order)
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}
	for w := 1; w <= 2; w++ {
		wg.Add(1)
		go emailWorker(w, emails, &wg)
	}
	for w := 1; w <= 5; w++ {
		wg.Add(1)
		go orderWorker(results, &wg)
	}

	// Load Data into Channels
	for _, j := range jobsData {
		jobs <- j
	}
	for _, e := range emailsData {
		emails <- e
	}
	for _, o := range ordersData {
		results <- o
	}

	// Close channels to signal workers to stop when empty
	close(jobs)
	close(emails)
	close(results)

	// Wait for all workers across all pools to finish
	wg.Wait()

	return time.Since(start)
}

// --- Main ---

func main() {
	// 1. Prepare Data
	jobsPayload := []int{1, 2, 3, 4, 5}

	outbox := []Email{
		{"alice@example.com", "Welcome to our platform!"},
		{"bob@example.com", "Password Reset Link"},
		{"charlie@example.com", "Weekly Newsletter"},
		{"diana@example.com", "Your Invoice #9921"},
		{"eve@example.com", "Security Alert: New Login"},
	}

	orderPayload := []NewOrder{
		{[]Product{{"Shirt", 149.50, 2}}, User{"Vandan", "vandan1@boghara.com"}},
		{[]Product{{"Pant", 299.99, 1}, {"Belt", 99.99, 1}}, User{"Rahul", "rahul2@boghara.com"}},
		{[]Product{{"Shoes", 799.00, 1}, {"Socks", 49.00, 3}}, User{"Amit", "amit3@boghara.com"}},
		{[]Product{{"T-Shirt", 199.99, 2}}, User{"Karan", "karan4@boghara.com"}},
		{[]Product{{"Watch", 1299.00, 1}, {"Cap", 150.00, 2}}, User{"Neha", "neha5@boghara.com"}},
		{[]Product{{"Bag", 899.99, 1}}, User{"Riya", "riya6@boghara.com"}},
		{[]Product{{"Jeans", 599.50, 2}, {"Shirt", 249.00, 1}}, User{"Priya", "priya7@boghara.com"}},
		{[]Product{{"Laptop", 55999.00, 1}, {"Mouse", 799.00, 1}}, User{"Arjun", "arjun8@boghara.com"}},
		{[]Product{{"Keyboard", 1499.00, 1}}, User{"Manav", "manav9@boghara.com"}},
		{[]Product{{"Mobile", 18999.00, 1}, {"Cover", 499.00, 1}}, User{"Sneha", "sneha10@boghara.com"}},
		{[]Product{{"Charger", 699.00, 2}}, User{"Jay", "jay11@boghara.com"}},
		{[]Product{{"Earphones", 1299.00, 1}, {"PowerBank", 1999.00, 1}}, User{"Pooja", "pooja12@boghara.com"}},
		{[]Product{{"Notebook", 59.00, 5}}, User{"Krishna", "krishna13@boghara.com"}},
		{[]Product{{"Pen", 10.00, 10}, {"Pencil", 5.00, 12}}, User{"Vivek", "vivek14@boghara.com"}},
		{[]Product{{"Bottle", 199.00, 2}}, User{"Meera", "meera15@boghara.com"}},
		{[]Product{{"Chair", 2499.00, 1}, {"Table", 4999.00, 1}}, User{"Rohit", "rohit16@boghara.com"}},
		{[]Product{{"Fan", 1899.00, 1}}, User{"Anjali", "anjali17@boghara.com"}},
		{[]Product{{"AC", 32999.00, 1}, {"Stabilizer", 2499.00, 1}}, User{"Dev", "dev18@boghara.com"}},
		{[]Product{{"TV", 45999.00, 1}}, User{"Nitin", "nitin19@boghara.com"}},
		{[]Product{{"Remote", 499.00, 2}}, User{"Kishan", "kishan20@boghara.com"}},
		{[]Product{{"Mixer", 3499.00, 1}}, User{"Asha", "asha21@boghara.com"}},
		{[]Product{{"Cooker", 1999.00, 1}, {"Pan", 899.00, 2}}, User{"Suresh", "suresh22@boghara.com"}},
		{[]Product{{"Plate", 120.00, 6}}, User{"Bhavya", "bhavya23@boghara.com"}},
		{[]Product{{"Cup", 80.00, 4}, {"Bottle", 250.00, 2}}, User{"Hardik", "hardik24@boghara.com"}},
		{[]Product{{"Sofa", 25999.00, 1}}, User{"Tina", "tina25@boghara.com"}},
		{[]Product{{"Curtain", 999.00, 3}}, User{"Varun", "varun26@boghara.com"}},
		{[]Product{{"Bed", 18999.00, 1}, {"Pillow", 499.00, 4}}, User{"Nisha", "nisha27@boghara.com"}},
		{[]Product{{"Blanket", 1499.00, 2}}, User{"Parth", "parth28@boghara.com"}},
		{[]Product{{"Lamp", 699.00, 2}}, User{"Jatin", "jatin29@boghara.com"}},
		{[]Product{{"Clock", 499.00, 1}}, User{"Komal", "komal30@boghara.com"}},
		{[]Product{{"Helmet", 2499.00, 1}}, User{"Yash", "yash31@boghara.com"}},
		{[]Product{{"Bike Cover", 799.00, 1}, {"Gloves", 599.00, 2}}, User{"Dhruv", "dhruv32@boghara.com"}},
		{[]Product{{"Car Perfume", 299.00, 3}}, User{"Nakul", "nakul33@boghara.com"}},
		{[]Product{{"Tyre", 4999.00, 2}}, User{"Ankit", "ankit34@boghara.com"}},
		{[]Product{{"Oil", 899.00, 2}, {"Filter", 399.00, 1}}, User{"Milan", "milan35@boghara.com"}},
		{[]Product{{"Camera", 65999.00, 1}}, User{"Heena", "heena36@boghara.com"}},
		{[]Product{{"Tripod", 1999.00, 1}}, User{"Sahil", "sahil37@boghara.com"}},
		{[]Product{{"Mic", 3499.00, 1}, {"Stand", 899.00, 1}}, User{"Rakesh", "rakesh38@boghara.com"}},
		{[]Product{{"Speaker", 5499.00, 2}}, User{"Aarav", "aarav39@boghara.com"}},
		{[]Product{{"Router", 2499.00, 1}}, User{"Isha", "isha40@boghara.com"}},
		{[]Product{{"SSD", 4999.00, 1}}, User{"Om", "om41@boghara.com"}},
		{[]Product{{"RAM", 3299.00, 2}}, User{"Tushar", "tushar42@boghara.com"}},
		{[]Product{{"Processor", 18999.00, 1}, {"Motherboard", 12999.00, 1}}, User{"Viraj", "viraj43@boghara.com"}},
		{[]Product{{"Graphic Card", 45999.00, 1}}, User{"Zara", "zara44@boghara.com"}},
		{[]Product{{"Cabinet", 3499.00, 1}}, User{"Mahi", "mahi45@boghara.com"}},
		{[]Product{{"UPS", 5999.00, 1}}, User{"Ayaan", "ayaan46@boghara.com"}},
		{[]Product{{"Printer", 8999.00, 1}, {"Ink", 1499.00, 2}}, User{"Harsh", "harsh47@boghara.com"}},
		{[]Product{{"Scanner", 6499.00, 1}}, User{"Rupal", "rupal48@boghara.com"}},
		{[]Product{{"Projector", 28999.00, 1}}, User{"Chirag", "chirag49@boghara.com"}},
		{[]Product{{"Whiteboard", 1999.00, 1}}, User{"Payal", "payal50@boghara.com"}},
		{[]Product{{"Football", 799.00, 1}}, User{"Nirav", "nirav51@boghara.com"}},
		{[]Product{{"Bat", 1499.00, 1}, {"Ball", 199.00, 6}}, User{"Umesh", "umesh52@boghara.com"}},
		{[]Product{{"Gloves", 599.00, 1}}, User{"Harit", "harit53@boghara.com"}},
		{[]Product{{"Racket", 2499.00, 2}}, User{"Nency", "nency54@boghara.com"}},
		{[]Product{{"Shoes", 3999.00, 1}}, User{"Sejal", "sejal55@boghara.com"}},
		{[]Product{{"Track Pant", 999.00, 2}}, User{"Deep", "deep56@boghara.com"}},
		{[]Product{{"Gym Bag", 1499.00, 1}}, User{"Raj", "raj57@boghara.com"}},
		{[]Product{{"Dumbbell", 2499.00, 2}}, User{"Monika", "monika58@boghara.com"}},
		{[]Product{{"Yoga Mat", 799.00, 1}}, User{"Kajal", "kajal59@boghara.com"}},
		{[]Product{{"Cycle", 14999.00, 1}}, User{"Akash", "akash60@boghara.com"}},
		{[]Product{{"Milk", 60.00, 5}}, User{"Naman", "naman61@boghara.com"}},
		{[]Product{{"Bread", 40.00, 3}, {"Butter", 120.00, 1}}, User{"Reena", "reena62@boghara.com"}},
		{[]Product{{"Rice", 899.00, 1}}, User{"Sonal", "sonal63@boghara.com"}},
		{[]Product{{"Wheat", 699.00, 1}}, User{"Ajay", "ajay64@boghara.com"}},
		{[]Product{{"Oil", 1499.00, 1}, {"Sugar", 499.00, 2}}, User{"Tarun", "tarun65@boghara.com"}},
		{[]Product{{"Salt", 20.00, 5}}, User{"Bina", "bina66@boghara.com"}},
		{[]Product{{"Tea", 399.00, 2}}, User{"Mitesh", "mitesh67@boghara.com"}},
		{[]Product{{"Coffee", 599.00, 1}}, User{"Dimple", "dimple68@boghara.com"}},
		{[]Product{{"Biscuits", 30.00, 10}}, User{"Pankaj", "pankaj69@boghara.com"}},
		{[]Product{{"Chocolate", 99.00, 8}}, User{"Geeta", "geeta70@boghara.com"}},
		{[]Product{{"Perfume", 2499.00, 1}}, User{"Sagar", "sagar71@boghara.com"}},
		{[]Product{{"Face Wash", 299.00, 2}}, User{"Irfan", "irfan72@boghara.com"}},
		{[]Product{{"Soap", 49.00, 6}, {"Shampoo", 249.00, 2}}, User{"Farhan", "farhan73@boghara.com"}},
		{[]Product{{"Toothpaste", 120.00, 3}}, User{"Aaliya", "aaliya74@boghara.com"}},
		{[]Product{{"Comb", 50.00, 2}}, User{"Zubin", "zubin75@boghara.com"}},
		{[]Product{{"Face Cream", 399.00, 1}}, User{"Rina", "rina76@boghara.com"}},
		{[]Product{{"Lipstick", 799.00, 2}}, User{"Mona", "mona77@boghara.com"}},
		{[]Product{{"Nail Polish", 199.00, 4}}, User{"Sana", "sana78@boghara.com"}},
		{[]Product{{"Hair Dryer", 1899.00, 1}}, User{"Ritu", "ritu79@boghara.com"}},
		{[]Product{{"Trimmer", 2499.00, 1}}, User{"Faiz", "faiz80@boghara.com"}},
		{[]Product{{"Toy Car", 499.00, 3}}, User{"Kabir", "kabir81@boghara.com"}},
		{[]Product{{"Puzzle", 299.00, 2}}, User{"Aryan", "aryan82@boghara.com"}},
		{[]Product{{"Doll", 799.00, 1}}, User{"Kiara", "kiara83@boghara.com"}},
		{[]Product{{"Lego Set", 2499.00, 1}}, User{"Vivaan", "vivaan84@boghara.com"}},
		{[]Product{{"Story Book", 199.00, 5}}, User{"Sara", "sara85@boghara.com"}},
		{[]Product{{"Color Box", 149.00, 4}}, User{"Tanvi", "tanvi86@boghara.com"}},
		{[]Product{{"School Bag", 1299.00, 1}}, User{"Yuvi", "yuvi87@boghara.com"}},
		{[]Product{{"Lunch Box", 399.00, 2}}, User{"Mira", "mira88@boghara.com"}},
		{[]Product{{"Water Bottle", 299.00, 2}}, User{"Devansh", "devansh89@boghara.com"}},
		{[]Product{{"Notebook", 89.00, 10}}, User{"Aditi", "aditi90@boghara.com"}},
		{[]Product{{"Drone", 45999.00, 1}}, User{"Rudra", "rudra91@boghara.com"}},
		{[]Product{{"VR Headset", 25999.00, 1}}, User{"Vihaan", "vihaan92@boghara.com"}},
		{[]Product{{"Smart Watch", 4999.00, 2}}, User{"Kruti", "kruti93@boghara.com"}},
		{[]Product{{"Tablet", 21999.00, 1}}, User{"Jenil", "jenil94@boghara.com"}},
		{[]Product{{"Stylus", 1999.00, 1}}, User{"Niyati", "niyati95@boghara.com"}},
		{[]Product{{"Gaming Mouse", 3499.00, 1}}, User{"Pratik", "pratik96@boghara.com"}},
		{[]Product{{"Gaming Chair", 14999.00, 1}}, User{"Bhargav", "bhargav97@boghara.com"}},
		{[]Product{{"Monitor", 17999.00, 2}}, User{"Het", "het98@boghara.com"}},
		{[]Product{{"HDMI Cable", 499.00, 3}}, User{"Viral", "viral99@boghara.com"}},
		{[]Product{{"Webcam", 2999.00, 1}, {"Ring Light", 1999.00, 1}}, User{"Anaya", "anaya100@boghara.com"}},
	}

	// 2. Measure Executions
	fmt.Println("Starting Sequential Processing... (This will take a while)")
	seqTime := runSequential(jobsPayload, outbox, orderPayload)
	fmt.Printf("✅ Sequential Processing Finished in: %v\n\n", seqTime)

	fmt.Println("Starting Concurrent Processing...")
	concTime := runConcurrent(jobsPayload, outbox, orderPayload)
	fmt.Printf("✅ Concurrent Processing Finished in: %v\n\n", concTime)

	// 3. Print Results
	fmt.Println("========================================")
	fmt.Println("          PERFORMANCE SUMMARY           ")
	fmt.Println("========================================")
	fmt.Printf("Sequential Time: %v\n", seqTime)
	fmt.Printf("Concurrent Time: %v\n", concTime)
	fmt.Printf("Time Saved:      %v\n", seqTime-concTime)
	fmt.Printf("Speedup:         %.2fx faster\n", float64(seqTime)/float64(concTime))
	fmt.Println("========================================")
}