// Package postgres provides primitives for connecting to and interacting with a Postgres database.
package postgres

import (
	"context"
	"time"

	"github.com/NEKETSKY/gg-test/internal/repository"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type postgresApiClient interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close(ctx context.Context) error
}

type psql struct {
	db postgresApiClient
}

// Init inits a new psql instance and returns it as a repository.Repository interface.
// Returns an error if the database connection fails.
func Init(connUrl string) (repository.Repository, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	conn, err := pgx.Connect(ctx, connUrl)
	if err != nil {
		return nil, err
	}
	return &psql{db: conn}, nil
}

// Close closes the database connection.
func (p *psql) Close() error {
	ctx, cancelF := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelF()
	return p.db.Close(ctx)
}
