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

	// In ra man hinh
	fmt.Println("Hello Classmates!")

	// Goi thu ham (func)
	// Say("Hi my Friends")

	// Doc tu man hinh

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name:")
	yourname, _ := reader.ReadString('\n')
	fmt.Print("Your name is " + yourname)

	// Khai bao bien
	var (
		err    error
		height float64
		weight float64
	)

	height, err = readNumberFromKeyboard("Chiều cao của bạn : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	height, err = readNumberFromKeyboard("Cân nặng của bạn : ")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bmi := CalculateBMI(height, weight)

	fmt.Print("Your BMI index is ", bmi)
}

// Tên hàm viết thường thì ý nghĩa là hàm chỉ dùng trong local package - private, còn viết hoa thì có thể dùng ngoài package - public
func CalculateBMI(height float64, weight float64) (index float64) {
	return weight / (height * height)
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
