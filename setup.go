package pcql

import (
	"log"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var sessionx gocqlx.Session
var session *gocql.Session

func Setup(address, keyspace string) {
	cluster := gocql.NewCluster(address)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.LocalOne
	cluster.ProtoVersion = 4
	// for authentication:
	//   cluster.Authenticator = gocql.PasswordAuthenticator{
	// 		Username: "user",
	// 		Password: "password"
	//  }
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	sessionx, err = gocqlx.WrapSession(session, err)
	if err != nil {
		log.Fatal(err)
	}
	// defer keyspaceWorenaSession.Close()
}

func SetupCassandra(id, token, keyspace string) {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		log.Fatalf("unable to create astra cluster config: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatalf("could not connect to astra db: %v", err)
	}
	sessionx, err = gocqlx.WrapSession(session, err)
	if err != nil {
		log.Fatal(err)
	}
}

func Execute(stmt string, names ...interface{}) error {
	return session.Query(stmt, names...).Exec()
}
