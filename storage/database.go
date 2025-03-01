package storage

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

// Session is the global Cassandra session
var Session *gocql.Session

// ConnectDatabase initializes the Cassandra connection
func ConnectDatabase() {
	cluster := gocql.NewCluster("127.0.0.1") // Change this if Cassandra runs elsewhere
	cluster.Keyspace = "go_crud"             // Ensure this keyspace exists in Cassandra
	cluster.Consistency = gocql.Quorum

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Failed to connect to Cassandra:", err)
	}

	fmt.Println("Connected to Cassandra successfully!")
}
