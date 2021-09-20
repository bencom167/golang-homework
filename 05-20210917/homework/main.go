package main

import (
	"fmt"
	"homework/docker"
	"homework/staff"
	"time"
)

func main() {

	/***** Bài 1 *****/

	// Xây dựng iDockerClient với tham số isAllContainer: true -> lấy tất cả container, false -> chỉ running container
	docClient := docker.BuildIDockerClient(true)

	// In danh sách Container
	docClient.ListAll()

	// Khởi động một Container
	conID := "e8659b128d"
	err := docClient.StartContainer(conID)
	if err == nil {
		fmt.Println("Start container successfully: ", conID)
	} else {
		fmt.Println(err)
	}

	// Dừng một Container
	conID = "3097eee86d"
	err = docClient.StopContainer(conID)
	if err == nil {
		fmt.Println("Stop container successfully: ", conID)
	} else {
		fmt.Println(err)
	}

	// In lại danh sách Container
	time.Sleep(time.Second * 5)
	docClient.ListAll()

	/***** Bài 2 *****/
	// Hàm BuildStaffs(n, m) tạo danh sách với n nhân viên chính thức, m nhân viên hợp đồng
	staffs := staff.BuildStaffs(7, 3)

	// In danh sách nhân viên
	staff.PrintStaffs(staffs)

	// Tổng tiền trả lương hàng tháng
	fmt.Printf("Tổng tiền công ty phải trả lương cho nhân viên: %12v\n", staff.TotalSalary(staffs))
}
