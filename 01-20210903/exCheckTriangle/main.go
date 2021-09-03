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
	var (
		err error
		a   float64
		b   float64
		c   float64
	)

	a, err = readNumberFromKeyboard("Cạnh 1 : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	b, err = readNumberFromKeyboard("Cạnh 1 : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c, err = readNumberFromKeyboard("Cạnh 1 : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Kiểm tra điều kiện 3 cạnh một tam giac
	if isTriangle(a, b, c) {
		fmt.Print("Ba số có thể là cạnh của một tam giác.")
		return
	}

	fmt.Print("Ba số không thể là cạnh của một tam giác.")
}

func isTriangle(a float64, b float64, c float64) bool {
	return (a+b > c) && (a+c > b) && (b+c > a)
}

/*
Hàm đọc từ bàn phím
*/
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
