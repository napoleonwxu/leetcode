func minCost(nums []int, cost []int) int64 {
    left, right := min(nums), max(nums)
    ans := getCost(nums, cost, left)
    for left < right {
        mid := (left + right) / 2
        c1 := getCost(nums, cost, mid)
        c2 := getCost(nums, cost, mid+1)
        if c1 < c2 {
            ans = c1
            right = mid
        } else {
            ans = c2
            left = mid + 1
        }
    }
    return ans
}

func getCost(nums, cost []int, target int) int64 {
    ans := int64(0)
    for i := 0; i < len(nums); i++ {
        ans += int64(abs(nums[i]-target)) * int64(cost[i])
    }
    return ans
}

func min(nums []int) int {
    ans := nums[0]
    for _, num := range nums {
        if num < ans {
            ans = num
        }
    }
    return ans
}

func max(nums []int) int {
    ans := nums[0]
    for _, num := range nums {
        if num > ans {
            ans = num
        }
    }
    return ans
}

func abs(x int) int {
    if x > 0 {
        return x
    }
    return -x
}
