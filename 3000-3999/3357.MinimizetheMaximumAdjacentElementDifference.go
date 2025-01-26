func minDifference(nums []int) int {
    n := len(nums)
    maxAdj := 0
    minA, maxB := math.MaxInt64, 0
    for i := 1; i < n; i++ {
        if nums[i] > 0 && nums[i-1] > 0 {
            maxAdj = max(maxAdj, abs(nums[i]-nums[i-1]))
        } else if nums[i] > 0 || nums[i-1] > 0 {
            minA = min(minA, max(nums[i], nums[i-1]))
            maxB = max(maxB, max(nums[i], nums[i-1]))
        }
    }
    ans := 0
    minR := (maxB-minA+2) / 3 * 2
    for i := 0; i < n; i++ {
        if (i > 0 && nums[i-1] == -1) || nums[i] > 0 {
            continue
        }
        j := i
        for ; j < n && nums[j] == -1; j++ {}
        a, b := math.MaxInt64, 0
        if i > 0 {
            a = min(a, nums[i-1])
            b = max(b, nums[i-1])
        }
        if j < n {
            a = min(a, nums[j])
            b = max(b, nums[j])
        }
        if j-i == 1 {
            ans = max(ans, min(maxB-a, b-minA))
        } else {
            ans = max(ans, min(maxB-a, min(b-minA, minR)))
        }
    }
    return max(maxAdj, (ans+1)/2)
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}