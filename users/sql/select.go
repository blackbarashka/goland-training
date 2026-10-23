package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func Select(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT id, phone_number, full_name 
	FROM users
	`
	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var fullName, phoneNumber string
		if err := rows.Scan(&id, &phoneNumber, &fullName); err != nil {
			return err
		}
		fmt.Printf("id - %d\nfull_name - %s\nphone_number - %s\n", id, fullName, phoneNumber)
	}
	return nil

}
