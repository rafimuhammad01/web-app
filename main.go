package main

import (
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = DEFAULT_PORT
	}

	server := http.Server{
		Addr: "0.0.0.0:" + port,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	log.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
