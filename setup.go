package pcql

import (
	"log"

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

func Execute(stmt string, names ...interface{}) error {
	return session.Query(stmt, names...).Exec()
}
