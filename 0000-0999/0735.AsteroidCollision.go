// Python solution is better to understand
func asteroidCollision(asteroids []int) []int {
    var ans []int
    i := -1
    for _, n := range asteroids {
        for i >= 0 && ans[i] > 0 && n < 0 {
            if ans[i] < -n {
                ans = ans[:i]
                i--
                continue
            } else if ans[i] == -n {
                ans = ans[:i]
                i--
            }
            goto next
        }
        ans = append(ans, n)
        i++
        next:
    }
    return ans
}