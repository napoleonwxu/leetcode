//Trie, O(K) + O(K), K: len(key)
type TrieNode struct {
    children map[byte]*TrieNode
    sum int
}

type MapSum struct {
    Map map[string]int
    root *TrieNode
}


/** Initialize your data structure here. */
func Constructor() MapSum {
    return MapSum{
        Map: make(map[string]int),
        root: &TrieNode{children: make(map[byte]*TrieNode)},
    }
}


func (this *MapSum) Insert(key string, val int)  {
    diff := val - this.Map[key]
    this.Map[key] = val
    cur := this.root
    for i := range key {
        _, ok := cur.children[key[i]]
        if !ok {
            trie := TrieNode{children: make(map[byte]*TrieNode)}
            cur.children[key[i]] = &trie
        }
        cur = cur.children[key[i]]
        cur.sum += diff
    }
}


func (this *MapSum) Sum(prefix string) int {
    cur := this.root
    for i := range prefix {
        _, ok := cur.children[prefix[i]]
        if !ok {
            return 0
        }
        cur = cur.children[prefix[i]]
    }
    return cur.sum
}


/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */