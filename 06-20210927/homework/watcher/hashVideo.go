package watcher

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"sync"

	"github.com/TechMaster/eris"
)

type VideoInfo struct {
	Duration   int    // Thời lượng in seconds
	Resolution string // Độ phân giải
}

/*
Hàm này làm 3 việc:
1. Tạo folder HLS để chưa kết quả băm
2. Chuẩn bị exec.Command để gọi lệnh ffmpeg
3. Gọi hàm executeHashing trong để băm
*/
func hashVideoToHLS(videoID string, videoName string, wg *sync.WaitGroup, ch chan<- error) {
	defer wg.Done() //defer để chạy hàm này vào lúc cuối trước khi hàm thoát

	args, hashedFolder, err := makeHlsDir(videoID, videoName)
	if err != nil {
		ch <- err
	}

	cmd := exec.Command("ffmpeg", args...)
	cmd.Dir = hashedFolder

	// Bắt đầu băm video
	err = executeHashing(cmd)
	if err == nil {
		hashedVideoWriter(videoID, videoName)
		err = errors.New("Hoàn thành convert: " + videoID + " - " + videoName)
	}
	ch <- err // Đẩy thông báo vào channel
}

/*
Hàm này thực sự tiến hành băm
*/
func executeHashing(cmd *exec.Cmd) error { // Đã bỏ tham số  uploadMeta UploadVideoMeta
	err := cmd.Start()
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi khi bắt đầu chạy lênh băm ffmpeg!")
	}

	//Đoạn này chạy rất lâu !
	err = cmd.Wait()
	if err != nil {
		return eris.NewFromMsg(err, "Lỗi trong quá trình chạy lênh băm ffmpeg!")
	}
	return nil
}

/*
Tạo thư mục HLS chờ sẵn để băm
*/
func makeHlsDir(videoID string, videoName string) (args []string, hashedFolder string, err error) {
	/*Tạo thư mục chứa file m3u8 và các file sau khi băm video
	Lệnh băm ffmpeg sẽ chạy ở thư mục này do đó chỉ cần truyền tên file manifestFile m3u8 là đủ
	*/
	hashedFolder = path.Join(HLSDir, videoID)
	if err := os.MkdirAll(hashedFolder, 0777); err != nil {
		return nil, hashedFolder, eris.NewFromMsg(err, "Lỗi tạo thư mục HLS "+videoID)
	}

	manifiestFile := videoID + ".m3u8"

	videoPath := path.Join(UploadDir, videoName)

	//Lấy thông tin độ phân giải
	videoInfo, err := GetVideoInfo(videoPath)
	if err != nil {
		return nil, hashedFolder, err
	}
	resolution := strings.Split(videoInfo.Resolution, "x")
	scale := fmt.Sprintf("scale=w=%s:h=%s:force_original_aspect_ratio=decrease", resolution[0], resolution[1])

	hlsKey := randomString(10)
	hlsVector := randomString(10)

	return []string{
		"-i", videoPath,
		"-vf", scale, "-c:v", "h264", "-c:a", "aac", "-ar", "48000", "-profile:v", "main", "-crf", "20",
		"-r", "30", "-g", "60", "-maxrate", "5350k", "-bufsize", "7500k", "-b:a", "192k",
		"-hls_enc", "1", "-hls_enc_key", hlsKey, "-hls_enc_iv", hlsVector,
		"-hls_playlist_type", "vod", manifiestFile,
		"-hide_banner", "-loglevel", "quiet", "-progress", "/dev/stdout",
	}, hashedFolder, nil
}

// Truyền vào đường dẫn file video trả về thời lượng và độ phân giải
func GetVideoInfo(filePath string) (videoInfo VideoInfo, err error) {
	var result []byte
	durationArgs := []string{
		"-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	}
	result, err = exec.Command("ffprobe", durationArgs...).Output()
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi chạy ffprobe").InternalServerError()
	}
	durationStr := string(result)
	indexOfDot := strings.Index(durationStr, ".") //Chỉ lấy phần nguyên, bỏ phần thập phân
	if indexOfDot != -1 {
		durationStr = durationStr[:indexOfDot]
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi đọc thời lượng video").InternalServerError()
	}
	videoInfo.Duration = duration

	// Đọc độ phân giải
	resolutionArgs := []string{
		"-v", "error", "-show_entries", "stream=width,height",
		"-of", "csv=p=0:s=x",
		filePath,
	}
	result, err = exec.Command("ffprobe", resolutionArgs...).Output()
	if err != nil {
		return videoInfo, eris.NewFromMsg(err, "Lỗi đọc độ phân giải video").InternalServerError()
	}

	resolution := strings.ReplaceAll(string(result), "\n", "")
	videoInfo.Resolution = resolution

	return videoInfo, nil
}
