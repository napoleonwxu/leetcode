func smallestEvenMultiple(n int) int {
    return n << (n & 1)
}