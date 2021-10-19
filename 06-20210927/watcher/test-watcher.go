package watcher

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

// main
func DemoWatcher() {

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
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Them file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/Users/bendn/code/golang/06-20210927/homework/upload")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
