func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// O(n)
    stack := []int{}
    Map := make(map[int]int)
    for _, num := range nums2 {
        n := len(stack)
        for n > 0 && stack[n-1] < num {
            Map[stack[n-1]] = num
            stack = stack[:n-1]
            n--
        }
        stack = append(stack, num)
    }
    ans := []int{}
    for _, num := range nums1 {
        grater, ok := Map[num]
        if ok {
            ans = append(ans, grater)
        } else {
            ans = append(ans, -1)
        }
    }
    return ans
}