func nextGreaterElement(nums1 []int, nums2 []int) []int {
    stack, n := make([]int, len(nums2)), 0
    Map := make(map[int]int)
    for _, num := range nums2 {
        for n > 0 && stack[n-1] < num {
            Map[stack[n-1]] = num
            n--
        }
        stack[n] = num
        n++
    }
    ans := make([]int, len(nums1))
    for i, num := range nums1 {
        if _, ok := Map[num]; ok {
            ans[i] = Map[num]
        } else {
            ans[i] = -1
        }
    }
    return ans
}