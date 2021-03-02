package car

import "fmt"

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
