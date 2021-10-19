## Chủ đề : GORM + MySQL
Dựa vào bài quan hệ many-to-many : https://github.com/TechMaster/go/tree/main/12/many-to-many

Viết tiếp các function để hoàn thiện các yêu cầu sau:

1.Sửa tên thành viên theo ID

2.Sửa tên club theo ID

3.Xóa 1 thành viên bất kỳ theo ID (hoặc name)

4.Trả về danh sách tên thành viên và số lượng club mà thành viên đó tham gia

Ví dụ kết quả

Name	Clubs  
Alice	2  
Bob	3  
Anna	2  
John	1  
5.Trả về danh sách tên club cùng sô lượng các thành viên đã tham gia

Ví dụ kết quả

Name	Members  
Sport	2  
Music	3  
Math	1  
6.Chỉ ra thông tin club có số lượng thành viên tham gia nhiều nhất

(Các yêu cầu trên có viết unit test trong file test/repo_test.go)

