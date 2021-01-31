package IntToWords

func numberToWords(num int) string {
	if num == 0 {
		return "Zero"
	}
	billion := num / 1000000000
	million := (num - billion*1000000000) / 1000000
	thousand := (num - billion*1000000000 - million*1000000) / 1000
	rest := num % 1000
	result := ""
	if billion != 0 {
		result = f3(billion) + " Billion"
	}
	if million != 0 {
		// add if billion exist
		if billion != 0 {
			result += " "
		}
		result += f3(million) + " Million"
	}
	if thousand != 0 {
		// add if billion exist
		if len(result) > 0 {
			result += " "
		}
		result += f3(thousand) + " Thousand"
	}
	if rest != 0 {
		if len(result) > 0 {
			result += " "
		}
		result += f3(rest)
	}
	return result
}

func f3(num int) string {

}
