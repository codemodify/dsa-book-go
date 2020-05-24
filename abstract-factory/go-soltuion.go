package main

import "fmt"

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Abstract Factory based on that.
//

type Car interface {
	BuildWheels() []IWheel
	BuildShape() IShape
}

type IWheel interface {
	GetDiameter() int
}

type IShape interface {
	GetShape() string
}

// SPORTS
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
type sportsCar struct{}

func (l sportsCar) BuildWheels() []IWheel {
	return []IWheel{
		// ...
	}
}
func (l sportsCar) BuildShape() IShape {
	// return ...
	return nil
}

// OFFROAD
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
type offroadCar struct{}

func (l offroadCar) BuildWheels() []IWheel {
	return []IWheel{
		offroadCarWheel{
			diameter: 120,
		},
		offroadCarWheel{
			diameter: 120,
		},
		offroadCarWheel{
			diameter: 120,
		},
		offroadCarWheel{
			diameter: 120,
		},
	}
}
func (l offroadCar) BuildShape() IShape {
	return offroadCarShape{
		shape: "boxy",
	}
}

type offroadCarWheel struct {
	diameter int
}

func (w offroadCarWheel) GetDiameter() int {
	return w.diameter
}

type offroadCarShape struct {
	shape string
}

func (s offroadCarShape) GetShape() string {
	return s.shape
}

// FACTORY
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
type carType int

const (
	carTypeSports carType = iota
	carTypeOffroad
)

func createCar(lt carType) Car {
	switch lt {
	case carTypeSports:
		return &sportsCar{}
	case carTypeOffroad:
		return &offroadCar{}
	}

	return nil
}

func main() {
	car := createCar(carTypeOffroad)
	car.BuildWheels()

	fmt.Println(car.BuildWheels()[0].GetDiameter())
	fmt.Println(car.BuildShape().GetShape())
}
