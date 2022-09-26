package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!!"))
	})

	log.Printf("Server running on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
