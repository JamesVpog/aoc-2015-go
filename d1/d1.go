package d1

import "fmt"

// prints the answer to AOC day 1 part 1
func P1(content string) {

	total := 0

	for _, d := range content {
		switch d {
		case '(':
			total += 1
		case ')':
			total -= 1
		}
	}

	fmt.Printf("%d\n", total) // 74
}

// prints the answer to AOC day 1 part 2
func P2(content string) {

	total := 0

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
}
