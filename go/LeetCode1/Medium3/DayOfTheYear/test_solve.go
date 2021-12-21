package dayoftheyear

import (
	"testing"
)

func TestSolve(t *testing.T) {
	date := "2019-01-09"
	res := 9
	if i := dayOfYear(date); i != res {
		t.FailNow()
	}
}
