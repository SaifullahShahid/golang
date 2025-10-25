package main

import (
	"fmt"
	//"strconv"
)

type car struct {
	name      string
	year      int
	color     string
	ownerInfo owner
}
type owner struct {
	name string
	age  int
}

func main() {
	var car1 car = car{"Toyota Platz", 2007, "silver", owner{"Saif", 21}}
	fmt.Println(car1.name, car1.year, car1.color, car1.ownerInfo.name)
	car1.year = 2008
	car1.ownerInfo.name = "Alex"
	fmt.Println(car1.year, car1.ownerInfo.name, car1.ownerInfo.age)

	var car2 car
	car2.name = "Honda Civic"
	car2.ownerInfo.name = car1.ownerInfo.name

	fmt.Println(car1.carDetails())
	fmt.Println(car2.carDetails())

	//var car3 = struct {
	//	carName string
	//	modelNo string
	//}{"Carrera GT", "L3ST3R"}
	//fmt.Println(car3)
	//var car4 = struct {
	//	carName string
	//	modelNo string
	//}{}
	//car4 = car3
	//fmt.Println(car4)

	var rect = Rectangle{
		Width:  10,
		Length: 5,
	}
	circle := Circle{10.75}
	//var tri = Triangle{
	//	height: 20,
	//	base:   30.5,
	//}
	var areaR, perimeterR = describeShape(rect)
	fmt.Println("Rectangle Area:", areaR, "Perimeter:", perimeterR)
	var areaC, perimeterC = describeShape(circle)
	fmt.Println("Triangle Area:", areaC, "Perimeter:", perimeterC)

	anyValue := interface{}("Saif")
	describeValue(anyValue)
	describeValue(rect)
	describeValue(circle)

	val, ok := anyValue.(string)
	if ok {
		fmt.Println("Value:", val)
	} else {
		fmt.Println("The value is not a string")
	}

	sqrt, err := squareRoot(-64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(sqrt)
	}

}

func (c car) carDetails() string {
	return fmt.Sprintf("Car name: %s, Year: %d, Color: %s, Car ownerInfo: %s:%d",
		c.name, c.year, c.color, c.ownerInfo.name, c.ownerInfo.age)
	//"Car name: " + c.name + "\nCar year: " + strconv.Itoa(c.year) + "\nCar color: " +
	//	c.color + "\nCar ownerInfo: " + c.ownerInfo.name
}
