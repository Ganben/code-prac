package solution

import "fmt"

var pubgrid [][]byte

// func numIslands(grid [][]byte) int {
// 	result := 0

// 	return result
// }

func numIslands(grid [][]byte) int {
	result := 0
	pubgrid = [][]byte{}
	for i := 0; i < len(grid); i++ {
		pubgrid = append(pubgrid, grid[i])
	}
	// copy a pub grid

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if pubgrid[i][j] == '1' {
				fmt.Printf("found 1 %v \n", pubgrid)
				pubgrid[i][j] = byte(2)
				searchAndReplace(i, j)
				result++
				fmt.Printf("del %v \n", pubgrid)
			}
		}
	}
	return result

}

func searchAndReplace(x, y int) {
	//d1
	if x > 0 {
		if pubgrid[x-1][y] == '1' {
			// assign byte(2) and search deeper
			pubgrid[x-1][y] = byte(2)
			searchAndReplace(x-1, y)
		}
	}

	if y > 0 {
		if pubgrid[x][y-1] == '1' {
			// assign and search deeper
			pubgrid[x][y-1] = byte(2)
			searchAndReplace(x, y-1)
		}
	}

	if x < len(pubgrid)-1 {
		if pubgrid[x+1][y] == '1' {
			pubgrid[x+1][y] = byte(2)
			searchAndReplace(x+1, y)
		}
	}

	if y < len(pubgrid[0])-1 {
		if pubgrid[x][y+1] == '1' {
			pubgrid[x][y+1] = byte(2)
			searchAndReplace(x, y+1)
		}
	}

}

func nearbyTest(ar []byte, val, newVal byte) byte {
	for _, e := range ar {
		if e == newVal {
			if val == '1' {
				val = e
			}
		}
	}
	// return 0 or newVal
	return val
}
