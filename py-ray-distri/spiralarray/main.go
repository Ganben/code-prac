package main

import "fmt"

var array [][]int

func spiral(n int) [][]int {
	array = [][]int{
		{1, 2},
		{4, 3},
	}
	//
	if n <= 4 {
		array = [][]int{
			{1, 2},
			{4, 3},
		}
		return array
	}
	//

	for i := 4; (i-1)*(i-1) < n; i += 2 {
		row1 := []int{}
		for j := 0; j < i; j++ {
			row1 = append(row1, (i-2)*(i-2)+i-1+j)
		}
		rowe := []int{}
		for j := 0; j < i; j++ {
			rowe = append(rowe, i*i-j)
		}
		//
		newArr := [][]int{}
		for j := 0; j < i; j++ {
			if j == 0 {
				newArr = append(newArr, row1)
				continue
			} else if j == i-1 {
				newArr = append(newArr, rowe)
				break
			} else {
				//
				row := []int{}
				row = append(row, newArr[0][0]-j)
				row = append(row, array[j-1]...)
				row = append(row, newArr[0][i-2]+j+1)
				newArr = append(newArr, row)
			}
		}
		//
		array = newArr
		// fmt.Printf("%v\n", array)
	}
	return array
}

func main() {
	arr := spiral(400)
	for _, v := range arr {
		for _, nn := range v {
			fmt.Printf("%5d", nn)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("%v,", arr[14][7])
	ans := 0
	for i := 0; i < 20; i++ {
		ans += arr[i][1]
	}
	fmt.Printf("sum: %v", ans)
}
