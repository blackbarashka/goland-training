package repository

import (
	"context"
	"sync"
	"task_list/internal/features/tasks"
	"time"

	"github.com/jackc/pgx/v5"
)

type TaskRepository interface {
	Create(ctx context.Context, task *tasks.Task) error
	Get(ctx context.Context, id string) (*tasks.Task, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, id string, task *tasks.Task) error
	List(ctx context.Context) ([]*tasks.Task, error)
}

type postgresRepository struct {
	conn *pgx.Conn
	mu   sync.Mutex
}

func NewPostgresRepository(conn *pgx.Conn) TaskRepository {
	return &postgresRepository{conn: conn}
}

func (p *postgresRepository) Create(ctx context.Context, task *tasks.Task) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	sqlQuery := `
    INSERT INTO tasks(title, description, completed, created_at) 
    VALUES ($1, $2, $3, $4)
    RETURNING id`

	err := p.conn.QueryRow(ctx, sqlQuery, task.Title, task.Description, task.Completed, time.Now()).Scan(&task.Id)
	return err
}

func (p *postgresRepository) Update(ctx context.Context, id string, task *tasks.Task) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	sqlQuery := `
    UPDATE tasks 
    SET title=$2, description=$3, completed=$4
    WHERE id=$1
    `
	_, err := p.conn.Exec(ctx, sqlQuery, id, task.Title, task.Description, task.Completed)
	return err
}
func (p *postgresRepository) Delete(ctx context.Context, id string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	sqlQuery := `DELETE FROM tasks WHERE id=$1`
	_, err := p.conn.Exec(ctx, sqlQuery, id)
	return err
}
func (p *postgresRepository) Get(ctx context.Context, id string) (*tasks.Task, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	sqlQuery := `
    SELECT id, title, description, completed, created_at
    FROM tasks
    WHERE id=$1
    `
	var task tasks.Task
	err := p.conn.QueryRow(ctx, sqlQuery, id).Scan(&task.Id, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (p *postgresRepository) List(ctx context.Context) ([]*tasks.Task, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	sqlQuery := `
    SELECT id, title, description, completed, created_at
    FROM tasks
    `

	rows, err := p.conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var taskList []*tasks.Task
	for rows.Next() {
		var task tasks.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Completed, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return taskList, nil
}
