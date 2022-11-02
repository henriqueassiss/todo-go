package main

import (
	"net/http"

	"github.com/henriqueassiss/todo-go/routes"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvVars()
	routes.Routes()
	http.ListenAndServe(":8000", nil)
}

func loadEnvVars() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}
