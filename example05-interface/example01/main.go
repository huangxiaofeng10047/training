package main

import (
	"fmt"
	"github.com/go-training/training/example05-interface/example01/benchi"

	"github.com/go-training/training/example05-interface/example01/lexus"
	"github.com/go-training/training/example05-interface/example01/toyota"
)

type car interface {
	Name() string
	Price() float64
	Discount() float64
	Color() string
}

func detail(c car) {
	fmt.Println("==================")
	fmt.Println("Name:", c.Name())
	fmt.Println("Price:", c.Price())
	fmt.Println("Discount:", c.Discount())
	fmt.Println("==================")
}

func main() {
	car1 := toyota.NewCar("car1", 3000, 0.8, "white")
	car2 := toyota.NewCar("car2", 4000, 0.9, "white")
	car3 := lexus.NewCar("car3", 5000, 0.7, "blue")
	car4 := lexus.NewCar("car4", 6000, 0.7, "red")

	detail(car1)
	detail(car2)
	detail(car3)
	detail(car4)
	//add benchi

	car5 := benchi.NewCar("car5", 5000, 0.7, "white")
	detail(car5)
}
