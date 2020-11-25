package ddbdriver

import (
	"context"
	"database/sql/driver"
)

type ddbConnector struct{}

var _ driver.Connector = &ddbConnector{}

func (c *ddbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	panic("not implemented") // TODO: Implement
}

func (c *ddbConnector) Driver() driver.Driver {
	return &ddbDriver{}
}
