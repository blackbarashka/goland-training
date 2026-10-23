package sql

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connection(ctx context.Context) (*pgx.Conn, error) {
	val, ok := os.LookupEnv("CONN_STRING")
	if !ok {
		return nil, errors.New("CONN_STRING environment variable not set")
	}
	conn, err := pgx.Connect(ctx, val)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
