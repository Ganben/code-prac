package main

func minMutation(start string, end string, bank []string) int {
arr := []int{}
curr := make([]bool, len(bank))

dfs(start, end bank, curr, 0, &arr)
sort.Ints(arr)
if len(arr) == 0 {
	return -1
}
return arr[0]
}

func dfs(temp, end, string, bank []string, visited []bool, level int, result *[]int) {
	if temp == end {
		*result = append(*result, level)
		return
	}
	//
	for i:= 0; i < len(bank); i++ {
		if !visited[i] && checkDifference(temp, bank[i]) {
			visited[i] = true
			//
			dfs(bank[i], end, bank, visited, level+1, result)
			visited[i] = false
		}
	}
}

func checkDifference(temp string, bankByOne string) bool {
	diff := 0
	for i := 0; i < len(temp); i ++ {
		if temp[i] != bankByOne[i] {
			diff ++
		}
	}
	if diff == 1 {
		return true
	}
	return false
}

func main() {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Printf("%v\n", minMutation(start, end, bank))	
}