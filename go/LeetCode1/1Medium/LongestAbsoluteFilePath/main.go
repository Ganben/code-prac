package main

import (
	"fmt"
	"strings"
)

type Path struct {
	Content string
	Level   int
}

func calLen(stack []Path) int {
	result := 0
	for _, v := range stack {
		result += len(v.Content)
	}
	return result + len(stack) - 1
}

func lengthLongestPath(input string) int {
	s := strings.Split(input, "\n")
	var paths []Path
	for _, v := range s {
		level := strings.Count(v, "\t")
		paths = append(paths, Path{v[level:], level})
	}
	var stack []Path
	var result int
	for _, v := range paths {
		for len(stack) != 0 && stack[len(stack)-1].Level >= v.Level {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, v)
		if strings.Contains(v.Content, ".") {
			result = max(result, calLen(stack))
		}
	}
	return result
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	s := "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext"

	fmt.Printf("%v\n", lengthLongestPath(s))

}
