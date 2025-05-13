package main

import (
	"log"
	"net/http"

	"github.com/samfeese/chat/data"
)

func main() {
	session := data.NewCassandraSession()
	defer session.Close()

	log.Println("Connected to Cassandra")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
