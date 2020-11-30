package ddbdriver

import (
	"context"
	"database/sql/driver"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ddbConnector struct {
	cfg *Config
}

var _ driver.Connector = &ddbConnector{}

func (c *ddbConnector) Connect(ctx context.Context) (driver.Conn, error) {
	awsCfg := aws.Config{}
	if c.cfg.Endpoint != "" {
		awsCfg.Endpoint = &c.cfg.Endpoint
	}
	ses, err := session.NewSessionWithOptions(session.Options{Config: awsCfg})
	if err != nil {
		return nil, err
	}
	client := dynamodb.New(ses)
	return &ddbConn{client: client}, nil
}

func (c *ddbConnector) Driver() driver.Driver {
	return &ddbDriver{}
}
