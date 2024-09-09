package main

import (
	"fmt"
	"net/http"
	"zephyr/handlers"
)

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", handlers.Handler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("Shutting down server ...")
}
