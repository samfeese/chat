package main

import (
	"log"
	"net/http"

	"github.com/samfeese/chat/data"
)

func main() {
	session := data.NewBootstrapSession()
	defer session.Close()

	if err := data.BootstrapCassandra(session); err != nil {
		log.Fatalf("Bootstrap failed: %v", err)
	}
	session.Close()

	appSession := data.NewCassandraSession()
	defer appSession.Close()

	log.Println("Connected to Cassandra")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { 
		w.Write([]byte("OK"))
	})
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
