package CompareVersionNumber

func compareVersion(version1, version2 string) int {
	p1, p2 := 0, 0
	b := 2
	if len(version1) > 1 {
		if version1[1] != '.' {
			b = 1
		}
	} else if len(version2) > 1 {
		if version2[1] != '.' {
			b = 1
		}
	}

	for p1 <= len(version1) && p2 <= len(version2) {

		if version1[p1] == version2[p2] {
			p1 += b
			p2 += b
		} else {
			if int(version1[p1]) > int(version2[p2]) {
				// p1 later
				return 1
			} else {
				// p2 later
				return -1
			}
		}

	}

	if p1 > len(version1) && p2 < len(version2) {
		// p1 later
		return 1
	}

	if p1 > len(version1) && p2 < len(version2) {
		// p2 later
		return -1
	}

	return 0
}
