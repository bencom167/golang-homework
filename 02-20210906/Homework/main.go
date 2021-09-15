package main

import (
	"math/rand"
	"time"
)

func main() {
	// Chạy tương ứng bài 1,2,3,4
	hw := 4

	switch {
	case hw == 1:
		hw01Max2Numbers()
	case hw == 2:
		hw02FindMaxLengthElement()
	case hw == 3:
		hw03RemoveDuplicates()
	default:
		hw04StaffsManage()
	}
}

/*
	Hàm sinh mảng số nguyên không âm ngẫu nhiên, count phần từ, giá trị phần tử <= max
*/
func randomIntArray(max int, count int) []int {
	var numArray []int

	// Khoi tao bien sinh ngau nhien.
	rand.Seed(time.Now().UnixNano())

	// Sinh ngau nhien so nguyên không âm <= max, bổ sung vào mảng
	for i := 0; i < count; i++ {
		numArray = append(numArray, int(rand.Intn(max)))
	}
	return numArray
}
