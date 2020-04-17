package solution

import "fmt"

var ans [][]int

func generateMatrix(n int) [][]int {
	ans = [][]int{}
	for i := 0; i < n; i++ {
		row := make([]int, n)
		ans = append(ans, row)
	}
	walker(n, 0, 0)
	return ans
}

func walker(length, pos, start int) {
	if length <= 0 {
		return
	}
	if length == 1 {
		start++
		ans[pos][pos] = start
		return
	}
	d := 0
	a, b := pos, pos
Loop:
	for i := 0; i < 4*length-4; i++ {
		start++
		ans[a][b] = start
		fmt.Printf("assign %v,%v for %v\n", a, b, start)
		switch d {
		case 0:
			if b+1 > length+pos-1 {
				d = 1
				a++
				continue
			}
			b++
		case 1:
			if a+1 > length+pos-1 {
				d = 2
				b--
				continue
			}
			a++
		case 2:
			if b-1 < pos {
				d = 3
				a--
				continue
			}
			b--
		case 3:
			if a-1 == pos {
				break Loop
			}
			a--
		}
	}
	fmt.Printf("visit deeper %v,%v,%v\n\n", length-2, pos+1, start)
	walker(length-2, pos+1, start)
}
