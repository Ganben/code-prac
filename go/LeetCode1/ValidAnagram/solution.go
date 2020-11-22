package ValidAnagram



func isAnagram(s string, t string) bool {
var (
	map1 map[rune]int
	map2 map[rune]int
)

map1 = map[rune]int{}
map2 = map[rune]int{}

for _, c := range "abcdefghijklmnopqrstuvwxyz" {
	map1[c] = 0
	map2[c] = 0
}

for _, c := range s {
	map1[c] += 1
}

for _, c := range t {
	map2[c] += 1
}

for k,_ := range map1 {
	if map1[k] != map2[k] {
		return false
	}
}
return true
}