package staff

import (
	"fmt"
	"math/rand"
	"time"
)

type iStaff interface {
	Print()
	CalculateSalary() int
	SetEmpID(int)
	SetBasicPay(int)
	SetPF(int)
	Clone() iStaff
}

func BuildStaffs(numPermanent int, numContract int) []iStaff {
	var staffList []iStaff
	rand.Seed(time.Now().UnixNano()) // Khởi tạo nhân sinh ngẫu nhiên

	if numPermanent > 0 { // Sử dụng phương thức Clone, điều kiện là số lượng nhân viên chính thức > 0
		staff := Permanent{ // Tạo mới nhân viên chính thức đầu tiên
			empId:    randomEmpID(),
			basicpay: randomBasicPay(),
			pf:       randomPF(),
		}
		staffList = append(staffList, &staff) // Đẩy vào danh sách

		for i := 1; i < numPermanent; i++ { // Lặp với số lượng n-1
			staffClone := staff.Clone()              // Clone nhân viên mới
			staffClone.SetEmpID(randomEmpID())       // Update mã nhân viên mới
			staffClone.SetBasicPay(randomBasicPay()) // Update lương mới
			staffClone.SetPF(randomPF())             // Update thưởng mới

			staffList = append(staffList, staffClone) // Đẩy vào danh sách
		}
	}

	for i := 0; i < numContract; i++ { // Không dùng Clone mà tạo mới lần lượt với các nhân viên hợp đồng
		staff := Contract{ // Tạo mới nhân viên hợp đồng
			empId:    randomEmpID(),
			basicpay: randomBasicPay(),
		}
		staffList = append(staffList, &staff) // Đẩy vào danh sách
	}

	return staffList
}

/*
	Hàm in danh sach tất cả nhân viên trong mảng staffList
*/
func PrintStaffs(staffList []iStaff) {
	for i := 0; i < len(staffList); i++ {
		fmt.Printf("Thông tin nhân viên %3v ", i+1)
		staffList[i].Print()
	}
}

/*
	Hàm tính tổng lương của tất cả nhân viên trong mảng staffList
*/
func TotalSalary(staffList []iStaff) int {
	totalSalary := 0
	for i := 0; i < len(staffList); i++ {
		totalSalary += staffList[i].CalculateSalary()
	}
	return totalSalary
}

//	Hàm sinh số ID ngẫu nhiên (10000->99999)
func randomEmpID() int {
	return 10000 + rand.Intn(89999)
}

//	Hàm sinh lương ngẫu nhiên (10->50tr)
func randomBasicPay() int {
	return 1000000 * (10 + rand.Intn(40))
}

//	Hàm sinh thưởng ngẫu nhiên (5->20tr)
func randomPF() int {
	return 1000000 * (5 + rand.Intn(15))
}
