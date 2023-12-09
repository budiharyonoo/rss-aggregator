package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error parse .env file: %s", err)
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("PORT is not bound in the .env file")
		return
	}

	fmt.Println("Port:", port)

	// Init Chi Router
	router := chi.NewRouter()

	// Router
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Create semantic router
	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerError)
	router.Mount("/v1", v1Router)

	// Init server
	server := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	errSrv := server.ListenAndServe()
	if errSrv != nil {
		log.Fatalln(errSrv)
		return
	}

	log.Printf("Server starting on port: %s", port)
}
