/*
	Đoán số
	Máy tính tự sinh ra một số nguyên dương X >= 0 và <= 100. Lập trình một vòng lặp để người dùng đoán số:
	-	Nếu số đoán lớn hơn X thì in ra "Số bạn đoán lớn hơn X"
	-	Nếu số đoán nhỏ hơn X thì in ra "Số bạn đoán nhỏ hơn X"
	-	Nếu bằng X thì in ra "Bạn đã đoán đúng"
*/

package main

import (
	// Thu vien chuan
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Khai bao bien
	var (
		guessNum int64
		X        int64
	)

	// Khoi tao bien sinh ngau nhien.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Sinh ngau nhien so nguyên dương X >= 0 và <= 100
	X = int64(r.Int31n(100))

	for {
		guessNum = readNumberFromKeyboard("Ban du doan so X la : ")
		switch {
		case guessNum > X:
			fmt.Println("Số bạn đoán lớn hơn X!")
		case guessNum < X:
			fmt.Println("Số bạn đoán nhỏ hơn X!")
		default:
			fmt.Println("Bạn đã đoán đúng!")
			return
		}
	}
}

// Ham do so nguyen tu ban phim
func readNumberFromKeyboard(msg string) (result int64) {
	var err error
	reader := bufio.NewReader(os.Stdin)

	// Lap den khi nhap lieu dung thi thoi
	for {
		fmt.Print(msg)
		str, _ := reader.ReadString('\n')
		str = strings.TrimSuffix(str, "\n")

		if result, err = strconv.ParseInt(str, 10, 64); err == nil {
			break
		}
		fmt.Println(err.Error())
	}

	return result
}
