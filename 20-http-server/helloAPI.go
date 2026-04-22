package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello API")
	w.Write([]byte("Hello API"))
}

func main() {
	http.Handle("/", logMiddleware(http.HandlerFunc(homeHandler)))
	http.Handle("/user", logMiddleware(http.HandlerFunc(userHandler)))
	http.Handle("/products", logMiddleware(http.HandlerFunc(productsHandler)))
	http.Handle("/order", logMiddleware(http.HandlerFunc(orderHandler)))
	http.Handle("/health", logMiddleware(http.HandlerFunc(healthHandler)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
