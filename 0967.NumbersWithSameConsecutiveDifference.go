func numsSameConsecDiff(N int, K int) []int {
    // BFS
    nums := []int{}
    for i := 1; i < 10; i++ {
        nums = append(nums, i)
    }
    if N == 1 {
        return append(nums, 0)
    }
    for step := 1; step < N; step++ {
        tmp := []int{}
        for _, num := range nums {
            d := num%10
            if d-K >= 0 {
                tmp = append(tmp, num*10 + d-K)
            }
            if K > 0 && d+K < 10 {
                tmp = append(tmp, num*10 + d+K)
            }
        }
        nums = tmp
    }
    return nums
}