package diagonaltrav

func findDiagonalOrder(mat [][]int) []int {
	flag := 1
	// sum := 0
	rN, cN := 0, 0
	res := []int{}
	for i := 0; i < len(mat)*len(mat[0]); i++ {
		//
		if rN > len(mat)-1 || cN > len(mat[0])-1 {
			break
		}
		res = append(res, mat[rN][cN])
		// inc: flag / edge test; 4 state;
		if flag == 1 && rN == 0 && cN < len(mat[0])-1 {
			rN += 0
			cN += 1
			flag *= -1
		} else if flag == -1 && cN == 0 && rN < len(mat)-1 {
			rN += 1
			cN += 0
			flag *= -1
		} else if flag == 1 && cN == len(mat[0])-1 {
			rN += 1
			cN += 0
			flag *= -1

		} else if flag == -1 && rN == len(mat)-1 {
			rN += 0
			cN += 1
			flag *= -1
		} else {
			if flag == 1 {
				rN += -1
				cN += 1
			} else {
				rN += 1
				cN += -1
			}
		}

		if rN == len(mat)-1 && cN == len(mat[0])-1 {
			res = append(res, mat[rN][cN])
			break
		}
	}
	return res
}
