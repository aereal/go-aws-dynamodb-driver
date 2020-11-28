package ddbdriver

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var (
	defaultEndpointScheme = "https"
)

type Config struct {
	Endpoint string
}

func ParseDSN(dsn string) (*Config, error) {
	parsed, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	scheme, err := parseScheme(parsed.Scheme)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	if parsed.Host != "" || scheme != defaultEndpointScheme {
		cfg.Endpoint = fmt.Sprintf("%s://%s", scheme, parsed.Host)
	}
	return cfg, nil
}

// "awsdynamodb" => "https", nil
// "awsdynamodb+http" => "http", nil
func parseScheme(scheme string) (string, error) {
	if !strings.Contains(scheme, DriverName) {
		return "", errors.New("invalid DSN scheme")
	}
	if scheme == DriverName {
		return defaultEndpointScheme, nil
	}
	return strings.Replace(scheme, DriverName+"+", "", 1), nil
}
