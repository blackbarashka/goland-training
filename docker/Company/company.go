package Company

import (
	"docker/employee"
	"sync"
)

type Company struct {
	mtx     sync.Mutex
	workers []employee.Employee
	id      int64
}

func NewCompany() *Company {
	return &Company{
		workers: make([]employee.Employee, 0),
	}
}

func (company *Company) Add(fullName, position string) {
	company.mtx.Lock()
	defer company.mtx.Unlock()
	worker := employee.NewEmployee(company.id, fullName, position)
	company.id++
	company.workers = append(company.workers, *worker)
}
func (company *Company) GetWorkers() []employee.Employee {
	return company.workers
}
func (company *Company) DeleteWorker(id int64) {
	company.mtx.Lock()
	defer company.mtx.Unlock()

	for i, worker := range company.workers {
		if worker.Id == id {
			company.workers = append(company.workers[:i], company.workers[i+1:]...)
			break
		}
	}
}
