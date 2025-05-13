package data

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

func NewCassandraSession() *gocql.Session {
	host := os.Getenv("CASSANDRA_HOST")
	keyspace := os.Getenv("CASSANDRA_KEYSPACE")

	cluster := gocql.NewCluster(host) // Update with your Cassandra host(s)
	cluster.Keyspace = keyspace       // Make sure this keyspace exists
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra: %v", err)
	}
	return session
}
