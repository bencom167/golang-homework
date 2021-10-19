package watcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TechMaster/eris"
	"github.com/spf13/viper"
)

func DemoJSONReader() {
	content, err := ioutil.ReadFile("/Users/bendn/code/golang/06-20210927/homework/hls/result.json") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content)) // This is some content

	var hashedVideos []VideoNameID

	err = json.Unmarshal([]byte(content), &hashedVideos)
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = eris.NewFromMsg(err, "Lỗi không tìm thấy file config!")
		} else {
			err = eris.NewFromMsg(err, "Lỗi đọc file config!")
		}
	}

	fmt.Println(hashedVideos)
}
