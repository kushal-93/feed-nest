package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")

	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found. Exiting.")
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		MaxAge:         300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server staring on port=%v", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
