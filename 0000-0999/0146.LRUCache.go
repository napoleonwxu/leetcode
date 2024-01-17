type Node struct {
    key int
    val int
    pre *Node
    nxt *Node
}

type LRUCache struct {
    cap  int
    m    map[int]*Node
    head *Node
    tail *Node
}

func Constructor(capacity int) LRUCache {
    h := &Node{}
    t := &Node{pre: h}
    h.nxt = t
    return LRUCache{cap: capacity, m: make(map[int]*Node, capacity+1), head: h, tail: t}
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.m[key]
    if !ok {
        return -1
    }
    this.remove(node)
    this.add(node)
    return node.val
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.m[key]; ok {
        this.remove(node)
    }
    node := &Node{key: key, val: value}
    this.add(node)
    this.m[key] = node
    if len(this.m) > this.cap {
        delete(this.m, this.tail.pre.key)
        this.remove(this.tail.pre)
    }
}

func (this *LRUCache) add(node *Node) {
    n := this.head.nxt
    this.head.nxt = node
    node.nxt = n
    n.pre = node
    node.pre = this.head
}

func (this *LRUCache) remove(node *Node) {
    p, n := node.pre, node.nxt
    p.nxt = n
    n.pre = p
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */