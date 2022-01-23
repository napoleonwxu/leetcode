type DinnerPlates struct {
    stack [][]int
    n int
    cnt int
    left int
    right int
}


func Constructor(capacity int) DinnerPlates {
    return DinnerPlates{stack: [][]int{}, n: capacity, cnt: 0, left: 0, right: 0}
}


func (this *DinnerPlates) Push(val int)  {
    for len(this.stack) > this.left && len(this.stack[this.left]) == this.n {
        this.left++
    }
    if this.left < len(this.stack) {
        this.stack[this.left] = append(this.stack[this.left], val)
    } else {
        this.stack = append(this.stack, []int{val})
    }
    this.cnt++
    this.right = max(this.right, this.left)
}


func (this *DinnerPlates) Pop() int {
    if this.cnt == 0 {
        return -1
    }
    for this.right >= 0 && len(this.stack[this.right]) == 0 {
        this.right--
    }
    l := len(this.stack[this.right])
    p := this.stack[this.right][l-1]
    this.stack[this.right] = this.stack[this.right][:l-1]
    this.cnt--
    this.left = min(this.left, this.right)
    return p
}


func (this *DinnerPlates) PopAtStack(index int) int {
    if len(this.stack) <= index || len(this.stack[index]) == 0 {
        return -1
    }
    l := len(this.stack[index])
    p := this.stack[index][l-1]
    this.stack[index] = this.stack[index][:l-1]
    this.cnt--
    this.left = min(this.left, index)
    return p
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */