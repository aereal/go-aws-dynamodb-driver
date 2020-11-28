package ddbdriver

import (
	"context"
	"database/sql/driver"
)

const (
	// DriverName is name of the driver this package provides
	DriverName = "awsdynamodb"
)

type ddbDriver struct{}

var _ interface {
	driver.Driver
} = &ddbDriver{}

func (d *ddbDriver) Open(dsn string) (driver.Conn, error) {
	connector, err := d.OpenConnector(dsn)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return connector.Connect(ctx)
}

func (d *ddbDriver) OpenConnector(dsn string) (driver.Connector, error) {
	return &ddbConnector{}, nil
}
