package main

import "fmt"
import "unicode/utf8"

func main() {
	var numInt int
	fmt.Println(numInt)

	num2Int := 21
	fmt.Println(num2Int)

	var numFloat = 3.12
	var num3Int = int(numFloat)
	fmt.Println(num3Int)

	var num2Float float64 = float64(num2Int)
	fmt.Println("num2Float value:", num2Float)
	fmt.Printf("%.1f\n", num2Float)

	fmt.Println("num1Float + num2Float value:", numFloat+num2Float)

	var Name string = "Saif" //public variable
	fmt.Println(Name)

	var Age int = 21
	fmt.Println("name:", Name, "age:", Age)

	fmt.Println(len(Name))

	var chineseString string = "你好"
	fmt.Println(len(chineseString))
	fmt.Println(utf8.RuneCountInString(chineseString))

	var myBool bool
	fmt.Println(myBool)

	myBool = true
	fmt.Println(myBool)

	var s1, s2, s3 string = "I", "am", "Saif"
	fmt.Println(s1, s2, s3)

	str1, str2 := "Learning", "Go"
	fmt.Println(str1, str2)

	str1, str2 = "Learning", "Golang"
	fmt.Println(str1, str2)

	const myConst string = "Constant value"
	fmt.Println(myConst)
	//myConst = "Changing value"
	//const id int
	const pi = 3.1415926
	fmt.Println("pi:", pi)
	//numInt = int(pi) ->cannot do this due to untyped constant
	var num3Float = pi
	fmt.Println(num3Float)
}
