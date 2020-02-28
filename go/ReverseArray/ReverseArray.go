package ReverseArray

func ReverseArray(array []int) {
	i := 0
	u := len(array) - 1
	for i < u {
		array[i], array[u] = array[u], array[i]
		i, u = i+1, u-1
	}
}