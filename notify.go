package pcql

import (
	"fmt"
	"time"

	astra "github.com/datastax/gocql-astra"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v3"
)

var notifySessionx gocqlx.Session
var notifySession *gocql.Session

func SetupCassandraNotify(id, token, keyspace string) error {
	cluster, err := astra.NewClusterFromURL("https://api.astra.datastax.com", id, token, 10*time.Second)
	if err != nil {
		return fmt.Errorf("unable to create cassandra(astra) cluster: %v", err)
	}

	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.ConnectTimeout = 5 * time.Second

	notifySession, err = cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("could not connect to astra db: %v", err)
	}
	notifySessionx, err = gocqlx.WrapSession(notifySession, err)
	if err != nil {
		return fmt.Errorf("could not wrap session: %v", err)
	}
	return nil
}

func InsertNotify(stmt string, names []string, row any) error {
	err := notifySessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func GetNotify(stmt string, names []string, input, ret any) error {
	return notifySessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func SelectPaginatedNotify(stmt string, names []string, input M, rows any, pageState []byte) (*PaginationResult, error) {
	pageSize := 20
	query := notifySessionx.Query(stmt, names).BindMap(input).PageSize(pageSize)

	if len(pageState) > 0 {
		query = query.PageState(pageState)
	}

	iter := query.Iter()
	if err := iter.Select(rows); err != nil {
		iter.Close()
		return nil, err
	}

	nextPageState := iter.PageState()
	if err := iter.Close(); err != nil {
		return nil, err
	}

	return &PaginationResult{
		PageState: nextPageState,
		HasMore:   len(nextPageState) > 0,
	}, nil
}

func UpdateNotify(stmt string, names []string, row any) error {
	err := notifySessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}
