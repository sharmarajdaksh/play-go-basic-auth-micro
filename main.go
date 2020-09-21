package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sharmarajdaksh/basic-auth-microservice/db"
	"github.com/sharmarajdaksh/basic-auth-microservice/middleware"

	"github.com/sharmarajdaksh/basic-auth-microservice/config"
	"github.com/sharmarajdaksh/basic-auth-microservice/handlers"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config. Exiting.")
	}

	if err := db.InitializeDB(); err != nil {
		log.Fatal("Failed to create database connection. Exiting.")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/verify", middleware.WithMiddleware(handlers.Verify))
	mux.HandleFunc("/register", middleware.WithMiddleware(handlers.Register))
	mux.HandleFunc("/delete", middleware.WithMiddleware(handlers.Delete))

	// http.Handle("/graphql", handlers.GraphQLHandler)

	s := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", config.C.Global.ListenPort),
		Handler: mux,
	}

	log.Fatal(
		s.ListenAndServe())
}
