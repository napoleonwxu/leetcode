type Node struct {
    Map map[byte]*Node
    word bool
}

type Trie struct {
    root *Node
}


/** Initialize your data structure here. */
func Constructor() Trie {
    node := &Node{Map: make(map[byte]*Node)}
    return Trie{root: node}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    node := this.root
    for i := range word {
        if node.Map[word[i]] == nil {
            node.Map[word[i]] = &Node{Map: make(map[byte]*Node)}
        }
        node = node.Map[word[i]]
    }
    node.word = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.root
    for i := range word {
        if node.Map[word[i]] == nil {
            return false
        }
        node = node.Map[word[i]]
    }
    return node.word
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this.root
    for i := range prefix {
        if node.Map[prefix[i]] == nil {
            return false
        }
        node = node.Map[prefix[i]]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */