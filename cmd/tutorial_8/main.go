package main

import (
	"fmt"
)

type owner struct {
	name        string
	age         int
	nationality string
}

func (o *owner) setName(name string) {
	o.name = name
}
func (o *owner) setAge(a int) {
	o.age = a
}
func (o owner) setNationality(n string) { //--> value receiver and will not modify the pointer struct in main
	o.nationality = n
}

func main() {
	//var pointer *int = new(int)
	//fmt.Println("pointer address", pointer)
	//fmt.Println("value pointer points to", *pointer)
	//var i = 10
	//fmt.Println("Address of i: ", &i)
	//pointer = &i
	//fmt.Println("pointer address", pointer)
	//fmt.Println("value pointer points to", *pointer)
	//i = 20
	//fmt.Println("value pointer points to", *pointer)
	//*pointer = 10
	//fmt.Println("Value of i: ", i)
	//fmt.Println("value pointer points to", *pointer)
	//var name = "Saif"
	//fmt.Println("Address of name: ", &name)
	//pointer = &name --> cannot change type

	var arr = [5]int32{2, 4, 6, 8, 10}
	fmt.Println("Array:", arr)
	fmt.Println("Square of the array:", squareArr(&arr))

	val1, val2 := 50, 100
	fmt.Printf("val1: %v, val2: %v\n", val1, val2)
	swapInt(&val1, &val2)
	fmt.Printf("val1: %v, val2: %v\n", val1, val2)

	var address = &val1
	fmt.Printf("address of val1 %v and value: %v\n", address, *address)
	swapInt(address, &val2)
	fmt.Printf("val1: %v, val2: %v\n", val1, val2)

	var ownerPointer = &owner{}
	ownerPointer.setName("Saif")
	ownerPointer.setAge(21)
	ownerPointer.setNationality("Pakistan")
	fmt.Printf("owner name: %v of age: %v and nationality: %v\n", ownerPointer.name, ownerPointer.age,
		ownerPointer.nationality)

}

func squareArr(arr *[5]int32) [5]int32 {
	for i := range *arr {
		(*arr)[i] = (*arr)[i] * (*arr)[i]
	}
	return *arr
}

func swapInt(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
