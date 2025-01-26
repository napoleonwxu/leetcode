func minZeroArray(nums []int, queries [][]int) int {
    n := len(nums)
    prefix := make([]int, n+1)
    sum, k := 0, 0
    for i, num := range nums {
        for sum + prefix[i] < num {
            k++
            if k-1 >= len(queries) {
                return -1
            }
            l, r := queries[k-1][0], queries[k-1][1]
            val := queries[k-1][2]
            if i > r {
                continue
            }
            prefix[max(i, l)] += val
            prefix[r+1] -= val
        }
        sum += prefix[i]
    }
    return k
}
