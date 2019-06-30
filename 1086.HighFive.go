func highFive(items [][]int) [][]int {
    sort.Slice(items, func(i, j int) bool {
        if items[i][0] == items[j][0] {
            return items[i][1] > items[j][1]
        }
        return items[i][0] < items[j][0]
    })
    n := len(items)
    ans := [][]int{}
    for i := 0; i < n; {
        id, sum := items[i][0], 0
        for j := 0; j < 5; j++ {
            sum += items[i][1]
            i++
        }
        ans = append(ans, []int{id, sum/5})
        for ; i < n && items[i][0] == id; i++ {}
    }
    return ans
}