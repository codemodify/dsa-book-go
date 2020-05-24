package main

import "fmt"

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Builder based on that.
//
// The example below is kept minimal to avoid code noise so you can capture
// the idea.
//

type Car struct {
	Wheels string
	Shape  string
	Color  string
}

type CarBuilder interface {
	SetWheels()
	SetShape()
	SetColor()
}

// SPORTS
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
type sportsCar struct {
	car *Car
}

func (l *sportsCar) SetWheels() {
	l.car.Wheels = "michelin, black 20 inch rims, summer tire"
}
func (l *sportsCar) SetShape() {
	l.car.Shape = "slim, F1 type"
}
func (l *sportsCar) SetColor() {
	l.car.Color = "white"
}

// OFFROAD
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
type offroadCar struct {
	car *Car
}

func (l *offroadCar) SetWheels() {
	l.car.Wheels = "open country, all weather"
}
func (l *offroadCar) SetShape() {
	l.car.Shape = "slim, F1 type"
}
func (l *offroadCar) SetColor() {
	l.car.Color = "black"
}

// usage
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
func main() {
	car := Car{}

	var builder CarBuilder

	builder = &sportsCar{car: &car}
	builder.SetWheels()
	builder.SetShape()
	builder.SetColor()
	fmt.Println(car)

	builder = &offroadCar{car: &car}
	builder.SetWheels()
	builder.SetShape()
	builder.SetColor()
	fmt.Println(car)
}
