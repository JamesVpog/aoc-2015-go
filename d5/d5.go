package d5

import (
	"fmt"
	"strings"
)

func containsAny(s string, substrings []string) bool {
	for _, sub := range substrings {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func hasAtLeast3Vowels(s string) bool {
	// Source - https://stackoverflow.com/a
	// Posted by Eve Freeman
	// Retrieved 2025-11-25, License - CC BY-SA 3.0
	t := 0
	for _, value := range s {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			t++
		}
	}
	return t >= 3
}

func oneLetterAppearsTwiceInARow(s string) bool {
	// check if cur letter is same as next letter
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

//////////////////////////////////////

func P1(content string) {
	// how many nice strings are there?
	// a string is nice if it has:
	// 1. at least 3 vowels
	// 2. one letter that appears twice in a row
	// 3. does not contain substrings "ab", "cd", "pq", or "xy"

	disallowed := []string{"ab", "cd", "pq", "xy"}

	total := 0
	words := strings.Split(content, "\n")
	for _, s := range words {
		if !containsAny(s, disallowed) && hasAtLeast3Vowels(s) && oneLetterAppearsTwiceInARow(s) {
			total += 1
		}
	}
	fmt.Printf("%d\n", total)
}

func has_non_overlapping_pairs(s string) bool {
	// pick a pair of chars [i, i + 2] in the string

	for i := 0; i < len(s)-1; i++ {
		pair := s[i : i+2]
		if strings.Count(s, pair) >= 2 {

			return true
		}
	}
	// find that pair in the next i+2
	return false
}

func has_one_letter_betwee_two_same_chars(s string) bool {
	// check in substrings of 3
	for i := 2; i < len(s); i++ {
		substr := s[i-2 : i+1] //3 characters from with string of i-2, i-1, and i
		if substr[0] == substr[2] {
			return true
		}
	}
	return false
}
func P2(content string) {
	// new nice rules
	// 1. It contains at least one letter which repeats with exactly one letter between them
	// 2. It contains a pair of any two letters that appears at least twice in the string
	// 	  without overlapping

	total := 0
	words := strings.Split(content, "\n")
	for _, w := range words {
		if has_non_overlapping_pairs(w) && has_one_letter_betwee_two_same_chars(w) {
			total += 1
		}
	}
	fmt.Printf("%d\n", total)
}
