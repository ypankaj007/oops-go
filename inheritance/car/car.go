package car

import "fmt"

// Car struct
type Car struct {
	Model  string
	Wheets int
}

func (c *Car) Start() {
	// check if car is ready to start
	// like: hasFeul etc
	fmt.Println(fmt.Sprintf("%s %s", c.Model, "started!"))
}
