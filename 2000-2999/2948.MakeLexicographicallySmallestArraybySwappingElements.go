import "sort"

func lexicographicallySmallestArray(nums []int, limit int) []int {
    n := len(nums)
    tmp := make([][]int, n)
    for i, num := range nums {
        tmp[i] = []int{num, i}
    }
    sort.Slice(tmp, func(i, j int) bool {
        return tmp[i][0] < tmp[j][0]
    })
    groups := [][][]int{}
    groups = append(groups, [][]int{tmp[0]})
    for i := 1; i < n; i++ {
        if tmp[i][0]-tmp[i-1][0] <= limit {
            m := len(groups)
            groups[m-1] = append(groups[m-1], tmp[i])
        } else {
            groups = append(groups, [][]int{tmp[i]})
        }
    }
    ans := make([]int, n)
    for _, group := range groups {
        idx := make([]int, len(group))
        for i, g := range group {
            idx[i] = g[1]
        }
        sort.Ints(idx)
        for i := range idx {
            ans[idx[i]] = group[i][0]
        }
    }
    return ans
}