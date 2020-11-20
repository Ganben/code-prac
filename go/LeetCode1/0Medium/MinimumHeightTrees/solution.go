package MinimumHeightTrees

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	inG := make([]int, n)
	outG := make(map[int][]int)
	for i := 0; i < len(edges); i++ {
		outG[edges[i][0]] = append(outG[edges[i][0]], edges[i][1])
		inG[edges[i][1]]++
		outG[edges[i][1]] = append(outG[edges[i][1]], edges[i][0])
		inG[edges[i][0]]++
	}

	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if inG[i] == 1 {
			queue = append(queue, i)
		}
	}

	var rst []int
	for len(queue) > 0 {
		rst = make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			rst = append(rst, top)
			for _, v := range outG[top] {
				inG[v]--
				if inG[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
	}
	return rst
}
