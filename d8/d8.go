package d8

import (
	"fmt"
	"strings"
)

// per line
// number of chars of code for string literals - number of chars in memory for values of the strings in entire file
func P1(content string) {
	total := 0
	lines := strings.Split(content, "\n")
	for _, l := range lines[:len(lines)-1] {
		codeChars := len(l)
		inMemChars := getNumOfCharsForInMemString(l)
		res := codeChars - inMemChars
		total += res
	}
	fmt.Println(total)
}

func getNumOfCharsForInMemString(s string) int {

	// if len(s) is 2, its just the ""
	if len(s) == 2 {
		return 0
	}

	// ignore " at start and end
	res := s[1 : len(s)-1]

	count := 0

	// parse looking for double-quoted string literals
	// " \\ ", " \" " , and "\x[00-FF]"
	for i := 0; i < len(res); i++ {
		if res[i] == '\\' { // we find a forward slash,
			// make sure not to index beyond the string
			if i+1 < len(res) {
				next := res[i+1]
				if next == '\\' || next == '"' {
					// find another backslash or "
					count += 1
					i++
				} else if next == 'x' && i+3 < len(res) {
					count += 1
					i += 3
				}
			}
		} else {
			count += 1
		}
	}
	return count
}

// add / behind " and /
func getNumCharNewlyEncodedString(s string) int {
	new_s := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			new_s += "\\" + string('"')
		} else if s[i] == '\\' {
			new_s += "\\\\"
		} else {
			new_s += string(s[i])
		}
	}
	new_s = string('"') + new_s + string('"')
	fmt.Println(new_s)
	return len(new_s)
}

// add " to start and end
func P2(content string) {
	lines := strings.Split(content, "\n")
	total := 0
	for _, l := range lines[:len(lines)-1] {
		numCharCode := len(l)
		numCharNewlyEnc := getNumCharNewlyEncodedString(l)
		res := numCharNewlyEnc - numCharCode
		total += res
	}
	fmt.Println(total)
}
