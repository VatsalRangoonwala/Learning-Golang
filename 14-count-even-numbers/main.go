package main

import "fmt"

func main() {
	n := 20
	count := 0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			count += 1
		}
	}

	fmt.Println("Count of Even numbers between 1 to 20 is", count)
}
