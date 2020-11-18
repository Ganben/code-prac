package RemoveInvalidParentheses

func removeInvalidParenthese(s string) []string {
	var (
		res     []string
		isValid func(s string) bool
	)
	isValid = func(s string) bool {
		left := 0
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				left++
			} else if s[i] == ')' {
				left--
			}
			if left < 0 {
				return false
			}
		}
		return left == 0
	}

	level := map[string]bool{s: true}
	for len(level) > 0 {
		
	for v, _ := range level {
		if isValid(v) {
			res = append(res, v)
		}
		if len(res) > 0 {
			break
		}
		nextLevel := map[string]bool{}
		for v, _ := range level {
		for i:= 0; i <len(v); i++ {
			if v[i] == '(' || v[i] == ')' {
				new := []byte{}
				new = append(new, []byte(v)[:i]...)
				new = append(new, []byte(v)[i+1:]...)
				nextLevel[string(new)] = true
			}
		}
	}
	level = nextLevel
}
return res
}


var ansMap map[string]bool

func removeInvalidParentheses(s string) []string {
	if s == "" {
		return []string{""}
	}
	ansMap = make(map[string]bool)
	ans := make([]string, 0)
	lcnt, rcnt := GetRemoveCnt(s)
	helper1("", s, 0, lcnt, rcnt)
	for k, _ := range ansMap {
		ans = append(ans, k)
	}
	return ans

}

func helper1(pre, s string, stackcnt, lcnt, rcnt int) {
	if lcnt == 0 && rcnt == 0 {
		isValid(pre, s, stackcnt)
		return
	}
	if lcnt < 0 || rcnt < 0 || s == "" {
		return
	}

	if s[0] != '(' && s[0] != ')' {
		pre += s[:1]
		helper1(pre, s[1:], stackcnt, lcnt, rcnt)
	}
	if s[0] == ')' {
		if stackcnt == 0 { //delete
			helper1(pre, s[1:], stackcnt, lcnt, rcnt-1)
		} else {
			helper1(pre, s[1:], stackcnt, lcnt, rcnt-1) //delete
			pre += s[:1]
			helper1(pre, s[1:], stackcnt-1, lcnt, rcnt) //use
			pre = pre[:len(pre)-1]
		}
	}
	if s[0] == '(' {
		helper1(pre, s[1:], stackcnt, lcnt-1, rcnt) //delete
		pre += s[:1]
		helper1(pre, s[1:], stackcnt+1, lcnt, rcnt) //use
		pre = pre[:len(pre)-1]
	}

}

func isValid(pre, s string, stackcnt int) {
	for i := 0; i < len(s); i++ {
		if s[i] != '(' && s[i] != ')' {
			continue
		}
		if s[i] == ')' {
			if stackcnt == 0 {
				return
			}
			stackcnt--
		}
		if s[i] == '(' {
			stackcnt++
		}
	}
	if stackcnt == 0 {
		ansMap[string(pre)+s] = true
	}
}

func GetRemoveCnt(s string) (int, int) {
	n := len(s)
	lcnt, rcnt := 0, 0
	for i := 0; i < n; i++ {
		if s[i] != '(' && s[i] != ')' {
			continue
		}
		if s[i] == '(' {
			lcnt++
		}
		if s[i] == ')' {
			if lcnt == 0 {
				rcnt++
			} else {
				lcnt--
			}
		}
	}
	return lcnt, rcnt
}

作者：changwenlong
链接：https://leetcode-cn.com/problems/remove-invalid-parentheses/solution/dfshui-su-by-changwenlong/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。