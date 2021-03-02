# Inheritance

##### In Go, inheritance can be achived by embebbbing structs and uning interfaces.
#
##### To inherit a struct,  child struct can use parent struct as embedded field.
#
##### To inherit a interface, struct must implements all it's method. (If struct missed any method to implement of interface then it will consider as it's not inherired) 
#
Lets look at the below example:
Car struct has model and wheels as field member, and start method.
There are two interfaces Flyable and Floatable. 
ArmoredCar embedded Car as field member and implements all methods of Flyable and Floatable interface. So ArmoredCar inherits Car, Flyable and Floatable
#
```go
package main

import "fmt"

// Flyable interface
type Flyable interface {
	Fly()
}

// Floatable interface
type Floatable interface {
	FloatOnWater()
}

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

// ArmoredCar features
type ArmoredCar struct {
	BulletProofWindows int
	*Car               // embebbing Car in ArmoredCar
}

func (a *ArmoredCar) RemoteStart() {
	fmt.Println("Starting ArmoredCar by remote")
	a.Start() // calling car's start method
}

// implementing Floatable interface floatOnWater()
func (a *ArmoredCar) FloatOnWater() {
	fmt.Println("I can float!")
}

// implementing Flyable interface fly()
func (a *ArmoredCar) Fly() {
	fmt.Println("I can fly!")
}

func main() {

	armoredCar := &ArmoredCar{
		BulletProofWindows: 4,
		Car: &Car{
			Model:  "SUV DXI",
			Wheets: 4,
		},
	}

	armoredCar.RemoteStart()
	armoredCar.Fly()
	armoredCar.FloatOnWater()

}

```
