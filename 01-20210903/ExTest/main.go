package main

import (
	// Thu vien chuan

	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Khoi tao bien sinh ngau nhien.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Sinh ngau nhien so nguyên dương X >= 0 và <= 100
	for i := 0; i < 100; i++ {
		fmt.Println(r.Int63n(10))
	}
}
