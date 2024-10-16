package main

import "fmt"


type Vehicle interface {
	Mileage(liter int) int
}

type Car struct {
	speed int
}

type MotorBike struct {
	speed int
}

type Bajaj struct {
	speed int
}

// Menghitung jarak tempuh berdasarkan liter bensin

func (c Car) Mileage(oil int) int{
	return oil * 1
}

func (m MotorBike) Mileage(oil int) int{
	return oil * 3
}

func (b Bajaj) Mileage(oil int) int{
	return oil * 4
}
 

func eff(liter int, vehicle ...Vehicle) string {
	
	var max int
	var efficiency string

	for _, vehicle := range vehicle {
		distence := vehicle.Mileage(liter)

		switch vehicle.(type) {

		case Car:
			if distence > max {
				max = distence
				efficiency = "Mobil"
			}
		case MotorBike:
			if distence > max {
				max = distence
				efficiency = "Motor"
			}
		case Bajaj:
			if distence > max {
				max = distence
				efficiency = "Bajaj"
			}
		}
	}

	return efficiency + " adalah Kendaraan paling efisiensi"
}



func main()  {
	
	motorBike := MotorBike{speed: 60}
	car := Car{speed: 60}
	bajaj := Bajaj{speed: 60}

	oil := 10

	result := eff(oil, motorBike, car,bajaj)

	fmt.Println(result)
}