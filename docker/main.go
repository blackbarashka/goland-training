package main

import (
	"context"
	"docker/Company"
	"docker/database"
	"docker/http_server"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	if os.Getenv("DATABASE_URL") == "" {
		err := os.Setenv("DATABASE_URL", "postgres://postgres:admin@localhost:5432/postgres")
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	ctx := context.Background()
	conn, err := database.Connection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, ctx)
	if err := database.CreateTable(ctx, conn); err != nil {
		log.Fatal(err)
		return
	}

	company := Company.NewCompany(ctx, conn)
	handler := http_server.NewHandler(company)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /employees", handler.AddHTTP)
	mux.HandleFunc("GET /employees", handler.GetWorkersHTTP)
	mux.HandleFunc("DELETE /employees/{id}", handler.RemoveHTTP)
	if err := http_server.StartHTTPServer(mux); err != nil {
		log.Fatal(err)
	}
}
