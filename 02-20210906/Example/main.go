package main

import "fmt"

func main() {
	rawReverseLoop()
}

func rawReverseLoop() {
	cars := [3]string{"Toyota", "Mercedes", "BMW"}

	for i := len(cars) - 1; i >= 0; i-- {
		fmt.Println(i, " ", cars[i])
	}
}
