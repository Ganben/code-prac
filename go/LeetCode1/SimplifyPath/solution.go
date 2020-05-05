package SimplifyPath

import "strings"

func simplifyPath(path string) string {
	// stack
	sArr := []string{}
	splited := strings.Split(path, "/")
	for _, c := range splited {
		if c == "." {
			// donothing
		} else if c == ".." {
			if len(sArr) >= 1 {
				sArr = sArr[:len(sArr)-1]
			}

		} else if c != "" {
			sArr = append(sArr, c)
		}
	}
	ans := "/"
	for i, cn := range sArr {
		ans += cn
		if i < len(sArr)-1 {
			ans += "/"
		}
	}
	return ans

}
