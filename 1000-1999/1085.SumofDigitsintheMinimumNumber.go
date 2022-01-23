func sumOfDigits(A []int) int {
    min := A[0]
    for i := 1; i < len(A); i++ {
        if A[i] < min {
            min = A[i]
        }
    }
    sum := 0
    for min > 0 {
        sum += min%10
        min /= 10
    }
    if sum & 1 == 1 {
        return 0
    }
    return 1
}