package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/bhemi28/go-rss-aggregator/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Error getting PORT environment variable", os.Getenv("PORT"))
		return
	}

	dbString := os.Getenv("DB_URL")
	if dbString == "" {
		log.Fatal("Error getting DB_URL environment variable", os.Getenv("DB_URL"))
		return
	}

	conn, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
		return
	}

	dbConfig := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		}),
	)

	v1Router := chi.NewRouter()
	v1Router.Get("/json", jsonHandler)
	v1Router.Get("/error", errorHandler)
	v1Router.Post("/user", dbConfig.handlerCreateUser)
	v1Router.Get("/user", dbConfig.middlewareAuthHandler(dbConfig.handlerGetUser))
	v1Router.Post("/feed", dbConfig.middlewareAuthHandler(dbConfig.handlerCreateFeed))
	v1Router.Post("/feed/add", dbConfig.middlewareAuthHandler(dbConfig.addFeedToUser))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Starting server on port %s", port)
	error := server.ListenAndServe()
	if error != nil {
		log.Fatal("Error starting server:", error)
		return
	}

}
