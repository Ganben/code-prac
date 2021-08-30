package licensekeyformatting

func licenseKeyFormatting(s string, k int) string {
	list := []rune{}
	for _, v := range s {
		if v != '-' {
			if v >= 'a' && v <= 'z' {
				v += 'A' - 'a'
			}
			list = append(list, v)
		}
	}
	//
	listLen := len(list)
	first := listLen % k
	result := ""
	if first != 0 {
		result = string(list[:first])
	}
	for i := first; i < listLen; i += k {
		if i != 0 {
			result += "-"
		}
		result += string(list[i : i+k])
	}
	return result

}
