package RepeatedDNA

import "math"

func findRepeatedDnaSequences(s string) []string {
	L := 10
	n := len(s)
	a := 4
	aL := int(math.Pow(float64(a), float64(L)))
	c2Int := map[byte]int{}
	c2Int['A'] = 0
	c2Int['C'] = 1
	c2Int['G'] = 2
	c2Int['T'] = 3
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = c2Int[s[i]]
	}
	h := 0
	seen := map[int]int{}
	output := map[string]int{}
	for start := 0; start < n-L+1; start++ {
		if start != 0 {
			h = h*a - nums[start-1]*aL + nums[start+L-1]
		} else {
			for i := 0; i < L; i++ {
				h = h*a + nums[i]
			}
		}

		if _, ok := seen[h]; ok {
			output[s[start:start+L]] = 1
		}
		seen[h] = 1
	}
	ans := []string{}
	for k, _ := range output {
		ans = append(ans, k)
	}
	return ans
}
