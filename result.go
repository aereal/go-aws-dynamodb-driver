package ddbdriver

import "database/sql/driver"

type ddbResult struct{}

var _ driver.Result = &ddbResult{}

func (r *ddbResult) LastInsertId() (int64, error) {
	panic("not implemented") // TODO: Implement
}

func (r *ddbResult) RowsAffected() (int64, error) {
	panic("not implemented") // TODO: Implement
}
