package main

import (
	"database/sql"
	"github.com/budiharyonoo/rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

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

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatalln("DB_URL is not bound in the .env file")
		return
	}

	// Open DB Connection
	dbConn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("DB connection error", err)
		return
	}

	apiCfg := apiConfig{DB: database.New(dbConn)}

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

	// v1/user
	v1Router.Post("/user", apiCfg.handlerCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	// v1/feed
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))

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
