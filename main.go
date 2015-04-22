package main

import (
	"fmt"
)

func main() {
	var filename string
	var data []string
	filename = GetArgs()
	data = ParseFile(filename)
	for i := range data {
		fmt.Println(data[i])
	}
}
