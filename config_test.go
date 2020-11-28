package ddbdriver

import (
	"reflect"
	"testing"
)

func TestParseDSN(t *testing.T) {
	tests := []struct {
		name    string
		dsn     string
		want    *Config
		wantErr bool
	}{
		{"OK", "awsdynamodb://", &Config{}, false},
		{"custom scheme", "awsdynamodb+http://", &Config{Endpoint: "http://"}, false},
		{"custom endpoint", "awsdynamodb+http://localhost:8000", &Config{Endpoint: "http://localhost:8000"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDSN(tt.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("actual=%q expected=%v", got, tt.want)
			}
		})
	}
}
