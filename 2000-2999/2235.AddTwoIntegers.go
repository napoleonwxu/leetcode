func sum(num1 int, num2 int) int {
    if num1 == 0 {
        return num2
    }
    return sum((num1 & num2) << 1, num1 ^ num2)
}