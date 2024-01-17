func maximumStrongPairXor(nums []int) int {
    ans := 0
    for _, x := range nums {
        for _, y := range nums {
            if abs(x-y) <= min(x, y) && x^y > ans {
                ans = x ^ y
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}