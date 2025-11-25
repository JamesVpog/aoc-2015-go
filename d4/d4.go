package d4

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// content is a secret key, and md5 hash with content and i
// i is from 1..inf
// check if the md5 hash has 5 zeroes
// if it does break and print
// if it doesnt reset the hash and try with new i
func P1(content string) {
	content = strings.Replace(content, "\n", "", 1)

	var hashSum []byte
	var i int

	for true {
		h := md5.New()
		io.WriteString(h, content)
		io.WriteString(h, strconv.Itoa(i))
		hashSum = h.Sum(nil)

		hexHash := fmt.Sprintf("%x", hashSum)
		if strings.HasPrefix(hexHash, "00000") {
			fmt.Printf("%s\n", hexHash)
			fmt.Printf("%d\n", i)
			break
		}
		i += 1
	}

}

// 6 zeroes
func P2(content string) {
	content = strings.Replace(content, "\n", "", 1)

	var hashSum []byte
	var i int

	for true {
		h := md5.New()
		io.WriteString(h, content)
		io.WriteString(h, strconv.Itoa(i))
		hashSum = h.Sum(nil)

		hexHash := fmt.Sprintf("%x", hashSum)
		if strings.HasPrefix(hexHash, "000000") {
			fmt.Printf("%s\n", hexHash)
			fmt.Printf("%d\n", i)
			break
		}
		i += 1
	}
}
