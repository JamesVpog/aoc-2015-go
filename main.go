package main

import (
	"2015-aoc/d1"
	"log"
	"os"
)

func main() {
	inputPath := "d1/1.txt"

	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	content := string(data)
	d1.P1(content)
	d1.P2(content)

}
