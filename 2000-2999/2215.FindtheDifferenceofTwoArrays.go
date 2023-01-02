func findDifference(nums1 []int, nums2 []int) [][]int {
    map1, map2 := make(map[int]bool), make(map[int]bool)
    for _, num := range nums1 {
        map1[num] = true
    }
    for _, num := range nums2 {
        map2[num] = true
    }
    
    ans := make([][]int, 2)
    for num := range map1 {
        if !map2[num] {
            ans[0] = append(ans[0], num)
        }
    }
    for num := range map2 {
        if !map1[num] {
            ans[1] = append(ans[1], num)
        }
    }
    return ans
}