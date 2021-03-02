package main

import (
	"oops-go/inheritance/car"
)

func main() {

	armoredCar := &car.ArmoredCar{
		BulletProofWindows: 4,
		Car: &car.Car{
			Model:  "SUV DXI",
			Wheets: 4,
		},
	}

	armoredCar.RemoteStart()
	armoredCar.Fly()
	armoredCar.FloatOnWater()

}
