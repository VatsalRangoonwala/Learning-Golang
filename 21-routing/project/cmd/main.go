package main

import (
	"net/http"
	"project/handlers"
	"project/middleware"

	"github.com/go-chi/chi/v5"
)

func main() {
	const version = "/v1"

	r := chi.NewRouter()

	r.Use(middleware.LogMiddleware)

	r.Route("/api", func(r chi.Router) {
		r.Get("/", handlers.HomeHandler)
		r.Get("/health", handlers.HealthHandler)

		r.Route(version, func(r chi.Router) {
			r.Get("/user", handlers.UserHandler)
			r.Get("/products", handlers.ProductsHandler)
			r.Post("/order", handlers.OrderHandler)
		})
	})
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
