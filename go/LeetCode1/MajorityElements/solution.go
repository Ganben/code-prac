package MajorityElements

func majorityElement(nums []int) int {
	nmap := countNum(nums)
	maxN := 0
	res := -1
	for k, v := range nmap {
		if v > maxN {
			maxN = v
			res = k
		}
	}
	return res
}

func countNum(nums []int) map[int]int {
	counts := map[int]int{}
	for _, v := range nums {
		if _, ok := counts[v]; ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}
	return counts
}
