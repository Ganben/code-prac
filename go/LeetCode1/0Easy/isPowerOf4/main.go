package main

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	for n >= 4 {
		if n%4 == 0 {
			n /= 4
		} else {
			return false
		}
		if n == 1 {
			return true
		}
	}
	return true
}

func main() {

}
