func nextPermutation(nums []int)  {
    n := len(nums)
    i := n - 2
    for ; i >= 0 && nums[i] >= nums[i+1]; i-- {}
    if i < 0 {
        reverse(nums, 0, n-1)
        return
    }
    j := n - 1
    for ; j > i && nums[j] <= nums[i]; j-- {}
    nums[i], nums[j] = nums[j], nums[i]
    reverse(nums, i+1, n-1)
}

func reverse(nums []int, left, right int) {
    for ; left < right; left, right = left+1, right-1 {
        nums[left], nums[right] = nums[right], nums[left]
    }
}