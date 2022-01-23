func findNumberOfLIS(nums []int) int {
    n := len(nums)
    lens := make([]int, n)
    cnts := make([]int, n)
    maxLen, ans := 0, 0
    for i := 0; i < n; i++ {
        lens[i], cnts[i] = 1, 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                if lens[i] <= lens[j] {
                    lens[i] = lens[j] + 1
                    cnts[i] = cnts[j]
                } else if lens[i] == lens[j]+1 {
                    cnts[i] += cnts[j]
                }
            }
        }
        if lens[i] > maxLen {
            maxLen, ans = lens[i], cnts[i]
        } else if lens[i] == maxLen {
            ans += cnts[i]
        }
    }
    return ans
}