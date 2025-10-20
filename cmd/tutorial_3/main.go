package main

import (
	"errors"
	"fmt"
)

func main() {
	greetMsg := "Perform arithmetic operations"
	printMe(greetMsg)
	var num = 8
	var den = 0
	var err error
	var divResult float64
	var remainder int
	divResult, remainder, err = division(num, den)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The divison of %v by %v is %v with remainder %v \n", num, den, divResult, remainder)
	}
	var mulResult float64
	mulResult = multiplication(num, den)
	fmt.Printf("The multiplication %v and %v is %v \n", num, den, mulResult)
	fmt.Println("Enter num 1 and num 2:")
	var num1, num2 int
	_, err = fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var choice int
	fmt.Println("Enter your choice: (1 for addition 2 for subtraction) ")
	_, err = fmt.Scan(&choice)
	if err != nil {
		fmt.Println(err.Error())
	}
	switch choice {
	case 1:
		fmt.Println("The addition is", addition(num1, num2))
	case 2:
		fmt.Println("The subtraction is", subtraction(num1, num2))
	default:
		fmt.Println("Choose 1 or 2!")
	}

}

func printMe(greet string) {
	fmt.Println(greet)

}

func division(num int, den int) (float64, int, error) {
	var err error
	if den == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	return float64(num) / float64(den), num % den, err
}

func multiplication(num1 int, num2 int) float64 {
	return float64(num1 * num2)
}

func addition(num1 int, num2 int) float64 {
	return float64(num1 + num2)
}
func subtraction(num1 int, num2 int) float64 {
	return float64(num1 - num2)
}
