package employee

type Employee struct {
	Id       int64
	FullName string
	Position string
}

func NewEmployee(id int64, fullName string, position string) *Employee {
	return &Employee{
		Id:       id,
		FullName: fullName,
		Position: position,
	}
}
