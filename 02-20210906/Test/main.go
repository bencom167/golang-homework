/*
	Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
	Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Testcase mảng số sinh ngẫu nhiên (30 phần tử, 0 <= giá trị <= 100)
	// numArray := randomIntArray(8, 10)
	//fmt.Println(numArray)

	fmt.Println("Ben" > "Quan")
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
