package solution

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	times := n1
	p1 := 0
	s2Times := 0
	s1Times := 0
Loop:
	for i := 0; i < times; i++ {
		// try i cycles for string matching

		for id1, ch := range s1 {
			if ch == rune(s2[p1]) && p1 == len(s2)-1 {
				//reach end
				s2Times++
				// fmt.Printf("matched by %v at %v of %v\n", s2Times, id1, i+1)
				p1 = 0

			} else if ch == rune(s2[p1]) && p1 < len(s2)-1 {
				// fmt.Printf("char match at %v--\n", p1)
				p1++

			}
			if id1 == len(s1)-1 && p1 == 0 {
				s1Times = i + 1
				break Loop
			}

		}
		s1Times = i + 1
	}
	if s2Times == 0 {
		return 0
	} else {
		return int(n1 * s2Times / s1Times / n2)
	}
}
