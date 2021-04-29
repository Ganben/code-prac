package main

import (
	"fmt"
	"strconv"
)

func compress(chars []byte) int {
	anchor, write := 0, 0
	for i, c := range chars {
		if i+1 == len(chars) || chars[i+1] != c {
			chars[write] = chars[anchor]
			write += 1
			if i > anchor {
				for _, digit := range strconv.Itoa(i - anchor + 1) {
					chars[write] = byte(digit)
					write += 1
				}
			}
			anchor = i + 1
		}
	}
	return write
}

func main() {
	// i := []byte{b"a", b"a", b"b", b"b", b"c", b"c", b"c"}
	i := "aabbccc"
	fmt.Printf("%v\n", compress([]byte(i)))
}
