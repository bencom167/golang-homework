## Yêu cầu
- Porting fiber-postgres sang sử dụng gin-postgres
- Đánh giá benchmark API giữa 3 framework : fiber, iris (đã có sẵn code) và gin khi cùng kết nối với CSDL postgres
- Học viên gửi lại link github phần porting sang gin-postgres và bảng đánh giá như bên dưới trong phần README.md
- Mỗi framework đánh giá 3 lần đối với mỗi API

Bảng đánh giá:

|Method    |API          |   Gin                |      Iris        |    Fiber            |    Command            |
|----------|-------------|------------------|------------------|---------------------|-----------------------|
|GET       |/users       | 6086, 5974, 5934 | 5174, 5140, 5183 | 6921, 6798, 6719 | hey -n 60000 -c 50 -m GET http://localhost:3000/users |
|GET       |/users/{id}  | 6937, 6461, 7004 | 6027, 6862, 6680 | 6916, 7064, 6838 | hey -n 60000 -c 50 -m GET http://localhost:3000/users/oIlE6mcq |
|PUT       |/users/{id}  | 210, 223, 330, 1341, 402 | 258, 299, 180, 754, 305 | 365, 766, 313, 264, 206 | hey -n 10000 -c 50 -m PUT -H "Content-Type: application/json" -d '{"full_name" : "Ngo Van A", "email" : "b@gmail.com", "phone" : "0123456789"}' http://localhost:3000/users/IepsxLUd |
|POST      |/users       | 4771, 4895, 4864 | 1553, 4836, 4389, 1733 | 1796, 4762, 4880, 1375 | hey -n 10000 -c 50 -m POST -H "Content-Type: application/json" -d '{"full_name" : "Tran Thi C", "email" : "c@gmail.com", "phone" : "0123456789"}' http://localhost:3000/users |
