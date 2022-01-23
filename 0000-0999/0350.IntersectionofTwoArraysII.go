func intersect(nums1 []int, nums2 []int) []int {
    Map := make(map[int]int, len(nums1))
    for _, num := range nums1 {
        Map[num]++
    }
    ans := []int{}
    for _, num := range nums2 {
        if Map[num] > 0 {
            ans = append(ans, num)
            Map[num]--
        }
    }
    return ans
}