type Node struct {
    Key string
    Val int
}

type AllOne struct {
    Map map[string]int
    Arr []*Node
}


/** Initialize your data structure here. */
func Constructor() AllOne {
    return AllOne{Map: make(map[string]int), Arr: []*Node{}}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
    this.Map[key]++
    n := len(this.Arr)
    idx := 0
    for ; idx < n; idx++ {
        if this.Arr[idx].Key == key {
            break
        }
    }
    /* Wrong idx ?
    idx := sort.Search(n, func(i int) bool {
        return this.Arr[i].Key == key
    }) */
    if idx == n {
        this.Arr = append([]*Node{&Node{Key: key, Val: 1}}, this.Arr...)
    } else {
        this.Arr[idx].Val++
        if idx < n-1 && this.Arr[idx].Val > this.Arr[idx+1].Val {
            sort.Slice(this.Arr, func(i, j int) bool {
                return this.Arr[i].Val <= this.Arr[j].Val
            })
        }
    }
}


/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
    if this.Map[key] == 0 {
        return
    }
    this.Map[key]--
    n := len(this.Arr)
    idx := 0
    for ; idx < n; idx++ {
        if this.Arr[idx].Key == key {
            break
        }
    }
    /* Wrong idx ?
    idx := sort.Search(n, func(i int) bool {
        return this.Arr[i].Key == key
    })*/
    if idx < n {
        this.Arr[idx].Val--
        if this.Arr[idx].Val == 0 {
            this.Arr = append(this.Arr[:idx], this.Arr[idx+1:]...)
        } else if idx > 0 && this.Arr[idx].Val < this.Arr[idx-1].Val {
            sort.Slice(this.Arr, func(i, j int) bool {
                return this.Arr[i].Val <= this.Arr[j].Val
            })
        }
    }
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
    n := len(this.Arr)
    if n == 0 {
        return ""
    }
    return this.Arr[n-1].Key
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
    if len(this.Arr) == 0 {
        return ""
    }
    return this.Arr[0].Key
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */