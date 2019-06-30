func findNumberOfLIS(nums []int) int {
    n := len(nums)
    lens := make([]int, n)
    cnts := make([]int, n)
    maxlen, ans := 0, 0
    for j := 0; j < n; j++ {
        lens[j] = 1
        cnts[j] = 1
        for i := 0; i < j; i++ {
            if nums[i] < nums[j] {
                if lens[i] >= lens[j] {
                    lens[j] = lens[i] + 1
                    cnts[j] = cnts[i]
                } else if lens[i]+1 == lens[j] {
                    cnts[j] += cnts[i]
                }
            }
        }
        if lens[j] > maxlen {
            maxlen = lens[j]
            ans = cnts[j]
        } else if lens[j] == maxlen {
            ans += cnts[j]
        }
    }
    return ans
}