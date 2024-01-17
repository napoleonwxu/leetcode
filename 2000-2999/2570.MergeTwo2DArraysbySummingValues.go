func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    m, n := len(nums1), len(nums2)
    tmp := make([]int, max(nums1[m-1][0], nums2[n-1][0])+1)
    for _, num := range nums1 {
        tmp[num[0]] += num[1]
    }
    for _, num := range nums2 {
        tmp[num[0]] += num[1]
    }
    ans := make([][]int, 0, len(tmp))
    for id, val := range tmp {
        if val != 0 {
            ans = append(ans, []int{id, val})
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}