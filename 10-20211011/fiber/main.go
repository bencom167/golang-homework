/*
	Chủ đề : GoPG + Postgresql
	Yêu cầu:

	Đề bài : học viên tham khảo lại bài tập buổi 7 ( https://techmaster.vn/user/dashboard/bai-tap/c4o6gkv0k7qjcrne2i90/1104)
	Sử dụng Fiber Framework
	Các API thao tác trực tiếp với postgresql (CRUD) để trả về dữ liệu
*/

package main

import (
	"homework/repo"
)

func main() {
	// Khởi tạo dữ liệu cho lần chạy đầu tiên
	//if err := repo.InitUsersList(); err != nil {
	//	fmt.Println(err)
	//}

	// Thao tác với Repo
	repo.RepoUsersWithMiddleware("8080")

	//repo.DemoDB()
}
