package main

import "math/rand"

type RandomizedSet struct {
	valMap map[int]int
	arr    []int
}

// ConstructorRandomizedSet /** Initialize your data structure here. */
func Constructor_RandomizedSet() RandomizedSet {
	return RandomizedSet{map[int]int{}, []int{}}
}

// Insert /** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.valMap[val]; ok {
		return false
	}
	this.valMap[val] = len(this.arr)
	this.arr = append(this.arr, val)
	return true
}

// Remove /** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if index, ok := this.valMap[val]; !ok {
		return false
	} else {
		last := len(this.arr) - 1
		this.arr[index] = this.arr[last]
		this.valMap[this.arr[index]] = index
		this.arr = append(this.arr[:last])
		delete(this.valMap, val)
		return true
	}
}

// GetRandom /** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
