/*
	Ứng dụng Calendar
	Ở trên terminal, nếu bạn gõ lệnh:
		$ cal
	Kết quả in ra lịch tháng này:
			September 2021
		Su Mo Tu We Th Fr Sa
				1  2  3  4
		5  6  7  8  9 10 11
		12 13 14 15 16 17 18
		19 20 21 22 23 24 25
		26 27 28 29 30
	Hãy viết ứng dụng Golang thực sự (không dùng exec.Command)
	sử dụng thư viện fmt, time để viết lại ứng dụng cal

	Gợi ý: hãy chia ứng dụng thành 3 phần:

	In ra tháng và năm hiện tại September 2021
	In ra Su Mo Tu We Th Fr Sa
	Tìm ngày 1 tháng hiện tại ở thứ mấy bắt đầu in ra cho đến ngày cuối cùng của tháng.
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	curTime := inputTime()      //Nhập thời gian
	printMonthCalendar(curTime) // In ra lịch tháng hiện tại
}

func printMonthCalendar(curTime time.Time) {
	curYear, curMonth, curDay := curTime.Date() // Lấy thông tin ngày tháng năm

	monthYear := curMonth.String() + " " + string(strconv.Itoa(curYear))
	for len(monthYear) < 20 { // Tạo string tháng + năm có độ dài tương đương hàng ngày trong tuần (20 kí tự), để căn giữa khi in
		monthYear = " " + monthYear + " " // Thêm khoảng trắng đều vào đầu cuối
	}
	fmt.Println(monthYear)              // In tháng + năm
	fmt.Println("Su Mo Tu We Th Fr Sa") // In ngày trong tuần

	firstDateOfMonth := time.Date(curYear, curMonth, 1, 0, 0, 0, 0, curTime.Location()) // Ngày đầu tiên của tháng
	lastDateOfMonth := firstDateOfMonth.AddDate(0, 1, -1)                               // Ngày cuối cùng của tháng

	endOfWeek := 0                                                          // Giả sử cuối tuần là hết thứ 7, xem xét xuống dòng khi in
	for i := 1 - int(curTime.Weekday()); i < lastDateOfMonth.Day()+1; i++ { // Lặp để in ngày trong tháng, tính cả khoảng trắng đầu tháng
		switch {
		case i < 1: // In khoảng trắng khi những ngày đầu tháng không trùng với đầu tuần
			fmt.Printf("   ")
		case i == curDay: // Highlight, in màu đỏ ngày hiện tại, độ rộng 2 ký tự, căn phải
			fmt.Printf("\033[31m%2v \033[0m", i)
		default: // In ngày, độ rộng 2 ký tự, căn phải
			fmt.Printf("%2v ", i)
		}

		endOfWeek++
		if endOfWeek == 7 { // Nếu là cuối tuần (thứ 7) thì xuống dòng
			fmt.Printf("\n")
			endOfWeek = 0
		}
	}
	fmt.Printf("\n") // In xuống dòng cuối tháng
}

func inputTime() time.Time {
	return time.Now()
	//return time.Date(2021, 1, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 1
	//return time.Date(2021, 2, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 2
	//return time.Date(2021, 3, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 3
	//return time.Date(2021, 4, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 4
	//return time.Date(2021, 5, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 5
	//return time.Date(2021, 6, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 6
	//return time.Date(2021, 7, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 7
	//return time.Date(2021, 8, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 8
	//return time.Date(2021, 9, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 9
	//return time.Date(2021, 10, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 10
	//return time.Date(2021, 11, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 11
	//return time.Date(2021, 12, 13, 0, 0, 0, 0, time.Now().Location()) // Tháng 12
}
