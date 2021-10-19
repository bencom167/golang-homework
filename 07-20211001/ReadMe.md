Cho struct user có cấu trúc như sau

type User struct {
    Id int
    FullName string
    Email string
    Phone string
    Age int
    Sex string
}
Yêu cầu 1: Tạo user repository bao gồm các chức năng sau

METHOD	API	Chức năng	Dữ liệu trả về  
GET	/users	Lấy danh sách user	Trả về danh sách users  
POST	/users	Tạo user mới	Trả về user mới sau khi tạo  
GET	/users/id	Chi tiết user	Trả về thông tin user  
PUT	/users/id	Cập nhật user	Trả về thông tin user sau khi cập nhật  
DELETE	/users/id	Xóa user	Trả về danh sách user sau khi xóa  

Note:

- Với API GET /users : Yêu cầu có sử dụng query string để lọc user theo ít nhất 2 tiêu chí. Ngoài các tiêu chí về thuộc tính có thể áp dụng các tiêu chí khác như là: phân trang (page, limit), sắp xếp (sort))
- Dữ liệu trả về của các API đều là kiểu json
- Sử dụng thuần package net/http

Yêu cầu 2:  
Viết middleware để in ra thông tin của request trong terminal. Thông tin request bao gồm [METHOD] - URL (ví dụ [GET] - /users, [POST] - /users)