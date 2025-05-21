package data

import (
	"log"
	"os"
	"time"

	"github.com/gocql/gocql"
)

func NewCassandraSession() *gocql.Session {
	host := os.Getenv("CASSANDRA_HOST")
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")

	cluster := gocql.NewCluster(host) // Update with your Cassandra host(s)
	cluster.Keyspace = keyspace       // Make sure this keyspace exists
	cluster.Consistency = gocql.Quorum

	var session *gocql.Session
	var err error

	for i := 0; i < 10; i++ {
		session, err = cluster.CreateSession()
		if err == nil {
        log.Println("Connected to Cassandra")
				return session
    }


		log.Printf("Attempt %d: Failed to connect to Cassandra: %v", i+1, err)
		time.Sleep(5 * time.Second)
	}

	log.Fatalf("Failed to connect to Cassandra: %v", err)
	return nil

}
