package main

import "fmt"

type Trie struct {
	next [26]*Trie
	end  bool
}

func (this *Trie) add(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.next[word[i]-'a'] == nil {
			root.next[word[i]-'a'] = &Trie{
				next: [26]*Trie{},
				end:  false,
			}
		}
		root = root.next[word[i]-'a']
	}
	root.end = true
}

func findAllConcatenatedWordsInADict(words []string) []string {
	trie := new(Trie)
	for _, word := range words {
		if len(word) != 0 {
			trie.add(word)
		}
	}
	var (
		result = make([]string, 0)
		dfs    func(word string, index int) bool
	)
	//
	dfs = func(word string, index int) bool {
		root := trie
		for i := index; i < len(word); i++ {
			if root.next[word[i]-'a'] == nil {
				return false
			}
			root = root.next[word[i]-'a']
			if root.end && dfs(word, i+1) {
				return true
			}
		}
		//
		return root.end && index != 0
	}
	//
	for _, word := range words {
		if len(word) != 0 && dfs(word, 0) {
			result = append(result, word)
		}
	}
	return result

}

func main() {
	words := []string{
		"cat",
		"cats",
		"catsdogcats",
		"dog",
		"dogcatsdog",
		"hippopotamuses",
		"rat",
		"ratcatdogcat",
	}
	fmt.Printf("%v", findAllConcatenatedWordsInADict(words))
}
