package Solution2

func totalNQueens(n int) int {
	return len(solveNQueens(n))
}

var qstate = [][]int{}
var qpos = []int{}

// var qsize int
// var ans = [][]string{}
var ans int

func totalNQueens2(n int) int {
	ans = 0
	qpos = []int{}
	qstate = [][]int{}
	backtrack(n, qpos)
	return ans
}

// helper: shear new row's state
func shearleft(e1 int) int {
	switch e1 {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 2
	}
	return e1
}

func shearright(e1 int) int {
	switch e1 {
	case 0:
		return 3
	case 1:
		return 2
	case 2:
		return 2
	case 3:
		return 3
	}
	return e1
}

// helper test board
func testboard(n, pn int) bool {
	// state explain 1 left shear
	// 2 both side shear
	// 3 righ shear
	// check -1 row if 3th row
	if len(qstate) == 0 {
		return true
	}
	if len(qstate) >= 2 {
		for col, s := range qstate[len(qstate)-2] {
			switch s {
			case 1:
				if col > 0 {
					qstate[len(qstate)-1][col-1] = shearleft(qstate[len(qstate)-1][col-1])
				}
			case 2:
				if col == 0 {
					qstate[len(qstate)-1][col+1] = shearright(qstate[len(qstate)-1][col+1])
				} else if col > 0 && col < n-1 {
					qstate[len(qstate)-1][col+1] = shearright(qstate[len(qstate)-1][col+1])
					qstate[len(qstate)-1][col-1] = shearleft(qstate[len(qstate)-1][col-1])
				} else if col == n-1 {
					qstate[len(qstate)-1][col-1] = shearleft(qstate[len(qstate)-1][col-1])
				}
			case 3:
				if col < n-1 {
					qstate[len(qstate)-1][col+1] = shearright(qstate[len(qstate)-1][col+1])
				}
			}

		}
	}
	if pn == 0 {
		return qstate[len(qstate)-1][pn+1] == 0 || qstate[len(qstate)-1][pn+1] == 3
	} else if pn == n-1 {
		return qstate[len(qstate)-1][pn-1] == 0 || qstate[len(qstate)-1][pn-1] == 1
	} else {
		return (qstate[len(qstate)-1][pn+1] == 0 || qstate[len(qstate)-1][pn+1] == 3) && (qstate[len(qstate)-1][pn-1] == 0 || qstate[len(qstate)-1][pn-1] == 1)
	}
}

// backtrack standard solution
func backtrack(n int, qpos []int) {
	if len(qpos) == n {
		ans++
		return
	}
	// eliminate duplicated position
	availables := []int{}
	flag := false
	for i := 0; i < n; i++ {
		flag = false
		for _, pn := range qpos {
			if pn == i {
				flag = true
			}
		}
		if !flag {
			availables = append(availables, i)
		}
	}
	// fmt.Printf("availables %v", availables)

	for _, pn := range availables {
		// test compatible
		ok := testboard(n, pn)
		// fmt.Printf("we got %v\n %v is %v\n", qstate, pn, ok)
		if ok {
			row := make([]int, n)
			// 2 means both shear direction
			row[pn] = 2
			qstate = append(qstate, row)
			qpos = append(qpos, pn)
			backtrack(n, qpos)
			qpos = qpos[:len(qpos)-1]
			qstate = qstate[:len(qstate)-1]
		}
	}

}
