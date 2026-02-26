package pcql

import (
	"fmt"
)

type PaginationResult struct {
	PageState []byte
	HasMore   bool
}

func Insert(stmt string, names []string, row any) error {
	return sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func Get(stmt string, names []string, input, ret any) error {
	return sessionx.Query(stmt, names).BindStruct(input).GetRelease(ret)
}

func Delete(stmt string, names []string, row any) error {
	return sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func Select(stmt string, names []string, input M, rows any) error {
	return sessionx.Query(stmt, names).BindMap(input).SelectRelease(rows)
}

// SelectPaginated performs a paginated select query and returns the results along with pagination information.
// pageSize defaults to 20, and the caller can provide a pageState to fetch subsequent pages of results.
func SelectPaginated(stmt string, names []string, input M, rows any, pageState []byte) (*PaginationResult, error) {
	pageSize := 20
	query := sessionx.Query(stmt, names).BindMap(input).PageSize(pageSize)

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

func Update(stmt string, names []string, row any) error {
	return sessionx.Query(stmt, names).BindStruct(row).ExecRelease()
}

func Count(stmt string) (int, error) {
	var count int
	isSuccess := sessionx.Query(stmt, nil).Iter().Scan(&count)
	if isSuccess {
		return count, nil
	}
	return -1, fmt.Errorf("count failed")
}
