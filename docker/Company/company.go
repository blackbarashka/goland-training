package Company

import (
	"context"
	"docker/database"
	"docker/employee"
	"sync"

	"github.com/jackc/pgx/v5"
)

type Company struct {
	mtx  sync.Mutex
	ctx  context.Context
	conn *pgx.Conn
}

func NewCompany(ctx context.Context, conn *pgx.Conn) *Company {
	return &Company{
		ctx:  ctx,
		conn: conn,
	}
}

func (company *Company) Add(fullName, position string) error {
	company.mtx.Lock()
	defer company.mtx.Unlock()
	worker := employee.NewEmployee(0, fullName, position)
	err := database.InsertDB(company.ctx, company.conn, worker)
	return err
}
func (company *Company) GetWorkers() ([]employee.Employee, error) {

	workers, err := database.GetDB(company.ctx, company.conn)
	if err != nil {
		return nil, err
	}
	return workers, nil
}
func (company *Company) DeleteWorker(id int64) error {
	company.mtx.Lock()
	defer company.mtx.Unlock()

	err := database.DeleteDB(company.ctx, company.conn, id)
	return err
}
