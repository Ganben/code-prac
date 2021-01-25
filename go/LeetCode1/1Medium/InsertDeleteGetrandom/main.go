package main

import "math/rand"

type RandomizedSet struct {
	dict map[int]int
	list []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	dict := map[int]int{}
	list := []int{}

	return RandomizedSet{dict, list}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, exist := this.dict[val]
	if exist {
		return false
	}
	this.dict[val] = len(this.list)
	this.list = append(this.list, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	_, exist := this.dict[val]
	if exist {
		delete(this.dict, val)
		this.list = []int{}
		for k, _ := range this.dict {
			this.list = append(this.list, k)
		}
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	i := rand.Intn(len(this.list))
	return this.list[i]
}
