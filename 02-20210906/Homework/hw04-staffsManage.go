/*
	Bài 4: Một nhân viên trong công ty bao gồm các thuộc tính sau : Tên, Hệ số lương, Tiền trợ cấp
	Tạo 1 mảng nhân viên (số lượng tuỳ ý) và thực hiện các chức năng sau:

	Sắp xếp tên nhân viên tăng dần theo bảng chữ cái
	Sắp xếp nhân viên theo mức lương giảm dần (lương = Hệ số lương * 1.500.000 + Tiền trợ cấp)
	Lấy ra danh sách nhân viên có mức lương lớn thứ 2 trong mảng nhân viên
*/

package main

import "fmt"

type Staff struct {
	FirstName   string  // Tên
	LastName    string  // Họ
	SalaryRatio float64 // Hệ số lương
	BonusMoney  float64 // Trợ cấp
}

func main() {
	//staffs := []Staff{}
	//staffs := []Staff{{"Toan", "Vo", 3.0, 4500000}}
	//staffs := []Staff{{"Toan", "Vo", 3.0, 4500000}, {"Quan", "Dang", 3.0, 4500000}}
	//staffs := []Staff{{"Toan", "Vo", 3.0, 4500000}, {"Quan", "Dang", 5.7, 3000000}}
	staffs := []Staff{
		{"Toan", "Vo", 3.0, 4500000},
		{"Quan", "Dang", 5.7, 3000000},
		{"Khoa", "Tran", 4.5, 2000000},
		{"Duy", "Nguyen", 5.4, 5000000},
		{"Tan", "Nguyen", 4.8, 4500000},
		{"Bang", "Nguyen", 3.3, 3000000},
		{"Quan", "Phuong", 5.7, 4150000},
		{"Cuong", "Trinh", 4.8, 5500000},
		{"Ben", "Dang", 3.9, 5000000},
		{"Luan", "Vo", 4.8, 4000000}}

	// In danh sách gốc
	//printNiceFormat(staffs, "Danh sách gốc :")

	// Sắp xếp theo tên/họ và in
	printNiceFormat(sortByName(staffs), "Danh sach sap xep theo ten :")

	// Sắp xếp giảm dần theo lương
	staffsSortBySalary := sortBySalary(staffs)

	// In danh sách sau khi sắp xếp giảm dần theo lương
	printNiceFormat(staffsSortBySalary, "Danh sach sap xep giam dan theo luong :")

	// In danh sách nhân viên có mức lương lớn thứ 2, đầu vào là danh sách đã sắp xếp lương giảm dẩn
	printNiceFormat(max2StaffsSalary(staffsSortBySalary), "Danh sach nhân viên có mức lương lớn thứ 2 :")

}

/*
	Hàm sắp xếp danh sách theo Tên/Họ
*/
func sortByName(staffs []Staff) []Staff {
	for i := 0; i < len(staffs); i++ { // Sử dụng sắp xếp nổi bọt
		for j := i + 1; j < len(staffs); j++ {
			if staffs[i].FirstName > staffs[j].FirstName || // Điều kiện tên lớn hơn
				(staffs[i].FirstName == staffs[j].FirstName && staffs[i].LastName > staffs[j].LastName) { // Hoặc tên bằng, họ lớn hơn
				staffTemp := staffs[i] // Đảo vị trí 2 nhân viên
				staffs[i] = staffs[j]
				staffs[j] = staffTemp
			}
		}
	}
	return staffs
}

/*
	Hàm sắp xếp danh sách giảm dần theo lương
*/
func sortBySalary(staffs []Staff) []Staff {
	for i := 0; i < len(staffs); i++ { // Sử dụng sắp xếp nổi bọt
		for j := i + 1; j < len(staffs); j++ {
			if calculateSalary(staffs[i]) < calculateSalary(staffs[j]) { // Điều kiện lương nhỏ hơn
				staffTemp := staffs[i] // Đảo vị trí 2 nhân viên
				staffs[i] = staffs[j]
				staffs[j] = staffTemp
			}
		}
	}
	return staffs
}

/*
	Hàm lấy danh sách nhân viên có lương cao thứ 2
*/
func max2StaffsSalary(staffs []Staff) []Staff {
	max2Staffs := []Staff{} // Khởi tạo mảng rỗng

	if len(staffs) < 2 { // Danh sách nhân viên có 0, 1 phần tử -> không có người lương cao thứ 2
		return max2Staffs
	}

	maxSalary := calculateSalary(staffs[0]) // Lương max
	max2Index := -1                         // Vị trí bắt đầu của nhân viên có lương max2
	for i := 1; i < len(staffs); i++ {
		if calculateSalary(staffs[i]) < maxSalary {
			max2Index = i
			break
		}
	}

	if max2Index == -1 { // Tất cả nhân viên bằng lương nhau -> không có người lương cao thứ 2
		return max2Staffs
	}

	max2Salary := calculateSalary(staffs[max2Index]) // Lương max2
	for i := max2Index; i < len(staffs); i++ {
		if max2Salary == calculateSalary(staffs[i]) { // Nếu lương = max2 thì đẩy vào danh sách
			max2Staffs = append(max2Staffs, staffs[i])
		}
	}
	return max2Staffs
}

// Hàm tính lương từ hệ số lương và trợ cấp
func calculateSalary(staff Staff) float64 {
	return staff.SalaryRatio*1500000 + staff.BonusMoney
}

// Hàm in danh sách nhân viên
func printNiceFormat(staffs []Staff, msg string) {
	fmt.Println("=========================")
	fmt.Println(msg)
	for i := 0; i < len(staffs); i++ {
		fmt.Printf(" %v. \t%v %v \t%v \t%v\n", i+1, staffs[i].FirstName, staffs[i].LastName, staffs[i].SalaryRatio, staffs[i].BonusMoney)
	}
}
