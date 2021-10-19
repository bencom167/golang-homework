1. Ứng dụng Calendar
Ở trên terminal, nếu bạn gõ lệnh

$ cal
Kết quả in ra lịch tháng này

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

2. Docker SDK
Hãy sử dụng thư viện https://docs.docker.com/engine/api/sdk/ để viết ứng dụng có chức năng sau:

Hiển thị danh sách các container gồm các thông tin

Container ID
Container Name
Image Name
Port (host:container)
Trạng thái: running hay stop
In ra dấu nhắc để người dùng lựa chọn lệnh:

F5 để toggle liệt kê chỉ những container đang chạy, chỉ những container đang dừng, hoặc tất cả các container
Người dùng gõ stop container_id hoặc stop container_name thì dừng container đang chạy
Người dùng gõ start container_id hoặc start container_name thì khởi động container đang dừng

3. Dành cho người xuất sắc. Chỉ cần làm bài này, khỏi làm bài 1 và 2.
Hãy lập trình ứng dụng tương tự như lệnh tree liệt kê cây thư mục. Chú ý trên mạng có nhiều gợi ý thuật toán.
Chỉ cần đọc, hiểu, áp dụng được vào ứng dụng Go là tốt nghiệp lớp cấu trúc dữ liệu giải thuật Golang rồi.

Thầy Cường cũng chưa nghĩ ra bài này.

.
├── 01
│   ├── Github_Language_Stats.jpg
│   ├── ReadMe.md
│   ├── go.mod
│   ├── go.sum
│   ├── main
│   ├── main.go
│   ├── md
│   ├── structs
│   └── triangle_test.go
├── 02
│   ├── ReadMe.md
│   ├── array2d.go
│   ├── array_loop.go
│   ├── conditional_break_point.jpg
│   ├── demo.md
│   ├── go.mod
│   ├── main.go
│   ├── map.go
│   ├── remove_slice_bench_test.go
│   ├── return_error.go
│   ├── slice.go
│   ├── sort.go
│   ├── struct.go
│   └── switch.go