package longestwordindict

import "testing"

func TestSolve(t *testing.T) {
	s := "abpcplea"
	dictionary := []string{"ale", "apple", "monkey", "plea"}
	res := findLongestWord(s, dictionary)
	if res != "apple" {
		t.Log(res)
		t.FailNow()
	}
}
