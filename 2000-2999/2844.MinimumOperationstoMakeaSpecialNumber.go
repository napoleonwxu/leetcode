func minimumOperations(num string) int {
    n := len(num)
    zero, five := false, false
    for i := n - 1; i >= 0; i-- {
        if zero && (num[i] == '0' || num[i] == '5') {
            return n - i - 2
        }
        if five && (num[i] == '2' || num[i] == '7') {
            return n - i - 2
        }
        if num[i] == '0' {
            zero = true
        }
        if num[i] == '5' {
            five = true
        }
    }
    if zero {
        return n - 1
    }
    return n
}
