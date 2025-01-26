type MyCircularQueue struct {
    queue             []int
    size, front, rear int
}

func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{queue: make([]int, k), size: 0, front: 0, rear: -1}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    this.rear = (this.rear + 1) % len(this.queue)
    this.queue[this.rear] = value
    this.size++
    return true
}

func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false
    }
    this.front = (this.front + 1) % len(this.queue)
    this.size--
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.front]
}

func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.queue[this.rear]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return this.size == len(this.queue)
}

/**
* Your MyCircularQueue object will be instantiated and called as such:
* obj := Constructor(k);
* param_1 := obj.EnQueue(value);
* param_2 := obj.DeQueue();
* param_3 := obj.Front();
* param_4 := obj.Rear();
* param_5 := obj.IsEmpty();
* param_6 := obj.IsFull();
*/