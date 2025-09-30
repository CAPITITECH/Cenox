package main

import (
	"cenox/router"
	"log"
	"net/http"
)

func main() {
	r := router.NewRouter()
	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}