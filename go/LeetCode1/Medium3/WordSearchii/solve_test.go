package wordsearchii

import "testing"

func TestSolve(t *testing.T) {
	board := [][]byte{
		[]byte{'o', 'a', 'a', 'n'},
		[]byte{'e', 't', 'a', 'e'},
		[]byte{'i', 'h', 'k', 'r'},
		[]byte{'i', 'f', 'l', 'v'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	res := findWord(board, words)
	if res[0] != "eat" || res[1] != "oath" {
		t.Log(res)
		t.FailNow()
	}
}
