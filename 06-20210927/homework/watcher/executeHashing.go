package watcher

import (
	"io/fs"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/TechMaster/eris"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

var (
	UploadDir    string
	HLSDir       string
	HashedResult string
)

type VideoNameID struct {
	ID   string // ID video
	Name string // Tên file video
}

/*
	Hàm băm toàn bộ file mp4 có trong thư mục upload
*/
func HashAllVideo(inputFile string) error {
	videos := prepareInputs(inputFile) // Lấy danh sách video cần băm

	convert_result_channel := make(chan error)
	handle_convert_error(convert_result_channel)

	wait_group := sync.WaitGroup{}

	// Băm video theo danh sách
	for _, video := range videos {
		log.Printf("Bắt đầu convert: %v <- %v\n", video.ID, video.Name)
		wait_group.Add(1)
		go hashVideoToHLS(video.ID, video.Name, &wait_group, convert_result_channel)
	}

	// Bổ sung code watcher, theo dõi file mp4 bổ sung thêm và băm
	// Code tiếp
	hashedVideoWatcher(&wait_group, convert_result_channel)

	wait_group.Wait() // Đợi các routine hoàn thành hết

	log.Printf("Hoàn thành check/convert %v MP4 videos!\n", len(videos))
	return nil
}

/*
	Hàm đọc dữ liệu đầu vào
	1. Lấy thông tin thư mục input/output
	2. Lấy danh sách file trong thư mục input upload
	3. Loại bỏ những file sai định dạng và những file mp4 đã băm rồi
*/
func prepareInputs(inputFile string) []VideoNameID {
	// Đọc lấy thông tin thư mục input/output, nếu có lỗi thì in thông báo
	if err := inputReader(inputFile); err != nil {
		log.Fatal(err)
	}

	hashedVideos, _ := hashedVideoReader()
	var unhashedVideos []VideoNameID

	files, err := ioutil.ReadDir(UploadDir) // Đọc thư mục, trả lại danh sách File Infor
	if err != nil {                         // Nếu có lỗi thì in thông báo lỗi
		log.Fatal(err)
	}
	for _, file := range files {
		if isNotHashedVideo(file.Name(), hashedVideos) && isValidMP4(file) {
			unhashedVideos = append(unhashedVideos, VideoNameID{ID: randomString(8), Name: file.Name()})
		}
	}
	return unhashedVideos
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
	Hàm đọc danh sách file mp4 đã băm
*/
func hashedVideoReader() ([]VideoNameID, error) {
	var hashedVideos []VideoNameID

	content, err := ioutil.ReadFile(HashedResult)
	if err != nil {
		return hashedVideos, eris.NewFromMsg(err, "Lỗi đọc file config!")
	}

	data := make(map[string]VideoNameID)
	err = yaml.Unmarshal(content, &data)
	if err != nil {
		err = eris.NewFromMsg(err, "Lỗi đọc file config format yaml!")
	}

	for _, v := range data {
		hashedVideos = append(hashedVideos, v)
	}

	return hashedVideos, err
}

/*
	Hàm đọc danh sách file mp4 đã băm
*/
func hashedVideoWriter(videoID string, videoName string) error {
	hashedVideo := map[string]VideoNameID{"video " + videoID: {videoID, videoName}}

	data, err := yaml.Marshal(&hashedVideo)
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi định dạng dữ liệu sang format yaml!")
	}

	file, err := os.OpenFile(HashedResult, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi mở/tạo file ghi kết quả băm!")
	}
	defer file.Close()
	if _, err := file.Write(data); err != nil {
		return eris.NewFromMsg(err, "Lỗi ghi kết quả băm!")
	}

	return nil
}

/*
	Hàm theo dõi thưc mục upload, Nếu có file mp4 bổ sung thì băm thêm
*/
func hashedVideoWatcher(wg *sync.WaitGroup, convert_result_channel chan<- error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					convert_result_channel <- eris.NewFromMsg(nil, "Lỗi watcher!")
					return
				}
				if event.Op&fsnotify.Create == fsnotify.Create { // Có file mới đc upload
					myVideo, _ := os.Stat(event.Name)
					if isValidMP4(myVideo) {
						videoID := randomString(8)
						log.Printf("File mới upload, bắt đầu convert: %v <- %v\n", videoID, myVideo.Name())
						wg.Add(1)
						go hashVideoToHLS(videoID, myVideo.Name(), wg, convert_result_channel)
					}
				}
			case err, ok := <-watcher.Errors:
				convert_result_channel <- eris.NewFromMsg(err, "Lỗi watcher!")
				if !ok {
					return
				}
			}
		}
	}()

	err = watcher.Add(UploadDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

/*
	Hàm đọc kiểm tra xem có phải mp4 hay không
*/
func isValidMP4(file fs.FileInfo) bool {
	extension := file.Name()[strings.LastIndex(file.Name(), ".")+1:]
	return strings.ToLower(extension) == "mp4" && (!file.IsDir()) // Ko phải thư mục, đuôi là mp4
}

/*
	Hàm kiểm tra xem file mp4 đã băm chưa
*/
func isNotHashedVideo(videoName string, hashedVideos []VideoNameID) bool {
	for _, video := range hashedVideos {
		if video.Name == videoName {
			return false
		}
	}
	return true
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

func DemoReader(inputFile string) error {
	videos := prepareInputs(inputFile) // Lấy danh sách video cần băm
	log.Println(videos)
	return nil
}

func DemoWriter() error {
	hashedVideoWriter("cccccccc", "test.mp4")
	return nil
}
