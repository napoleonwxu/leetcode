func maxProduct(nums []int) int {
    ans := math.MinInt32
    mi, ma := 1, 1
    for _, num := range nums {
        if num < 0 {
            mi, ma = ma, mi
        }
        mi = min(mi*num, num)
        ma = max(ma*num, num)
        ans = max(ans, ma)
    }
    /*
    n := len(nums)
    l, r := 0, 0
    for i := 0; i < n; i++ {
        if l != 0 {
            l *= nums[i]
        } else {
            l = nums[i]
        }
        if r != 0 {
            r *= nums[n-1-i]
        } else {
            r = nums[n-1-i]
        }
        ans = max(ans, max(l, r))
    }*/
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