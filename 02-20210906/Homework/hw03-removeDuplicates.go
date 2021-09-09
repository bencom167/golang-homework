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
	//numArr := []int{}
	//numArr := []int{1}
	//numArr := []int{1, 1}
	//numArr := []int{1, 2}
	//numArr := []int{1, 2, 5, 2, 6, 2, 5}

	// Testcase mảng số sinh ngẫu nhiên (30 phần tử, 0 <= giá trị <= 100)
	numArr := randomIntArray(100, 30)

	// In mảng ban đầu
	fmt.Println(numArr)

	// Thực hiện tối ưu
	numArr = removeDuplicates(numArr)

	// In sau khi thuc hien xoa trung
	fmt.Println(numArr)
}

/*
	Hàm xoá các phần tử mảng trùng giá trị
*/
func removeDuplicates(numArr []int) []int {
	for i := 0; i < len(numArr); i++ { // Duyệt từng phần tử mảng từ đầu
		for j := len(numArr) - 1; j > i; j-- { // Với mỗi phần tử, duyệt từ cuối mảng lên
			if numArr[i] == numArr[j] { // Gặp trùng giá trị thì xoá ở phần tử gần cuối mảng hơn
				numArr = removeItemSliceNotKeepOrder(numArr, j)
			}
		}
	}
	return numArr
}

/*
	Hàm xoá 1 phần tử mảng, không giữ thứ tự
*/
func removeItemSliceNotKeepOrder(numArr []int, index int) []int {
	numArr[index] = numArr[len(numArr)-1] // Gán phần tử cuối cùng vào ô cần xoá
	return numArr[:len(numArr)-1]         // Cắt bớt phần tử cuối của mảng
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
