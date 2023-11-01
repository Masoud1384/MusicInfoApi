package main

import (
	"github.com/go-chi/chi"
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
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}
	srv.ListenAndServe()

}
