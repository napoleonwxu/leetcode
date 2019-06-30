/*
type Costs [][]int

func (costs Costs) Len() int {
    return len(costs)
}

func (costs Costs) Less(i, j int) bool {
    return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
}

func (costs Costs) Swap(i, j int) {
    costs[i], costs[j] = costs[j], costs[i]
}
*/
func twoCitySchedCost(costs [][]int) int {
    ans, n := 0, len(costs)>>1
    //sort.Sort(Costs(costs))
    sort.Slice(costs, func(i, j int) bool {
        return costs[i][0]-costs[i][1] < costs[j][0]-costs[j][1]
    })
    for i := 0; i < n; i++ {
        ans += costs[i][0] + costs[i+n][1]
    }
    return ans
}