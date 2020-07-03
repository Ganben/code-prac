package PalindromePartition

var res [][]string
var l int
var dp [][]bool
var ss string

func partition(s string) [][]string {
	n := len(s)
	ss = s
	l = n
	dp = [][]bool{}
	for i := 0; i < n; i++ {
		row := make([]bool, n)
		dp = append(dp, row)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i+1; j++ {
			if s[i] == s[j] && (i-j <= 2 || dp[j+1][i-1]) {
				dp[j][i] = true
			}
		}
	}
	res = [][]string{}
	return res
}

func helper(i int, tmp []string) {
	if i == l {
		res = append(res, tmp)
	}
	for j := i; j < l; j++ {
		if dp[i][j] {
			helper(j+1, append(tmp, ss[i:j+1]))
		}

	}
}
