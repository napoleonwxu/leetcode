func getNoZeroIntegers(n int) []int {
    for i := 1; i <= n/2; i++ {
        if noZero(i) && noZero(n-i) {
            return []int{i, n-i}
        }
    }
    return nil
}

func noZero(n int) bool {
    for n > 0 {
        if n%10 == 0 {
            return false
        }
        n /= 10
    }
    return true
}