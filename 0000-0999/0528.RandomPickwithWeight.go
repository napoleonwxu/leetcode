type Solution struct {
    sum []int
}


func Constructor(w []int) Solution {
    sum := make([]int, len(w))
    sum[0] = w[0]
    for i := 1; i < len(w); i++ {
        sum[i] = sum[i-1] + w[i]
    }
    return Solution{sum: sum}
}


func (this *Solution) PickIndex() int {
    /*
    r := rand.Intn(this.sum[len(this.sum)-1])
    idx := sort.Search(len(this.sum), func(i int) bool {
        return this.sum[i] > r
    })
    return idx
    */
    r := rand.Intn(this.sum[len(this.sum)-1]) + 1
    left, right := 0, len(this.sum)-1
    for left < right {
        mid := (left+right) >> 1
        if this.sum[mid] == r {
            return mid
        } else if this.sum[mid] < r {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */