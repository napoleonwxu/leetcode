func findDisappearedNumbers(nums []int) []int {
    ans := []int{}
    // O(n) + O(1)
    for i := range nums {
        idx := abs(nums[i]) - 1
        nums[idx] = -abs(nums[idx])
    }
    for i, num := range nums {
        if num > 0 {
            ans = append(ans, i+1)
        }
    }
    /* O(2n) = O(n), every swap put a new num in correct location and there are totally n number.
    for i := range nums {
        for nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i := range nums {
        if i+1 != nums[i] {
            ans = append(ans, i+1)
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