/*
	Hãy lập trình ứng dụng tương tự như lệnh tree liệt kê cây thư mục. Chú ý trên mạng có nhiều gợi ý thuật toán.
	Chỉ cần đọc, hiểu, áp dụng được vào ứng dụng Go là tốt nghiệp lớp cấu trúc dữ liệu giải thuật Golang rồi.

	Code giải quyết được với các yêu cầu:
		- Code chạy cho MacOS, Linux
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	directoryOrigin string
	depthOrigin     int
)

func main() {
	// Nhập dữ liệu thư mục và độ sâu của tree
	directoryOrigin, depthOrigin = inputData()

	// Kiểm tra thư mục
	directoryOrigin, valid := checkDirectory(directoryOrigin)
	if !valid {
		return
	}

	// In title thư mục
	fmt.Printf("%v in tree with %v depth\n", directoryOrigin, depthOrigin)
	// Duyệt và in thư mục/file dạng tree (bắt đầu thì prefix là rỗng)
	readDirectory(directoryOrigin, "", depthOrigin)
}

/*
	Hàm nhập dữ dữ liệu, tạm thời fix trong code :)
*/
func inputData() (string, int) {
	// Thư mục cần duyệt
	//directory := "/Users/bendn/code/golang/02-20210906"
	//directoryOrigin := "/Users/bendn/code/golang/"
	//directoryOrigin := "/Users/bendn/code/"
	directory := "."

	depth := 3 // Độ sâu thư mục cần duyệt

	return directory, depth
}

/*
	Hàm kiểm tra thư mục xem có đúng ko, có quyền đọc ko
*/
func checkDirectory(directory string) (string, bool) {
	if directory[len(directory)-1:] != "/" { // Nếu không tận cùng bằng dấu / thì bổ sung
		directory += "/"
	}

	_, err := ioutil.ReadDir(directory) // Đọc thử thư mục
	if err != nil {                     // Nếu có lỗi thì in thông báo lỗi
		fmt.Println("Lỗi khi đọc thư mục:")
		log.Fatal(err)
		return directory, false
	}
	return directory, true
}

/*
	Hàm đệ quy, duyệt thư mục, đi sâu thêm vào thư mục con nếu vẫn chưa hết độ sâu
	directory	: Thư mục cần duyệt
	prefix		: Xây dựng trong quá trình duyệt, để in tree đúng format cho mỗi file/thư mục con
	depth		: Độ sâu còn lại của quá trình duyệt
*/
func readDirectory(directory string, prefix string, depth int) {
	files, err := ioutil.ReadDir(directory) // Đọc thư mục, trả lại danh sách File Infor
	if err != nil {                         // Nếu có lỗi thì in thông báo lỗi
		log.Fatal(err)
	}

	count := 0 // Biến đếm số lượng file/thư mục con trong thư mục mẹ
	for _, file := range files {
		count++                                // Tăng biến đếm thêm 1
		isEnd := count == len(files)           // Xem xét có phải là file/thư mục con cuối cùng
		printEntry(file.Name(), prefix, isEnd) // In file/thư mục con

		if file.IsDir() && depth > 1 { // Nếu là thư mục và vẫn chưa duyệt hết độ sâu yêu cầu thì duyệt tiếp
			if isEnd {
				readDirectory(directory+file.Name()+"/", prefix+"    ", depth-1) // Nếu là file/thư mục con cuối cùng, duyệt với prefix không có nhánh
			} else {
				readDirectory(directory+file.Name()+"/", prefix+"│   ", depth-1) // Nếu không phải là file/thư mục con cuối cùng, duyệt với prefix có nhánh
			}
		}
	}
}

/*
	Hàm in tên file/thư mục theo format
*/
func printEntry(fileName string, preFix string, isEnd bool) {
	fmt.Printf("%v", preFix)

	if isEnd { // Nếu là phần thử cuối thi in móc vuông
		fmt.Printf("└── %v\n", fileName)
	} else { // Nếu không là phần thử cuối thi in móc chữ T
		fmt.Printf("├── %v\n", fileName)
	}
}
