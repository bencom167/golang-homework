/*
	Bài 2 Cho 1 mảng các chuỗi. Viết function lọc ra các phần tử có độ dài lớn nhất.
	Ví dụ: findMaxLengthElement["aba", "aa", "ad", "c", "vcd"] => ["aba", "vcd"]
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//stringArray := []string{} // Mảng rỗng
	//stringArray := []string{"aba"} // Mảng có 1 phần tử
	//stringArray := []string{"aba", "aa", "ad", "c", "vcd", "dfgd"}

	stringArray := randomStringArray(10, 30) // Sinh mảng 30 string, độ dài string từ 1->10

	// In mảng tất cả string
	fmt.Println(stringArray)

	// In tất cả string có độ dài lớn nhất
	fmt.Println(findMaxLengthElement(stringArray))
}

/*
	Hàm tìm tất cả các string có độ dài lớn nhất
*/
func findMaxLengthElement(stringArray []string) []string {
	var stringArrayMax []string

	// Tìm độ dài lớn nhất của các phần tử trong mảng
	maxLen := 0
	for i := 0; i < len(stringArray); i++ {
		if maxLen < len(stringArray[i]) {
			maxLen = len(stringArray[i])
		}
	}

	// Tìm tất cả các phần tử có độ dài lớn nhất
	for i := 0; i < len(stringArray); i++ {
		if maxLen == len(stringArray[i]) {
			stringArrayMax = append(stringArrayMax, stringArray[i])
		}
	}

	return stringArrayMax
}

/*
	Hàm sinh ngẫu nhiên mảng arrLen phần tử là string, độ dài của một string ngẫu nhiên 1->maxStrLen
*/
func randomStringArray(maxStrLen int, arrLen int) []string {
	stringArray := []string{}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < arrLen; i++ {
		stringArray = append(stringArray, randomString(rand.Intn(maxStrLen)+1))
	}

	return stringArray
}

/*
	Hàm sinh một string có độ dài length
*/
func randomString(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	str := make([]byte, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}
