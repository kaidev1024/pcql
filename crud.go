package pcql

import (
	"fmt"

	"github.com/scylladb/gocqlx/qb"
)

func Insert(stmt string, names []string, row any) error {
	err := sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func Get(stmt string, names []string, input, ret any) error {
	return sessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func Delete(stmt string, names []string, row any) error {
	return sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func Select(stmt string, names []string, input qb.M, rows any) error {
	err := sessionx.Query(stmt, names).BindMap(input).SelectRelease(rows)
	if err != nil {
		return err
	}
	return nil
}

func Update(stmt string, names []string, row any) error {
	err := sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func Count(stmt string) (int, error) {
	var count int
	isSuccess := sessionx.Query(stmt, nil).Iter().Scan(&count)
	if isSuccess {
		return count, nil
	}
	return -1, fmt.Errorf("count failed")
}
