package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("somthing went wrong :", err)
	}
	port := os.Getenv("PORT")
	host := os.Getenv("HOST")
	if port == "" || host == "" {
		log.Fatal("there is no host or port for running the project")
		os.Exit(1)
	}
	router := chi.NewRouter()
	router.Use(
		cors.Handler(cors.Options{
			AllowCredentials: false,
			AllowedHeaders:   []string{"*"},
			AllowedOrigins:   []string{"https://", "http://"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			ExposedHeaders:   []string{"LINKS"},
			MaxAge:           200,
		}),
	)
	v1router := chi.NewRouter()
	v1router.Get("/testget", responseHealth)
	router.Mount("/v1", v1router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	srv.ListenAndServe()

}
