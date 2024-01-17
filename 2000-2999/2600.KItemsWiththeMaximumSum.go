func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
    return min(numOnes, k) - max(0, k-numOnes-numZeros)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}