package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	var myString = "Go语言"
	fmt.Println(myString[7])
	fmt.Printf("%v %T\n", myString[3], myString[3])
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of myString is %v\n", len(myString))

	var myRune = []rune(myString)
	fmt.Printf("%v %T\n", myRune[1], myRune[1])
	for i, v := range myRune {
		fmt.Println(i, v)
	}
	fmt.Printf("The length of myRune is %v\n", len(myRune))

	var myRune2 = '$'
	fmt.Printf("%c : %v\n", myRune2, myRune2)

	var myString2 = string(myRune)
	fmt.Println(myString2)

	var strSlice = []string{"S", "a", "i", "f"}
	var catStr = ""
	fmt.Printf("The memory address of catStr is %v\n", &catStr)
	for i := range strSlice {
		catStr += strSlice[i]
		fmt.Printf("Iteration %d: header=%p, data=%p\n", i, &catStr, unsafe.StringData(catStr))
	}
	fmt.Println(catStr)
	//catStr[0] = 'l' -->cannot modify
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
		str2 := strBuilder.String()
		fmt.Printf("Iteration %d: data=%p, str=%c\n", i, unsafe.StringData(str2), str2[i])
	}

}
