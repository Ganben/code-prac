package solution

type MountainArray struct {
}

func (this *MountainArray) get(index int) int {}
func (this *MountainArray) length() int       {}

func findInMountainArray(target int, mountain *MountainArray) int {
	l := mountain.length()
	mtimes := 0
	index := 0
	for mtimes < 100 {
		if mountain.get(index) == target {
			return index
		}
		if mountain.get(index) < target && index < l/2 {
			index++
		}
		if mountain.get(index) > target && index < l/2 {
			index = l - 1
		}
		if mountain.get(index) < target && index > l/2 {
			index--
		}
		if mountain.get(index) > target && index > l/2 {
			index++
		}
		mtimes++
	}
	return -1

}
