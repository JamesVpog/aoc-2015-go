package main

import (
	"2015-aoc/d3"
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
	filePath := "d3/3.txt"
	content := inputReader(filePath)

	d3.P1(content)
	d3.P2(content)
}
