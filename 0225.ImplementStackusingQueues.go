type MyStack struct {
    queue []int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
    return MyStack{queue: []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.queue = append(this.queue, x)
    for i := 1; i < len(this.queue); i++ {
        this.queue = append(this.queue, this.queue[0])
        this.queue = this.queue[1:]
    }
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    p := this.queue[0]
    this.queue = this.queue[1:]
    return p
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.queue[0]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.queue) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */