import "sort"

func minOperations(nums []int, queries []int) []int64 {
    n := len(nums)
    sort.Ints(nums)
    preSum := make([]int64, n+1)
    for i := 0; i < n; i++ {
        preSum[i+1] = preSum[i] + int64(nums[i])
    }
    ans := make([]int64, len(queries))
    for i, q := range queries {
        j := binSearchLeft(nums, q)
        ans[i] += (int64(j*q) - preSum[j]) + (preSum[n] - preSum[j] - int64((n-j)*q))
    }
    return ans
}

func binSearchLeft(nums []int, target int) int {
    left, right := 0, len(nums)
    for left < right {
        mid := (left + right) / 2
        if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}