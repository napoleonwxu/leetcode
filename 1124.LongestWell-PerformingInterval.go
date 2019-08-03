func longestWPI(hours []int) int {
    longest := 0
    // O(n) + O(n)
    Map := make(map[int]int)
    tire := 0
    for i, hour := range hours {
        if hour > 8 {
            tire++
        } else {
            tire--
        }
        if tire > 0 {
            longest = i + 1
        } else {
            if _, ok := Map[tire]; !ok {
                Map[tire] = i
            }
            if _, ok := Map[tire-1]; ok {
                longest = max(longest, i-Map[tire-1])
            }
        }
    }
    /* O(n^2) + O(n)
    n := len(hours)
    nums := make([]int, n+1)
    for i, hour := range hours {
        if hour > 8 {
            nums[i+1] = nums[i] + 1
        } else {
            nums[i+1] = nums[i] - 1
        }
    }
    for i := 0; i < n; i++ {
        for j := i+1; j <= n; j++ {
            if nums[j] - nums[i] > 0 {
                longest = max(longest, j-i)
            }
        }
    } */
    return longest
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}