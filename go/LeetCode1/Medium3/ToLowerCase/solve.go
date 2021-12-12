package tolowercase

func toLowerCase(s string) string {
	res := ""
	for _, v := range s {

		if v-'Z' <= 0 && v-'A' >= 0 {
			res += string(v + 'a' - 'A')
		} else {
			res += string(v)
		}
	}
	return res
}
