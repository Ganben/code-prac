package WordSearchii

var (
	rBoard  [][]byte
	rResult []string
)

type Trie struct {
	links []*Trie
	isEnd bool
	val   byte
}

func Constructor() Trie {
	t := Trie{}
	t.links = make([]*Trie, 26)
	t.links = []*Trie{}
	for i := 0; i < 26; i++ {
		t.links = append(t.links, nil)
	}
	t.isEnd = false
	return t
}

func (this *Trie) Insert(word string) {
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
				cur.links[j].val = byte(v)
				if i == len(word)-1 {
					cur.links[j].isEnd = true
				}
				cur = cur.links[j]
			} else {
				if i == len(word)-1 {
					cur.links[j].isEnd = true
				}
				cur = cur.links[j]

			}
		}
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	if len(word) == 0 {
		return false
	}
	for i, v := range word {
		j := int(v - 'a')
		if cur.links[j] == nil {
			return false
		}
		cur = cur.links[j]
		if i == len(word)-1 {
			return cur.isEnd
		}
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i, v := range prefix {
		j := int(v - 'a')
		if cur.links[j] == nil {
			return false
		}
		cur = cur.links[j]
		if i == len(prefix)-1 {
			return true
		}
	}
	return false
}

func findWords(board [][]byte, words []string) []string {
	rBoard = board
	rResult = []string{}
	p := Constructor()
	n := &p
	for _, w := range words {
		n.Insert(w)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if n.links[board[i][j]] != nil {
				backtracking(i, j, n)
			}
		}
	}
	return rResult

}

func backtracking(row, col int, node *Trie) {

}
