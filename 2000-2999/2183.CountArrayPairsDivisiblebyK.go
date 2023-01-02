func countPairs(nums []int, k int) int64 {
    var cnt int64
    gcdMap := make(map[int]int)
    for _, num := range nums {
        gcd1 := gcd(num, k)
        for gcd2, c := range gcdMap {
            if gcd1 * gcd2 % k == 0 {
                cnt += int64(c)
            }
        }
        gcdMap[gcd1]++
    }
    return cnt
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}
