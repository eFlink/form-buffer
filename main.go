package main

import (
	"fmt"
	"net/http"

	"github.com/eFlink/form-buffer/handlers"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", handlers.PartialFormHandler)

	fmt.Println("Server running on port :8080")
	http.ListenAndServe(":8080", mux)
}
