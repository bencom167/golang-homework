package main

import (
	// Thu vien chuan
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Khai bao bien
	var value float64
	var (
		err    error
		target float64
	)

	target, err = readNumberFromKeyboard("Nhap so Fibonacci thu : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Su dung loop de tinh
	// value = methodUsingLoop(int(target))

	// Su dung De qui de tinh
	value = methodUsingRecursion(int(target))

	fmt.Println("So Fibonacci thu ", target, " co gia tri : ", value)
}

// Ham su dung loop For thong thuong
func methodUsingLoop(target int) float64 {
	var previousNum = 1.0
	var currentNum = 1.0

	// Voi thu tu so la 1,2 thi tra lai gia tri luon
	if target < 3 {
		return 1.0
	}

	// Tu 3 tro len thi lap tinh
	for i := 2; i < target; i++ {
		nextNum := previousNum + currentNum
		previousNum = currentNum
		currentNum = nextNum
	}

	return currentNum
}

// Ham De quy
func methodUsingRecursion(target int) float64 {
	if target <= 2 {
		return 1
	}
	// Goi de quy tinh cac so thu tu truoc do
	return methodUsingRecursion(target-1) + methodUsingRecursion(target-2)
}

// Ham do so tu ban phim
func readNumberFromKeyboard(msg string) (result float64, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(msg)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSuffix(str, "\n")

	if result, err = strconv.ParseFloat(str, 64); err != nil {
		fmt.Println(err.Error())
		return
	}

	return result, nil
}
