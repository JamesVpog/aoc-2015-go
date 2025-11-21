package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputPath := "inputs/1.txt"

	data, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	// process the data, which is a slice of bytes
	total := 0

	content := string(data)
	for i, d := range content {
		b := false
		switch d {
		case '(':
			total += 1
		case ')':
			total -= 1
			if total < 0 {
				fmt.Printf("%d\n", i+1)
				b = true
			}
		}

		if b {
			break
		}
	}

	fmt.Printf("%d\n", total) // 74

}
