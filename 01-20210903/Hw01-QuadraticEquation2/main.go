/*
	Giải phương trình bậc 2
	Nhập vào ba số a, b, c kiểu float64 hãy giải phương trình bậc 2.
*/

package main

import (
	// Thu vien chuan
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Khai bao bien
	var (
		a float64
		b float64
		c float64
	)

	// Nhap tham so
	a, b, c = inputPara()

	// Tinh gia tri delta
	delta := math.Pow(b, 2) - 4*a*c

	// Ghi ket qua
	switch {
	case delta < 0:
		fmt.Println("Phuong trinh vo nghiem!")
	case delta == 0:
		fmt.Println("Phuong trinh co nghiem kep: x1 = x2 = ", (-1*b)/(2*a))
	default:
		fmt.Println("Phuong trinh co nghiem :")
		fmt.Println("		x1 = ", ((-1*b)+math.Sqrt(delta))/(2*a))
		fmt.Println("		x2 = ", ((-1*b)-math.Sqrt(delta))/(2*a))
	}
}

// Nhap tham so a,b,c
func inputPara() (a float64, b float64, c float64) {

	a = readNumberFromKeyboard("Nhap tham so a : ")
	for {
		if a != 0 {
			break // Nhap lieu den khi a != 0
		}
		a = readNumberFromKeyboard("Nhap lai tham so a : ")
	}

	b = readNumberFromKeyboard("Nhap tham so b : ")
	c = readNumberFromKeyboard("Nhap tham so c : ")

	return a, b, c
}

// Ham do so tu ban phim
func readNumberFromKeyboard(msg string) (result float64) {
	var err error

	reader := bufio.NewReader(os.Stdin)

	// Lap den khi nhap lieu dung thi thoi
	for {
		fmt.Print(msg)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")

		if result, err = strconv.ParseFloat(str, 64); err == nil {
			break
		}
		fmt.Println(err.Error())
	}

	return result
}
