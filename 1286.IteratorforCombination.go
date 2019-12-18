type CombinationIterator struct {
    iter []string
    idx int
}


func Constructor(characters string, combinationLength int) CombinationIterator {
    c := CombinationIterator{iter: []string{}, idx: 0}
    c.dfs(characters, "", combinationLength, 0)
    return c
}


func (this *CombinationIterator) Next() string {
    this.idx++
    return this.iter[this.idx-1]
}


func (this *CombinationIterator) HasNext() bool {
    return this.idx < len(this.iter)
}

func (this *CombinationIterator) dfs(s, path string, n, i int) {
    if len(path) == n {
        this.iter = append(this.iter, path)
        return
    }
    for j := i; j < len(s); j++ {
        this.dfs(s, path+s[j:j+1], n, j+1)
    }
}

/**
 * Your CombinationIterator object will be instantiated and called as such:
 * obj := Constructor(characters, combinationLength);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */