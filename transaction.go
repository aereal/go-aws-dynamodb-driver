package ddbdriver

import "database/sql/driver"

type ddbTx struct{}

var _ interface {
	driver.Tx
} = &ddbTx{}

func (t *ddbTx) Commit() error {
	panic("not implemented") // TODO: Implement
}

func (t *ddbTx) Rollback() error {
	panic("not implemented") // TODO: Implement
}
