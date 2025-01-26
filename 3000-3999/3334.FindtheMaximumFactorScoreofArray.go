func maxScore(nums []int) int64 {
    ans := int64(0)
    for i := -1; i < len(nums); i++ {
        l, g := 1, 0
        for j := 0; j < len(nums); j++ {
            if j == i {
                continue
            }
            l = lcm(l, nums[j])
            g = gcd(g, nums[j])
        }
        ans = max(ans, int64(l)*int64(g))
    }
    return ans
}

func lcm(x, y int) int {
    return x * y / gcd(x, y)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}
