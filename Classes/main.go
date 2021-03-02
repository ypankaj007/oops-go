package main

import (
	"oops-go/classes/entity"
)

func main() {
	emp := entity.New(1, "Jack", 4509)
	emp.Display()
}
