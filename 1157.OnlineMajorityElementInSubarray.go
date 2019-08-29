type MajorityChecker struct {
    idxs map[int][]int
}


func Constructor(arr []int) MajorityChecker {
    Map := make(map[int][]int)
    for i, num := range arr {
        Map[num] = append(Map[num], i)
    }
    return MajorityChecker{idxs: Map}
}


func (this *MajorityChecker) Query(left int, right int, threshold int) int {
    for num, idx := range this.idxs {
        if len(idx) < threshold {
            continue
        }
        lo := sort.Search(len(idx), func(i int) bool {
            return idx[i] >= left
        })
        hi := sort.Search(len(idx), func(i int) bool {
            return idx[i] > right
        })
        if hi-lo >= threshold {
            return num
        }
    }
    return -1
}


/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */