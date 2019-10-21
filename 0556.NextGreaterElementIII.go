func nextGreaterElement(n int) int {
    nums := []int{}
    for n > 0 {
        nums = append(nums, n%10)
        n /= 10
    }
    reverse(nums, 0, len(nums)-1)
    ok := nextPermutation(nums)
    if !ok {
        return -1
    }
    ans := 0
    for _, num := range nums {
        ans = ans*10 + num
    }
    if ans > 1<<31-1 {
        return -1
    }
    return ans
}

func nextPermutation(nums []int) bool {
    n := len(nums)
    i := n - 2
    for ; i >= 0 && nums[i] >= nums[i+1]; i-- {}
    if i < 0 {
        return false
    }
    j := n - 1
    for ; j > i && nums[j] <= nums[i]; j-- {}
    nums[i], nums[j] = nums[j], nums[i]
    reverse(nums, i+1, n-1)
    return true
}

func reverse(nums []int, left, right int) {
    for ; left < right; left, right = left+1, right-1 {
        nums[left], nums[right] = nums[right], nums[left]
    }
}