func numOfSubarrays(arr []int, k int, threshold int) int {
    sum, cnt := 0, 0
    for i := 0; i < k-1; i++ {
        sum += arr[i]
    }
    for i := k-1; i < len(arr); i++ {
        sum += arr[i]
        if sum/k >= threshold {
            cnt++
        }
        sum -= arr[i-k+1]
    }
    return cnt
}