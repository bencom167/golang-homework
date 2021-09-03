package main

import (
	// Thu vien chuan

	"fmt"
)

func main() {

	// Khai bao bien
	var (
		firstNum  float64
		secondNum float64
	)

	firstNum = 1
	secondNum = 1

	target := 10

	number := returnNextNumberFibonacci(firstNum, secondNum)

	fmt.Print("So Fibonacci thu ", target, " : ", number)
}

/*
Hàm đọc từ bàn phím
*/
func returnNextNumberFibonacci(a float64, b float64) (result float64) {
	result = a + b
	return result
}
