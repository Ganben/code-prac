package ContestW

func modifyString(s string) string {
	sarr := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			avoids := []byte{}
			if i > 0 {
				avoids = append(avoids, sarr[i-1])
			}
			if i < len(s)-1 {
				avoids = append(avoids, s[i+1])
			}
			sarr[i] = choose(avoids)
		} else {
			sarr[i] = s[i]
		}
	}
	ans := ""
	for _, char := range sarr {
		ans += string(char)
	}
	return ans
}

func choose(avoids []byte) byte {
	for i := 0; i < 26; i++ {
		flag := false
		for _, v := range avoids {
			if i-int(v-'a') == 0 {
				flag = true
			}
		}
		if !flag {
			return byte(int('a') + i)
		}
	}
	return 'a'
}
