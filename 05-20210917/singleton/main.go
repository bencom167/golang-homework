package main

import (
	"fmt"
	"singleton-pattern/singleton"
)

func main() {
	/*
		ins1 := singleton.GetInstanceFunc()
		ins2 := singleton.GetInstanceFunc()
		fmt.Printf("%p\n", ins1)
		fmt.Printf("%p\n", ins2)
	*/
	ins3 := singleton.GetInstanceSync()
	ins4 := singleton.GetInstanceSync()
	fmt.Printf("%p\n", ins3)
	fmt.Printf("%p\n", ins4)
	/*
		ins5 := singleton.GetInstanceMutex()
		ins6 := singleton.GetInstanceMutex()
		fmt.Printf("%p\n", ins5)
		fmt.Printf("%p\n", ins6)
	*/
	//singleton.DemoGetInstanceFunc()
	//singleton.DemoGetInstanceCheck()
	//singleton.DemoGetInstanceMutex()
	//time.Sleep(time.Second * 5)

	//config.DemoConfig()
	// database.DemoDatabase()
	//post.DemoPost()
}
