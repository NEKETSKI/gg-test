package postgres

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		connUrl string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"Successfully connected",
			"postgres://postgres:postgres@localhost:5432/postgres",
			assert.NoError,
		},
		{
			"Connection error",
			"postgres://fake_user:postgres@localhost:5432/postgres",
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Init(tt.connUrl)
			if !tt.wantErr(t, err, fmt.Sprintf("Init(%v)", tt.connUrl)) {
				return
			}
		})
	}
}
