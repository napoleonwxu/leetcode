func findLengthOfShortestSubarray(arr []int) int {
    left, n := 0, len(arr)
    for ; left < n-1 && arr[left] <= arr[left+1]; left++ {}
    if left >= n-1 {
        return 0
    }
    right := n-1
    for ; right > left && arr[right-1] <= arr[right]; right-- {}
    ans := min(right, n-left-1)
    i, j := 0, right
    for i <= left && j < n {
        if arr[i] <= arr[j] {
            ans = min(ans, j-i-1)
            i++
        } else {
            j++
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}