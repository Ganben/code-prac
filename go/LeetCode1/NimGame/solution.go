package NimGame

func canWinNim(n int) bool {
	b := n % 4
	if b == 0 {
		return false
	}
	return true
}
