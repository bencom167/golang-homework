package main

import "homework/watcher"

func main() {
	/* Tác vụ 1
	1. Tải về 2 file binary `ffmpeg` và `ffprobe` để chạy
	2. Đọc đường dẫn thư mục chứa mp4 và thư mực chứa kết quả băm từ file cấu hình TXT, JSON, YAML...
	3. Quét tất cả file mp4 để băm
	4. Báo lỗi nếu gặp file mp4 không đúng định dạng
	5. Sinh tên mới độ dài 8 ký tự cho file băm
	Hàm băm với input là file JSON/YAML
	*/
	//hashvideo.HashAllVideo("./input.yaml")

	/* Tác vụ 2
	1. Tất cả chức năng Tác vụ 1
	2. Theo dõi thư mục upload, nếu có file mới đc tạo (upload) thì sẽ băm tiếp
	3. Ghi log kết quả băm thành công ra file yaml
	4. Khi chạy chương trình, đọc file yaml trên để không băm lại những video đã băm
	*/
	watcher.HashAllVideo("./input.yaml")

}
