func intersection(nums1 []int, nums2 []int) []int {
    Map := make(map[int]bool, len(nums1))
    for _, num := range nums1 {
        Map[num] = true
    }
    ans := []int{}
    for _, num := range nums2 {
        if Map[num] {
            ans = append(ans, num)
            Map[num] = false
        }
    }
    return ans
}