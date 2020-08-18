package DesignWordDictionary

// isn't this a trie?

type WordDictionary struct {
	links  []*WordDictionary
	isWord bool
	value  string
}

func Constructor() WordDictionary {
	d := WordDictionary{}
	d.links = make([]*WordDictionary, 60)
	d.isWord = false
	return d

}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}
	if !this.Search(word) {
		cur := this
		for i, v := range word {
			j := int(v - 'a')
			if cur.links[j] == nil {
				p := Constructor()
				cur.links[j] = &p
				// cur.links[j].val = byte(v)
				if i == len(word)-1 {
					cur.isWord = true
					cur.value = word
				}
				// cur = cur.links[j]
			} else {
				if i == len(word)-1 {
					cur.isWord = true
					cur.value = word
				}
				// cur = cur.links[j]
			}
			cur = cur.links[j]
		}
	}
}

func (this *WordDictionary) Search(word string) bool {
	cur := this
	if len(word) == 0 {
		return false
	}
	b := RecSearch(cur, word)
	return b
}

func RecSearch(d *WordDictionary, subs string) bool {
	j := int(subs[0] - 'a')

	if len(subs) == 1 {
		if d.isWord && subs[0] == '.' {
			return true
		}

		if d.isWord && d.links[j] != nil {
			return true
		}
		return false
	}
	if j == -51 || j == 205 {
		possible := []*WordDictionary{}
		for _, v := range d.links {
			if v != nil {
				possible = append(possible, v)
			}
		}
		res := false
		for _, v := range possible {
			if RecSearch(v, subs[1:]) {
				res = true
			}
		}
		return res
		// return RecSearch()
	}

	if d.links[j] == nil {
		return false
	}

	return RecSearch(d.links[j], subs[1:])
}
