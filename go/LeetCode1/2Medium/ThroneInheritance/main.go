package main

type ThroneInheritance struct {
	king  string
	edges map[string][]string
	dead  map[string]bool
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{kingName, map[string][]string{}, map[string]bool{}}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.edges[parentName] = append(this.edges[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string
	ans = []string{}
	var preorder func(string)
	preorder = func(name string) {
		if !this.dead[name] {
			ans = append(ans, name)
		}
		for _, childName := range this.edges[name] {
			preorder(childName)
		}
	}
	preorder(this.king)
	return ans
}

/**
 * Your ThroneInheritance object will be instantiated and called as such:
 * obj := Constructor(kingName);
 * obj.Birth(parentName,childName);
 * obj.Death(name);
 * param_3 := obj.GetInheritanceOrder();
 */
