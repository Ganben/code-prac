package RectangleArea

func computeArea(A int, B, C, D, E, F, G, H int) int {

	r1 := rectArea(A, B, C, D)
	r2 := rectArea(E, F, G, H)

	return r1 + r2 - interceptArea(A, B, C, D, E, F, G, H)

}

func rectArea(A, B, C, D int) int {
	return abs(A, C) * abs(B, D)
}

func interceptArea(A, B, C, D, E, F, G, H int) int {
	ix := 0
	iy := 0
	x1, y1 := min(A, C), min(B, D)
	x2, y2 := min(E, G), min(F, H)
	x1m, y1m := max(A, C), max(B, D)
	x2m, y2m := max(E, G), max(F, H)
	ix = min(x1m, x2m) - max(x1, x2)
	iy = min(y1m, y2m) - max(y1, y2)
	if ix < 0 || iy < 0 {
		return 0
	}
	return ix * iy
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
