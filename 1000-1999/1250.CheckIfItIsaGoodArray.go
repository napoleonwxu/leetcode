func isGoodArray(nums []int) bool {
    g := nums[0]
    for _, num := range nums {
        g = gcd(g, num)
        if g == 1 {
            return true
        }
    }
    return false
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}