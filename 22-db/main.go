package main

import (
	"net/http"
	"project/db"
	"project/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	const version = "/v1"
	db.ConnectMongo()

	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Get("/", handlers.HomeHandler)
		r.Get("/health", handlers.HealthHandler)

		r.Route(version, func(r chi.Router) {
			r.Get("/user", handlers.UserHandler)
			r.Get("/products", handlers.ProductsHandler)
			r.Post("/order", handlers.OrderHandler)
			r.Get("/orders",handlers.GetOrdersHandler)
			r.Get("/order/{id}", handlers.GetOrderByIDHandler)
		})
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
