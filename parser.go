package main

import (
	"bufio"
	"log"
	"os"
)

func GetArgs() (filename string) {
	if len(os.Args) == 2 {
		filename = os.Args[1]
	} else {
		log.Fatal("Error: no filename specified")
	}
	return filename
}

func ParseFile(filename string) (data []string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) != 0 {
			data = append(data, scanner.Text())
		}
	}
	return data
}
