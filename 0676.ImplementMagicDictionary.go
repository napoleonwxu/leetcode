type MagicDictionary struct {
    Map map[int][]string
}


/** Initialize your data structure here. */
func Constructor() MagicDictionary {
    return MagicDictionary{make(map[int][]string)}
}


/** Build a dictionary through a list of words */
func (this *MagicDictionary) BuildDict(dict []string)  {
    for _, w := range dict {
        this.Map[len(w)] = append(this.Map[len(w)], w)
    }
}


/** Returns if there is any word in the trie that equals to the given word after modifying exactly one character */
func (this *MagicDictionary) Search(word string) bool {
    n := len(word)
    for _, w := range this.Map[n] {
        diff := 0
        for i := 0; i < n; i++ {
            if w[i] != word[i] {
                diff++
                if diff > 1 {
                    break
                }
            }
        }
        if diff == 1 {
            return true
        }
    }
    return false
}


/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dict);
 * param_2 := obj.Search(word);
 */