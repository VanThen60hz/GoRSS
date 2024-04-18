package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Get Port .env
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Cors Route
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Router v1
	v1Router := chi.NewRouter()

	v1Router.Get("/heath", handlerReadiness)
	v1Router.Get("/err", handleErr)

	router.Mount("/v1", v1Router)

	srv := http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	// Log port
	log.Printf("Server starting on port %v", portString)

	// Listen serve
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
