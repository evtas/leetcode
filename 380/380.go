package main

import "math/rand"

type RandomizedSet struct {
	nums []int
	m    map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		m:    make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.m[val]
	if !ok {
		return false
	}
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last]
	this.m[this.nums[id]] = id
	this.nums = this.nums[:last]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	println(obj.Insert(0))
	println(obj.Insert(1))
	println(obj.Remove(0))
	println(obj.Insert(2))
	println(obj.Remove(1))
	println(obj.GetRandom())
}
