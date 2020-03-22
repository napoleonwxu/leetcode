var sub map[int][]int

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    sub = make(map[int][]int)
    for employee, boss := range manager {
        sub[boss] = append(sub[boss], employee)
    }
    return dfs(headID, informTime)
}

func dfs(id int, informTime []int) int {
    time := 0
    for _, sub_id := range sub[id] {
        time = max(time, dfs(sub_id, informTime))
    }
    return time + informTime[id]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}