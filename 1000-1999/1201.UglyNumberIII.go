func nthUglyNumber(n int, a int, b int, c int) int {
    // https://www.geeksforgeeks.org/find-the-nth-term-divisible-by-a-or-b-or-c/
    ab, ac, bc := lcm(a, b), lcm(a, c), lcm(b, c)
    abc := lcm(a, bc)
    left, right := 1, 2*int(1e9)
    for left < right {
        mid := left + (right-left)/2
        cnt := mid/a + mid/b + mid/c - mid/ab - mid/ac - mid/bc + mid/abc
        if cnt < n {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}

func lcm(x, y int) int {
    return x*y/gcd(x, y)
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}