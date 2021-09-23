package keyboardrow

import "testing"

func TestSolve(t *testing.T) {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	res := findWords(words)
	if res[0] != "Alaska" {
		t.FailNow()
	}
}
