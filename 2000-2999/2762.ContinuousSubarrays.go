func continuousSubarrays(nums []int) int64 {
    ans := int64(0)
    l, r := 0, 0
    j := 0
    for i := 0; i < len(nums); i++ {
        if i == 0 {
            l, r = nums[i]-2, nums[i]+2
            ans++
        } else if nums[i] >= l && nums[i] <= r {
            l, r = max(l, nums[i]-2), min(r, nums[i]+2)
            ans += int64(i - j + 1)
        } else {
            l, r = nums[i]-2, nums[i]+2
            for j = i - 1; nums[j] >= nums[i]-2 && nums[j] <= nums[i]+2; j-- {
                l, r = max(l, nums[j]-2), min(r, nums[j]+2)
            }
            j++
            ans += int64(i - j + 1)
        }
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}