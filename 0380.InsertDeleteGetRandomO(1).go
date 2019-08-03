type RandomizedSet struct {
    idx map[int]int
    num []int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    return RandomizedSet{idx: make(map[int]int), num: []int{}}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if _, ok := this.idx[val]; ok {
        return false
    }
    this.idx[val] = len(this.num)
    this.num = append(this.num, val)
    return true
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    i, ok := this.idx[val]
    if !ok {
        return false
    }
    last := len(this.num) - 1
    if last != i {
        this.num[i] = this.num[last]
        this.idx[this.num[i]] = i
    }
    this.num = this.num[:last]
    delete(this.idx, val)
    return true
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    return this.num[rand.Intn(len(this.num))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */