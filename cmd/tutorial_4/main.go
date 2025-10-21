package main

import "fmt"

func main() {
	var intArrA [11]int32 //type array [11]int32
	intArrA[0] = 0
	intArrA[1] = 1
	intArrA[2] = 2
	intArrA[3] = 3
	intArrA[4] = 4
	fmt.Println(&intArrA[0], &intArrA[1], &intArrA[2], &intArrA[3], &intArrA[4], &intArrA[5])
	fmt.Println(intArrA[0])
	fmt.Println(intArrA[1:2])
	fmt.Println(intArrA[2:2])
	fmt.Println(intArrA[2:5])
	fmt.Println(intArrA[:2])
	fmt.Println(intArrA[2:])
	fmt.Println(intArrA[4:5])

	var intArrB = intArrA[:5] //type slice []int32
	intArrA[1] = 11
	fmt.Println("Array A:", intArrA)
	fmt.Println("Array B:", intArrB)
	fmt.Printf("The len of arrA is %v and capacity is %v \n", len(intArrA), cap(intArrA))
	fmt.Printf("The len of arrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of array a is %v and array b is %v \n", &intArrA[0], &intArrB[0])
	intArrB = append(intArrB, 5, 6, 7, 8, 9, 10)
	intArrA[2] = 22
	fmt.Println("Array A:", intArrA)
	fmt.Println("Array B:", intArrB)
	fmt.Printf("The len of arrA is %v and capacity is %v \n", len(intArrA), cap(intArrA))
	fmt.Printf("The len of arrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of array a is %v and array b is %v \n", &intArrA[0], &intArrB[0])
	intArrB = append(intArrB, 11)
	intArrA[3] = 33
	fmt.Println("Array A:", intArrA)
	fmt.Println("Array B:", intArrB)
	fmt.Printf("The len of arrA is %v and capacity is %v \n", len(intArrA), cap(intArrA))
	fmt.Printf("The len of arrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of array a is %v and array b is %v \n", &intArrA[0], &intArrB[0])

	intSlice := []int32{55, 66, 77} //type slice []int32
	fmt.Println("Slice:", intSlice)
	fmt.Printf("The len of intSlice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	fmt.Printf("The memory of slice is %v and array b is %v \n", &intSlice[0], &intArrB[0])
	intSlice = append(intArrB, 88)
	fmt.Println("Slice:", intSlice)
	fmt.Println("Array B", intArrB)
	fmt.Printf("The len of intSlice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	fmt.Printf("The len of intArrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of slice is %v and array b is %v \n", &intSlice[0], &intArrB[0])
	intSlice = append(intSlice, 99)
	fmt.Println("Slice:", intSlice)
	fmt.Println("Array B:", intArrB)
	fmt.Printf("The len of intSlice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	fmt.Printf("The len of intArrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of slice is %v and array b is %v \n", &intSlice[0], &intArrB[0])
	fmt.Println("Changing Array B..")
	intArrB = append(intArrB, 1000)
	fmt.Println("Array B:", intArrB)
	fmt.Println("Slice:", intSlice)
	fmt.Printf("The len of intSlice is %v and capacity is %v \n", len(intSlice), cap(intSlice))
	fmt.Printf("The len of intArrB is %v and capacity is %v \n", len(intArrB), cap(intArrB))
	fmt.Printf("The memory of slice is %v and array b is %v \n", &intSlice[0], &intArrB[0])

	intArrC := [...]int32{1, 2, 3} //type array [3]int32
	fmt.Println("Array C:", intArrC)
	fmt.Println(len(intArrC), cap(intArrC))
	//intArrC = append(intArrC,4) --> cannot do this because it's not a slice

	var myMap map[string]uint8 = map[string]uint8{"Saif": 21, "Dev": 34} //type map[string]uint8
	fmt.Println(myMap)
	myMap["Adam"] = 3
	myMap["Saif"] = 22
	fmt.Println(myMap)
	var age, present = myMap["Saif"]
	if present {
		fmt.Println("Age of Saif:", age)
	}
	for k, v := range myMap {
		fmt.Println("Name:", k, "Age:", v)
	}
	for i, v := range intArrC {
		fmt.Println("Index:", i, "Value:", v)
	}
	for _, v := range intArrC {
		fmt.Println("Value:", v)
	}
	for _, v := range myMap {
		fmt.Println("Age:", v)
	}
	for i := 0; i < len(intArrC); i++ {
		fmt.Println("Index:", i, "Value:", intArrC[i])
	}

	var x int = 0
	for x < 10 { // -->while loop in go :]]
		fmt.Println(x)
		x++
	}
	for {
		if x == 20 {
			break
		}
		fmt.Println(x)
		x++
	}

}
