package Solution

func rotate(matrix [][]int) {
	if len(matrix) != len(matrix[0]) {
		return
	}
	// odd c = 0, even c = -1
	c := len(matrix)%2 - 1
	// center idx * 2
	q := len(matrix) - 1

	for i := 0; i < int(len(matrix)/2); i++ {
		for j := 0; j < int(len(matrix)/2); j++ {
			// 1/4 test
			//r 1
			//ri := q - i
			//rj := q - j
			// tmp := matrix[i][q-j]
			// r1
			matrix[j][q-i], matrix[q-i][q-j], matrix[q-j][i], matrix[i][j] = matrix[i][j], matrix[j][q-i], matrix[q-i][q-j], matrix[q-j][i]
		}
	}
	if c == 0 {
		// rotate row center
		e := int(q / 2)
		for i := 0; i < e; i++ {
			// r1
			matrix[e][q-i], matrix[q-i][e], matrix[e][i], matrix[i][e] = matrix[i][e], matrix[e][q-i], matrix[q-i][e], matrix[e][i]
			// r2
			// matrix[e][q-i], matrix[q-i][q-i] = matrix[q-i][q-i], matrix[e][q-i]
			// // r3
			// matrix[q-i][q-i], matrix[e][i] = matrix[e][i], matrix[q-i][q-i]
			// // r4
			// matrix[e][i], matrix[i][e] = matrix[i][e], matrix[e][i]
		}

	}
}
