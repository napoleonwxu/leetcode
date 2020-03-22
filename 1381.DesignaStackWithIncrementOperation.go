type CustomStack struct {
    stack, inc []int
    size, idx int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{stack: make([]int, maxSize), inc: make([]int, maxSize), size: maxSize, idx: 0}
}


func (this *CustomStack) Push(x int)  {
    // O(1)
    if this.idx < this.size {
        this.stack[this.idx] = x
        this.idx++
    }
}


func (this *CustomStack) Pop() int {
    // O(1)
    if this.idx == 0 {
        return -1
    }
    this.idx--
    if this.idx > 0 {
        this.inc[this.idx-1] += this.inc[this.idx]
    }
    this.stack[this.idx] += this.inc[this.idx]
    this.inc[this.idx] = 0
    return this.stack[this.idx]
}


func (this *CustomStack) Increment(k int, val int) {
    // O(1)
    i := min(k, this.idx) - 1
    if i >= 0 {
        this.inc[i] += val
    }
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */