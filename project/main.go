package main

import (
	"example/hello/project/internal/database"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	_ "github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type ApiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	db_port := os.Getenv("DATABASE_URL")

	if db_port == "" {
		log.Fatal("DATABASE_URL is not found in the enviroment")
	}

	fmt.Println("Port: ", port)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadliness)
	v1Router.Get("/err", handlerError)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server starting no port %v ", port)
	err := srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
