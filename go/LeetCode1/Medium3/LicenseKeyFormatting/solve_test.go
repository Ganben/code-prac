package licensekeyformatting

import "testing"

func TestSolve(t *testing.T) {
	t.Log("begin test")
	r := licenseKeyFormatting("5F3Z-2e-9-w", 4)
	if r != "5F3Z-2E9W" {
		t.Log(r)
		t.FailNow()
	}

}
func TestSolve2(t *testing.T) {
	t.Log("begin test")
	r := licenseKeyFormatting("2-5g-3-J", 2)
	if r != "2-5G-3J" {
		t.Log(r)
		t.FailNow()
	}

}
