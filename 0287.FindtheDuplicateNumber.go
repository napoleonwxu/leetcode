func findDuplicate(nums []int) int {
    // No modify, O(n) + O(1)
    slow, fast := nums[0], nums[nums[0]]
    for slow != fast {
        slow, fast = nums[slow], nums[nums[fast]]
    }
    slow2 := 0
    for slow != slow2 {
        slow, slow2 = nums[slow], nums[slow2]
    }
    return slow
    /* Swap, O(n) + O(1)
    for i := range nums {
        for nums[i] != nums[nums[i]-1] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i, num := range nums {
        if num != i+1 {
            return num
        }
    }
    return -1
    */
}