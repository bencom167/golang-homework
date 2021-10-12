package main

import (
	"fmt"
	"strings"
)

func main() {
	inputFile := "Users/bendn/code/golang/06-20210927/homework/input.abc.json"

	dotPos := strings.LastIndex(inputFile, ".")
	slashPos := strings.LastIndex(inputFile, "/")

	dir := inputFile[:slashPos+1]
	fileName := inputFile[slashPos+1 : dotPos]
	extension := inputFile[dotPos+1:]

	fmt.Println(dir)
	fmt.Println(fileName)
	fmt.Println(extension)

}
