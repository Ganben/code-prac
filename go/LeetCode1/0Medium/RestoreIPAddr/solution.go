package solution

import "strconv"

var n int
var segments []string
var output []string
var input string

func restoreIpAddresses(s string) []string {
	// backtrack with 3 dots position
	n = len(s)
	input = s
	segments = []string{}
	output = []string{}
	backtrack(-1, 3)
	return output
}

func backtrack(prev, dots int) {
	max_pos := min(n-1, prev+4)
	for curr := prev + 1; curr <= max_pos; curr++ {
		segmentString := input[prev+1 : curr+1]
		if valid(segmentString) {
			segments = append(segments, segmentString)
			if dots == 0 {
				if curr < n-1 {
					segments = segments[:len(segments)-1]
					continue
				}
				o := ""
				for i, e := range segments {
					o += e
					if i < len(segments)-1 {
						o += "."
					}
				}
				output = append(output, o)

			} else {
				backtrack(curr, dots-1)
			}
			segments = segments[:len(segments)-1]
		}

	}
}

func valid(s string) bool {
	m := len(s)
	if m > 3 {
		return false
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if v > 255 {
		return false
	}
	if s[0] == '0' && m > 1 {
		return false
	}
	if v == 0 && m > 1 {
		return false
	}
	return true
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
