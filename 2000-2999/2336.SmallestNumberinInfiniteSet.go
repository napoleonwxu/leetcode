type Node struct {
    val  int
    next *Node
}

type SmallestInfiniteSet struct {
    head *Node
}

func Constructor() SmallestInfiniteSet {
    return SmallestInfiniteSet{head: &Node{val: 1, next: nil}}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
    smallest := this.head.val
    if this.head.next == nil {
        this.head.val++
    } else {
        this.head = this.head.next
    }
    return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
    fake := &Node{val: 0, next: this.head}
    pre, cur := fake, fake.next
    for cur != nil && cur.val < num {
        pre, cur = pre.next, cur.next
    }
    if cur == nil || cur.val == num {
        return
    }
    node := &Node{val: num, next: cur}
    pre.next = node
    this.head = fake.next
}

/**
* Your SmallestInfiniteSet object will be instantiated and called as such:
* obj := Constructor();
* param_1 := obj.PopSmallest();
* obj.AddBack(num);
*/