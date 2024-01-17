func countSymmetricIntegers(low int, high int) int {
    cnt := 0
    for i := low; i <= high; i++ {
        if isSymmetric(i) {
            cnt++
        }
    }
    return cnt
}

func isSymmetric(num int) bool {
    n, sum := 0, 0
    for tmp := num; tmp > 0; tmp /= 10 {
        sum += tmp % 10
        n++
    }
    if n%2 == 1 {
        return false
    }
    half := 0
    for n /= 2; n > 0; n-- {
        half += num % 10
        num /= 10
    }
    return 2*half == sum
}