package entity

import "fmt"

// Employee as class
type Employee struct {
	ID     int
	Name   string
	Salary float32
}

// New as construtor
func New(id int, name string, salary float32) *Employee {
	return &Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}

// Display as method of Employee class
func (e *Employee) Display() {
	fmt.Println(fmt.Sprintf("ID\tName\tSalary\n%d\t%s\t%f", e.ID, e.Name, e.Salary))
}
