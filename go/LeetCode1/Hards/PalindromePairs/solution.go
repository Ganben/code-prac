package PalindromePairs

type Trie struct {
	children [60]*Trie
	// char is it's location derived from parent position
	val    rune
	id     int
	isWord bool
}

func (t *Trie) Insert(word string, id int) {
	cur := t
	add := 0
	for k, v := range word {
		c := v - 'A'

		if cur.children[c] == nil {
			cur.children[c] = &Trie{}
			cur.children[c].val = v
			cur.children[c].id = id
			if k == len(word)-1 {
				cur.isWord = true
			}
		}
	}
}

func (t *Trie) Query(word string) []int {
	cur := t
	ret := []int{}
	for _, v := range word {
		c := v - 'A'
		ret = append(ret, cur.id)
		if cur.children[c] == nil {
			return ret
		}

		cur = cur.children[c]
	}
	if cur.isWord == true {
		ret = append(ret, cur.id)
	} else {
		ret = append(ret, -1)
		// -1 is for identify no such word?? weird for algo transplant
	}
	return ret
}

func palindromePairs(words []string) [][]int {
	t1, t2 := &Trie{}, &Trie{}
	n := len(words)
	for i := 0; i < n; i++ {
		t1.Insert(words[i], i)
		tmp := words[i]
		tmp = reverse(tmp)
		t2.Insert(tmp, i)
	}

	ret := [][]int{}
	for i := 0; i < n; i++ {
		rec := manacher(words[i])
		id1 := t2.Query(words[i])
		words[i] = reverse(words[i])
		id2 := t1.Query(words[i])
		m := len(words[i])
		allId := id1[m]
		if allId != -1 && allId != i {
			ret = append(ret, []int{i, allId})
		}
		for j := 0; j < m; j++ {
			if rec[j][0] == 1 {
				leftId := id2[m-j-1]
				if leftId != -1 && leftId != i {
					ret = append(ret, []int{leftId, i})
				}
			}
			if rec[j][1] == 1 {
				rightId := id1[j]
				if rightId != -1 && rightId != i {
					ret = append(ret, []int{i, rightId})
				}
			}
		}

	}
	return ret

}

func manacher(s string) [][]int {
	n := len(s)
	tmp := "#"
	tmp += string(s[0])
	for i := 1; i < n; i++ {
		tmp += "*"
		tmp += string(s[i])
	}
	tmp += "!"
	m := n * 2
	l := []int{}
	r := [][]int{}
	for i := 0; i < m; i++ {
		l = append(l, 0)
	}
	for i := 0; i < n; i++ {
		r = append(r, []int{0, 0})
	}

	p, maxn := 0, -1
	for i := 1; i < m; i++ {

		if maxn >= 1 {
			l[i] = min(l[2*p-i], maxn-i)
		} else {
			l[i] = 0
		}

		for tmp[i-l[i]-1] == tmp[i+l[i]+1] {
			l[i]++
		}
		if i+l[i] > maxn {
			p = i
			maxn = i + l[i]
		}
		if i-l[i] == 1 {
			r[(i+l[i])/2][0] = 1
		}
		if i+l[i] == m-1 {
			r[(i - l[i]/2)][1] = 1
		}

	}
	return r
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func reverse(s string) string {
	n := len(s)
	b := []byte(s)
	for i := 0; i < n/2; i++ {
		b[i], b[n-i-1] = b[n-i-1], b[i]
	}
	return string(b)
}
