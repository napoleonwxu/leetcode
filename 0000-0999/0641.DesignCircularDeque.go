type MyCircularDeque struct {
	queue            []int
	cnt, front, last int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{queue: make([]int, k), cnt: 0, front: 0, last: k - 1}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.front = (this.front + len(this.queue) - 1) % len(this.queue)
	this.queue[this.front] = value
	this.cnt++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.last = (this.last + 1) % len(this.queue)
	this.queue[this.last] = value
	this.cnt++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % len(this.queue)
	this.cnt--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.last = (this.last + len(this.queue) - 1) % len(this.queue)
	this.cnt--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.last]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.cnt == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.cnt == len(this.queue)
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */