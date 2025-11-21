package main

import (
	"2015-aoc/d1"
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
	filePath := "d1/1.txt"
	content := inputReader(filePath)

	d1.P1(content)
	d1.P2(content)

}
