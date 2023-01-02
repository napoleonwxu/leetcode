func countEven(num int) int {
    tmp, sum := num, 0
    for tmp > 0 {
        sum += tmp % 10
        tmp /= 10
    }
    if sum&1 == 0 {
        return num / 2
    }
    return (num - 1) / 2
}
