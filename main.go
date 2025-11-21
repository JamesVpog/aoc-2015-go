package main

import (
	"2015-aoc/d2"
	"log"
	"os"
)

func inputReader(filePath string) (content string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content = string(data)
	return
}

func main() {
	filePath := "d2/2.txt"
	content := inputReader(filePath)

	d2.P1(content)
	d2.P2(content)
}
