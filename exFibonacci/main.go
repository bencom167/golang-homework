package main

import (
	// Thu vien chuan

	"fmt"
)

func main() {

	// Khai bao bien
	var (
		err       error
		firstNum  float64
		secondNum float64
	)

	firstNum = 1
	secondNum = 2

	bmi := returnNextNumberFibonacci(firstNum, secondNum)

	fmt.Print("Your BMI index is ", bmi)
}

/*
Hàm đọc từ bàn phím
*/
func returnNextNumberFibonacci(a float64, b float64) (result float64) {
	result = a + b
	return result
}
