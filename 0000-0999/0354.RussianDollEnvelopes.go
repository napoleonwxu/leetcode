func maxEnvelopes(envelopes [][]int) int {
    sort.Slice(envelopes, func(i, j int) bool {
        if envelopes[i][0] == envelopes[j][0] {
            return envelopes[i][1] > envelopes[j][1]
        }
        return envelopes[i][0] < envelopes[j][0]
    })
    dp := make([]int, 0, len(envelopes))
    for _, e := range envelopes {
        idx := binInsertLeft(dp, e[1])
        if idx == len(dp) {
            dp = append(dp, e[1])
        } else {
            dp[idx] = e[1]
        }
    }
    return len(dp)
}

func binInsertLeft(nums []int, target int) int {
    l, r := 0, len(nums)
    for l < r {
        mid := (l+r) / 2
        if nums[mid] < target {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l
}