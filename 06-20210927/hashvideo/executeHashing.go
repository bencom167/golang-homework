package hashvideo

import (
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/TechMaster/eris"
	"gopkg.in/yaml.v2"
)

var (
	UploadDir    string
	HLSDir       string
	HashedResult string
)

/*
	Hàm băm toàn bộ file mp4 có trong thư mục upload
*/
func HashAllVideo(inputFile string) error {
	// Đọc lấy thông tin thư mục input/output, nếu có lỗi thì in thông báo
	if err := inputReader(inputFile); err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(UploadDir) // Đọc thư mục, trả lại danh sách File Infor
	if err != nil {                         // Nếu có lỗi thì in thông báo lỗi
		log.Fatal(err)
	}

	convert_result_channel := make(chan error)
	handle_convert_error(convert_result_channel)

	wait_group := sync.WaitGroup{}
	count := 0

	for _, file := range files {
		if isValidMP4(file) {
			count++                    // Tăng biến đếm thêm 1
			videoID := randomString(8) // Tạo video ID với độ dài 8 kí tự
			log.Printf("Bắt đầu convert: %v <- %v\n", videoID, file.Name())

			// Hàm này chạy rất lâu do đó phải cho vào go routine để chạy
			wait_group.Add(1)
			go hashVideoToHLS(videoID, file.Name(), &wait_group, convert_result_channel)
		}
	}
	wait_group.Wait() // Đợi các routine hoàn thành hết

	log.Printf("Hoàn thành check/convert %v MP4 videos!\n", count)
	return nil
}

/*
	Hàm đọc file input, cung cấp thông tin thư mục A chưa file mp4 và thư mục B chứa định dang sau khi băm
*/
func inputReader(inputFile string) error {
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi đọc file config!")
	}

	data := make(map[string]string)
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi đọc file config format yaml!")
	}

	UploadDir = data["UploadDir"]
	HLSDir = data["HLSDir"]
	HashedResult = data["HashedResult"]

	return nil
}

/*
	Hàm đọc kiểm tra xem có phải mp4 hay không
*/
func isValidMP4(file fs.FileInfo) bool {
	extension := file.Name()[strings.LastIndex(file.Name(), ".")+1:]
	return strings.ToLower(extension) == "mp4" && (!file.IsDir()) // Ko phải thư mục, đuôi là mp4
}

/*
	Hàm sinh một string có độ dài length
*/
func randomString(length int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())

	str := make([]byte, length)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

/*
Receive channel để xử lý thông báo lỗi
*/
func handle_convert_error(ch <-chan error) {
	go func() {
		for {
			convert_result, more := <-ch        //khi không còn dữ liệu, more sẽ false
			log.Println(convert_result.Error()) // In thông báo của mỗi routine khi có lỗi hoặc hoàn thành

			if !more {
				log.Println("No more result!")
				return //thoát khỏi go routine
			}
		}
	}()

}
