/*
	Bài 3 Viết function remove những phần tử bị trùng nhau trong mảng
	Ví dụ: removeDuplicates([1,2,5,2,6,2,5]) => [1,2,5,6]
*/

package main

import (
	"fmt"
)

func hw03RemoveDuplicates() {
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
