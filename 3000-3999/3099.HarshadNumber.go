func sumOfTheDigitsOfHarshadNumber(x int) int {
    sum := 0
    for tmp := x; tmp > 0; tmp /= 10 {
        sum += tmp % 10
    }
    if x%sum == 0 {
        return sum
    }
    return -1
}