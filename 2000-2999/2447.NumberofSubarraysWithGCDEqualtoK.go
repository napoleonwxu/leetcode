func subarrayGCD(nums []int, k int) int {
    n, cnt := len(nums), 0
    for i := 0; i < n; i++ {
        g := nums[i]
        for j := i; j < n && nums[j]%k == 0; j++ {
            g = gcd(g, nums[j])
            if g == k {
                cnt++
            }
        }
    }
    return cnt
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}