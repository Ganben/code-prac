package solution

import "fmt"

func fullJustify(words []string, maxWidth int) []string {
	row := []string{}
	ans := []string{}
	ifEnd := false
	// detect which words put in to one line;
	for i := 0; i < len(words); i++ {
		ifEnd = false
		if i == len(words)-1 {
			fmt.Printf("end row at %v\n", i)
			ifEnd = true
		}
		if addRowTest(row, maxWidth, words[i]) {
			row = append(row, words[i])
		} else {
			// make row
			fmt.Printf("make row at %v\n", i)
			// this is called by next word (end word, therefore no end)
			ans = makeNewRow(ans, row, false, maxWidth)
			row = []string{}
			row = append(row, words[i])
		}
		if ifEnd {
			ans = makeNewRow(ans, row, ifEnd, maxWidth)
			break
		}
	}
	return ans
}

func addRowTest(row []string, maxWidth int, word string) bool {
	width := 0
	for _, w := range row {
		width += len(w)
		width++
	}
	if width+len(word) > maxWidth {
		return false
	}
	return true
}

func makeNewRow(ans []string, row []string, ifEnd bool, maxWidth int) []string {
	spaces := 0
	wordWidth := 0
	for _, w := range row {
		wordWidth += len(w)
	}
	spaces = maxWidth - wordWidth
	spaceCase := 1
	if ifEnd {
		// no assign spaces
		rowStr := ""
		for i, w := range row {
			rowStr += w
			if i < len(row)-1 {
				rowStr += " "
			}
		}
		if len(rowStr) < maxWidth {
			restN := maxWidth - len(rowStr)
			for i := 0; i < restN; i++ {
				rowStr += " "
			}
		}
		fmt.Printf("end row is %v\n", rowStr)
		return append(ans, rowStr)
	}

	if spaces == len(row)-1 && spaces > 0 {
		spaceCase = 1
	} else if spaces == 0 {
		spaceCase = 0
	} else {
		if len(row)-1 > 0 {
			spaceCase = int(spaces / (len(row) - 1))
			if spaces%(len(row)-1) > 0 {
				spaceCase++
			}
		} else {
			spaceCase = spaces
		}
	}
	// a to end loop
	rowStr := ""
	for i, w := range row {
		rowStr += w
		fmt.Printf("fill spaces at %v x %v\n", w, spaceCase)
		if i < len(row)-1 {
			for i := 0; i < spaceCase; i++ {
				rowStr += " "
			}
		} else if i == len(row)-1 && spaceCase > 0 {
			for i := 0; i < spaceCase; i++ {
				rowStr += " "
			}
		} else {
			break
		}
		// update rest spaces and spaceCase
		spaces -= spaceCase
		if len(row)-2-i > 0 {
			spaceCase = int(spaces / (len(row) - 1 - i - 1))
			if spaces%(len(row)-2-i) > 0 {
				spaceCase++
			}
		} else {
			spaceCase = spaces
		}

	}
	return append(ans, rowStr)
}
