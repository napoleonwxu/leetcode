type MinStack struct {
    stack []int
    min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{stack: []int{}, min: math.MaxInt32}
}


func (this *MinStack) Push(x int)  {
    if x <= this.min { // will WA if only "<"
        this.stack = append(this.stack, this.min)
        this.min = x
    }
    this.stack = append(this.stack, x)
}


func (this *MinStack) Pop()  {
    n := len(this.stack)
    if this.stack[n-1] == this.min {
        this.min = this.stack[n-2]
        this.stack = this.stack[:n-2]
    } else {
        this.stack = this.stack[:n-1]
    }
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */