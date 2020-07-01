package WordLadder

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if endWord == "" || beginWord == "" || len(wordList) == 0 {
		return 0
	}
	l := len(beginWord)
	d := make(map[string][]string)
	for _, w := range wordList {
		for i := 0; i < l; i++ {
			k := w[:i] + "*" + w[i+1:]
			d[k] = append(d[k], w)
		}
	}
	queWord := []string{}
	queWord = append(queWord, beginWord)
	queLevel := []int{}
	queLevel = append(queLevel, 1)
	visited := make(map[string]bool, len(wordList))
	visited[beginWord] = true
	for len(queWord) > 0 {
		currentWord := queWord[len(queWord)-1]
		level := queLevel[len(queLevel)-1]
		queWord = queWord[:len(queWord)-1]
		queLevel = queLevel[:len(queLevel)-1]

		for i := 0; i < l; i++ {
			testKey := currentWord[:i] + "*" + currentWord[i+1:]
			for _, w := range d[currentWord] {
				if w == endWord {
					return level + 1
				}
				_, ok := visited[w]
				if !ok {
					visited[w] = true
					queWord = append(queWord, w)
					queLevel = append(queLevel, level+1)
				}
			}
			d[testKey] = []string{}
		}
	}
	return 0
}
