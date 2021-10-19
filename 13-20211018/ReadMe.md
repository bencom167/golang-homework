## Yêu cầu
- Porting fiber-postgres sang sử dụng gin-postgres
- Đánh giá benchmark API giữa 3 framework : fiber, iris (đã có sẵn code) và gin khi cùng kết nối với CSDL postgres
- Học viên gửi lại link github phần porting sang gin-postgres và bảng đánh giá như bên dưới trong phần README.md
- Mỗi framework đánh giá 3 lần đối với mỗi API

Bảng đánh giá

Method	API	Gin	Iris	Fiber  
GET	/users	Ví dụ: 450, 467, 456		  
GET	/users/{id}			  
PUT	/users/{id}			  
POST	/users			  
