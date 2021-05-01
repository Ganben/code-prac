package main

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, Id int) int {

	em := findEmp(employees, Id)
	if em == nil {
		return 0
	}
	ans := 0
	ans += em.Importance
	for _, idx := range em.Subordinates {
		ans += getImportance(employees, idx)

	}
	return ans
}

func findEmp(employees []*Employee, id int) *Employee {
	for _, em := range employees {
		if id == em.Id {
			return em
		}
	}
	return nil
}

func main() {
	i := [][]int{}
}
