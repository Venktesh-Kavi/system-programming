package main

import "fmt"

type Vehicle interface {
	Drive()
}

type Engine struct {
	capacity  string
	torque    int
	cylinders int
}

func (e Engine) Stringify() string {
	return fmt.Sprintf("torque: %d, capacity: %s, cylinder: %d", e.torque, e.capacity, e.cylinders)
}

type Truck struct {
	Vehicle
	engine Engine
	brand  string
}

func (t Truck) Drive() {
	fmt.Println("driving now...")
}

func main() {
	e := Engine{capacity: "1.5L", torque: 220, cylinders: 3}
	t := Truck{engine: e, brand: "eicher"}
	t.Drive()
	fmt.Println("Injected Engine: ", t.engine.Stringify())
}
