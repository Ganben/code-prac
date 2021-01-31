package Solution

func isPalindrome(x int) bool {
	//
	if x < 0 || ( x!= 0 && x%10 == 0 ) {
		return false
	}
	halfReverseX := 0
	for ; x > halfReverseX; x /= 10 {
		halfReverseX = halfReverseX*10 + x%10
	}
	//
	return x == halfReverseX || halfReverseX/10 == x
}

func isPalindrome2(x int) bool {
	//
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	return x == reverse(x)
}

func reverse(x int) int {
	res := 0
	for ; x != 0; x /= 10 {
		res = res*10 + x%10
	}
	return res
}