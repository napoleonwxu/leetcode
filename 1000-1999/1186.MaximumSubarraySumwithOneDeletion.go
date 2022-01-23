func maximumSum(arr []int) int {
    n := len(arr)
    maxL, maxR := make([]int, n), make([]int, n)
    maxL[0], maxR[n-1] = arr[0], arr[n-1]
    Max := arr[0]
    for i := 1; i < n; i++ {
        maxL[i] = max(arr[i], arr[i]+maxL[i-1])
        Max = max(Max, maxL[i])
    }
    for i := n-2; i >= 0; i-- {
        maxR[i] = max(arr[i], arr[i]+maxR[i+1])
    }
    for i := 1; i < n-1; i++ {
        Max = max(Max, maxL[i-1]+maxR[i+1])
    }
    return Max
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}