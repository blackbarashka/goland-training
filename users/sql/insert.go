package sql

import (
	"context"
	"users/models"

	"github.com/jackc/pgx/v5"
)

func Insert(ctx context.Context, conn *pgx.Conn, user *models.User) error {
	sqlQuery := `
	INSERT INTO users(full_name, phone_number)
	VALUES ($1, $2)
`
	_, err := conn.Exec(ctx, sqlQuery, user.FullName, user.PhoneNumber)
	return err
}
