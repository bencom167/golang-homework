/*
	Bài 1: Viết function tìm ra số lớn thứ nhì trong mảng các số.
	Ví dụ: max2Numbers([2, 1, 3, 4]) => 3

	Code giải quyết được với các yêu cầu:
		- Mảng số nguyên không âm
		- Nếu nhiều phần từ bằng nhau, đều là max thì max2 vẫn là giá trị lớn thứ nhì
*/

package main

import (
	"fmt"
)

func hw01Max2Numbers() {
	//numArray := []int{} // Testcase mảng rỗng
	//numArray := []int{10} // Testcase mảng 1 phần tử
	//numArray := []int{2, 1, 3, 4}
	//numArray := []int{23, 45, 43, 91, 10, 19, 9, 21, 63, 72}
	//numArray := []int{23, 45, 43, 91, 10, 19, 9, 21, 63, 72, 91, 72} // Testcase mảng có nhiều phần tử =max và =max2

	// Testcase mảng số sinh ngẫu nhiên (30 phần tử, 0 <= giá trị <= 100)
	numArray := randomIntArray(100, 30)

	// Tìm giá trị max2
	max2 := max2Numbers(numArray)

	// In mảng
	fmt.Println(numArray)

	// In gia tri lon thu nhi
	if max2 == -1 {
		fmt.Println("Không tìm thấy phần tử lớn thứ nhì trong mảng!")
	} else {
		fmt.Println("Phần tử lớn thứ nhì trong mảng là : ", max2)
	}
}

/*
	Hàm tìm số lớn thứ nhì trong mảng số nguyên không âm
	Chấp nhận các phần tử trong mảng có giá trị bằng nhau
*/
func max2Numbers(numArray []int) int {

	// Trả lại not found nếu mảng rỗng
	if len(numArray) == 0 {
		return -1
	}

	// Tìm max
	max := numArray[0]
	for i := 1; i < len(numArray); i++ {
		if max < numArray[i] {
			max = numArray[i]
		}
	}

	max2 := -1 // Khởi tạo giá trị not found
	// Tìm max2, là số lớn nhất nhưng vẫn phải nhỏ hơn max
	for i := 1; i < len(numArray); i++ {
		if max2 < numArray[i] && numArray[i] < max {
			max2 = numArray[i]
		}
	}

	return max2
}
