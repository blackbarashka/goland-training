package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"task_list/core/db"
	"task_list/internal/features/tasks/delivery"
	"task_list/internal/features/tasks/repository"
	"task_list/internal/features/tasks/service"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx_signal, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	ctx := context.Background()
	connString := ""
	conn, err := db.NewConnection(ctx, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(conn, ctx)

	repo := repository.NewPostgresRepository(conn)
	svc := service.NewTaskService(repo)
	handler := delivery.NewTaskHandler(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /tasks", handler.CreateTask)
	mux.HandleFunc("GET /tasks", handler.ListTasks)
	mux.HandleFunc("GET /tasks/{id}", handler.GetTask)
	mux.HandleFunc("PUT /tasks/{id}", handler.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", handler.DeleteTask)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v\n", err)
		}
	}()

	<-ctx_signal.Done()
	log.Println("Получен сигнал. Останавливаем сервер...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Ошибка при остановке сервера: %v", err)
	}

	log.Println("Сервер успешно остановлен")

}
