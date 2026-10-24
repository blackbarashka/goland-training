package database

import (
	"context"
	"docker/employee"
	"errors"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func Connection(ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to database")
	return conn, nil
}

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
     CREATE TABLE IF NOT EXISTS employees (
        id SERIAL PRIMARY KEY,
        full_name VARCHAR(255) NOT NULL,
    	position VARCHAR(255) NOT NULL
     );`
	_, err := conn.Exec(ctx, sqlQuery)
	if err != nil {
		return errors.New("database creation failed: " + err.Error())
	}
	return nil
}

func InsertDB(ctx context.Context, conn *pgx.Conn, worker *employee.Employee) error {
	sqlQuery := `
    INSERT INTO employees( full_name, position) 
    VALUES ($1, $2);`
	_, err := conn.Exec(ctx, sqlQuery, worker.FullName, worker.Position)
	return err
}

func GetDB(ctx context.Context, conn *pgx.Conn) ([]employee.Employee, error) {
	sqlquery := `
	SELECT id, full_name, position
	FROM employees`
	rows, err := conn.Query(ctx, sqlquery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var workers []employee.Employee
	for rows.Next() {
		var worker employee.Employee
		err := rows.Scan(&worker.Id, &worker.FullName, &worker.Position)
		if err != nil {
			return nil, err
		}
		workers = append(workers, worker)
	}
	return workers, nil
}
func DeleteDB(ctx context.Context, conn *pgx.Conn, id int64) error {
	sqlQuery := `DELETE FROM employees WHERE id = $1;`
	_, err := conn.Exec(ctx, sqlQuery, id)
	return err
}
