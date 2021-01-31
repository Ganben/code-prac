package solution

import "fmt"

func isNumber(s string) bool {
	//TODO: preprocess, before space or first space

	if len(s) == 0 {
		return false
	}

	s = spaceRemover(s)

	if len(s) > 0 {
		return isInt(s) || isIntOrFloat(s) || isSci(s)
	}
	return false
}

func spaceRemover(s string) string {
	if len(s) == 1 {
		if s[0] == ' ' {
			return ""
		}
		return s
	}
	for s[0] == ' ' {
		s = s[1:]
	}
	if len(s) > 0 {
		for s[len(s)-1] == ' ' {
			s = s[:len(s)-1]
		}
		return s
	}
	return ""
}

func isInt(s string) bool {
	// 1 -1 -123 or 123
	// space  accepted
	for i, c := range s {
		if i == 0 && (c == '+' || c == '-') {
			// only skip first
			continue
		}

		if !isDigit(c) {
			return false
		}
	}
	if len(s) == 1 && !isDigit(rune(s[0])) {
		return false
	}
	return true
}

func isIntOrFloat(s string) bool {
	// 1 , 1.1 or -1.1
	// space accepted
	dotPos := -1
	numbered := false
	for i, c := range s {
		if i == 0 && (c == '-' || c == '+') {
			continue
		}
		if c == '.' {
			dotPos = i
			break
		}
		if !isDigit(c) {
			fmt.Printf("false 0 %v\n", s)
			return false
		}
		numbered = true
	}
	// check dotPos and
	if dotPos >= 0 && dotPos < len(s)-1 {
		for _, c := range s[dotPos+1:] {
			if !isDigit(c) {
				fmt.Printf("false %v\n", s)
				return false
			}
			numbered = true
		}
	}

	if len(s) == 1 && !isDigit(rune(s[0])) {
		fmt.Printf("false 2 %v\n", s)
		return false
	}
	fmt.Printf("false 4 %v, %v\n", s, numbered)
	return true && numbered

}

func isSci(s string) bool {
	// IntFloat e Int
	// space test required to next funcs
	strippedS := string(s)

	for i, ch := range strippedS {
		if i == 0 && ch == rune(' ') {
			strippedS = s[1:]
			continue
		}
		if ch == rune('e') && i != 0 && i != len(s)-1 {
			// test front isIntOrFloat, rear isInt, not space
			front := strippedS[0:i]
			rear := strippedS[i+1:]
			return isIntOrFloat(front) && isInt(rear)

		}
	}

	return isIntOrFloat(strippedS)

}

func isDigit(c rune) bool {
	digitstrings := "0123456789"
	for _, ch := range digitstrings {
		if ch == c {
			return true
		}
	}
	return false
}
