package pcql

import "github.com/scylladb/gocqlx/v3/table"

func CreateTable(name string, partKeys, sortKeys, columns []string) *table.Table {
	return table.New(table.Metadata{
		Name:    name,
		PartKey: partKeys,
		SortKey: sortKeys,
		Columns: columns,
	})
}
