# Classes

#### Structs as Classes
##### Golang (Go) does not provide classes but it has structs (Structs are user-defined data types/structures). A method can be added on structs. This provides the behaviour of bundling the data and methods that operate on the data together akin to a class.
#
Lets look at the below example: We have an employee as class with id, name and salary as attributes. We can define methods on types. A method is a function with a special receiver argument. In this example, the Display method has a receiver of type Employee named e. New func acting as construtor here which initialze Employee and return the reference of it.
#
```go

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

```
##
##### How to use employee class
#
###### In below, We imported entity package and called New functin with ID as 1, Name as Jack and Salary as 4509 to initialize Employee instance. Then we called Display method on that to print the employee records
##
```go
package main

import (
	"oops-go/Classes/entity"
)

func main() {
	emp := entity.New(1, "Jack", 4509)
	emp.Display()
}
```
