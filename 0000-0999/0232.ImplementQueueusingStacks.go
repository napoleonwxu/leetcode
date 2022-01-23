type MyQueue struct {
    stackIn []int
    stackOut []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{stackIn: []int{}, stackOut: []int{}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.stackIn = append(this.stackIn, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    p := this.Peek()
    this.stackOut = this.stackOut[:len(this.stackOut)-1]
    return p
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.stackOut) == 0 {
        for i := len(this.stackIn)-1; i >= 0; i-- {
            this.stackOut = append(this.stackOut, this.stackIn[i])
        }
        this.stackIn = []int{} 
    }
    return this.stackOut[len(this.stackOut)-1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.stackIn) == 0 && len(this.stackOut) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */