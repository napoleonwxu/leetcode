func findDuplicates(nums []int) []int {
    ans := []int{}
    // O(n) + O(1)
    for i := range nums {
        idx := abs(nums[i]) - 1
        if nums[idx] < 0 {
            ans = append(ans, idx+1)
        } else {
            nums[idx] = -nums[idx]
        }
    }
    /* O(2n) = O(n), no extra space
    // The outer for: O(n); the inner for(while): O(n), because every swap put a new num in correct location and there are totally n number.
    for i := range nums {
        for nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i, num := range nums {
        if num != i+1 {
            ans = append(ans, num)
        }
    }
    */
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}