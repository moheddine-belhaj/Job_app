package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/moheddine-belhaj/Job_app/pkg/routes"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting...")
	r := mux.NewRouter()

	// Register the routes for the book store
	routes.RegisterJobRoutes(r)

	// Use CORS middleware to allow requests from the Vue.js app
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the router with CORS middleware
	handler := c.Handler(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":9010", handler))
}
