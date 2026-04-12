package main

import (
	"fmt"
	"log"
	"net/http"

	"Users/database"
	"Users/handlers"
)

func main() {
	
	database.Init()

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/static/", handlers.StaticHandlers)
	fmt.Println("http://localhost:8081")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
