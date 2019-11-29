func minSubArrayLen(s int, nums []int) int {
	// O(n)
    ans := len(nums) + 1
    sum, i := 0, 0
    for j, num := range nums {
        sum += num
        for sum >= s {
            ans = min(ans, j-i+1)
            sum -= nums[i]
            i++
        }
    }
    if ans > len(nums) {
        return 0
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}