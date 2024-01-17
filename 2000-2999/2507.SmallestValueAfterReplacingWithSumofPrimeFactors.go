func smallestValue(n int) int {
    for true {
        sum := sumOfPrimeFactors(n)
        if sum == n {
            break
        }
        n = sum
    }
    return n
}

func sumOfPrimeFactors(n int) int {
    div, sum := 2, 0
    for n > 1 {
        if n%div == 0 {
            n /= div
            sum += div
        } else {
            div++
        }
    }
    return sum
}