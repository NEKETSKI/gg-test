package postgres

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
)

type mocks struct {
	execMock func(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

func (m mocks) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return m.execMock(ctx, sql, arguments...)
}

func (m mocks) Close(_ context.Context) error {
	return nil
}

func Test_psql_StoreStatusCode(t *testing.T) {
	type fields struct {
		db postgresApiClient
	}
	type args struct {
		requestUrl string
		code       int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			"No error",
			fields{db: mocks{execMock: func(ctx context.Context, _ string, arguments ...any) (pgconn.CommandTag, error) {
				assert.ElementsMatch(t, []any{"http://test.com", 200}, arguments[:2])
				return pgconn.CommandTag{}, nil
			}}},
			args{
				requestUrl: "http://test.com",
				code:       200,
			},
			assert.NoError,
		},
		{
			"Error case",
			fields{db: mocks{execMock: func(ctx context.Context, _ string, arguments ...any) (pgconn.CommandTag, error) {
				assert.ElementsMatch(t, []any{"http://example.com", 404}, arguments[:2])
				return pgconn.CommandTag{}, errors.New("test error")
			}}},
			args{
				requestUrl: "http://example.com",
				code:       404,
			},
			assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &psql{
				db: tt.fields.db,
			}
			tt.wantErr(t, p.StoreStatusCode(tt.args.requestUrl, tt.args.code),
				fmt.Sprintf("StoreStatusCode(%v, %v)", tt.args.requestUrl, tt.args.code))
		})
	}
}
