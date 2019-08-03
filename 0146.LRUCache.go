type Node struct {
    key int
    val int
    pre *Node
    nxt *Node
}

type LRUCache struct {
    size int
    Map map[int]*Node
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    h := &Node{key: 0, val: 0}
    t := &Node{key: 0, val: 0}
    h.nxt = t
    t.pre = h
    return LRUCache{size: capacity, Map: make(map[int]*Node, capacity+1), head: h, tail: t}
}


func (this *LRUCache) Get(key int) int {
    node, ok := this.Map[key]
    if ok {
        this.remove(node)
        this.add(node)
        return node.val
    }
    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.Map[key]
    if ok {
        this.remove(node)
    }
    node = &Node{key: key, val: value}
    this.add(node)
    this.Map[key] = node
    if len(this.Map) > this.size {
        delete(this.Map, this.tail.pre.key)
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