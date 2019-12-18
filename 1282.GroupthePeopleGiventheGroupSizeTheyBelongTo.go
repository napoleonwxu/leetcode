func groupThePeople(groupSizes []int) [][]int {
    ans := [][]int{}
    idx := make(map[int]int)
    for i, size := range groupSizes {
        j, ok := idx[size]
        if ok {
            ans[j] = append(ans[j], i)
        } else {
            j = len(ans)
            idx[size] = j
            ans = append(ans, []int{i})
        }
        if len(ans[j]) == size {
            delete(idx, size)
        }
    }
    return ans
}