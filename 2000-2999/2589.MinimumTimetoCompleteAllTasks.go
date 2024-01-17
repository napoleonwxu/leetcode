import "sort"

func findMinimumTime(tasks [][]int) int {
    time := make([]bool, 2001)
    sort.Slice(tasks, func(i, j int) bool {
        return tasks[i][1] < tasks[j][1]
    })
    for _, task := range tasks {
        for t := task[0]; t <= task[1]; t++ {
            if time[t] {
                task[2]--
            }
        }
        for t := task[1]; task[2] > 0; t-- {
            if !time[t] {
                time[t] = true
                task[2]--
            }
        }
    }
    ans := 0
    for _, t := range time {
        if t {
            ans++
        }
    }
    return ans
}