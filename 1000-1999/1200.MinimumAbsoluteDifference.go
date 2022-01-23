func minimumAbsDifference(arr []int) [][]int {
    sort.Ints(arr)
    diff := arr[1] - arr[0]
    for i := 2; i < len(arr); i++ {
        diff = min(diff, arr[i]-arr[i-1])
    }
    ans := [][]int{}
    for i := 1; i < len(arr); i++ {
        if arr[i]-arr[i-1] == diff {
            ans = append(ans, []int{arr[i-1], arr[i]})
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