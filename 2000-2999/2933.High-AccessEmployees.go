func findHighAccessEmployees(access_times [][]string) []string {
    m := make(map[string][]int)
    for _, t := range access_times {
        time, _ := strconv.Atoi(t[1])
        minute := 60*(time/100) + time%100
        m[t[0]] = append(m[t[0]], minute)
    }
    ans := make([]string, 0, len(m))
    for name, minutes := range m {
        sort.Ints(minutes)
        for i := 2; i < len(minutes); i++ {
            if minutes[i]-minutes[i-2] < 60 {
                ans = append(ans, name)
                break
            }
        }
    }
    return ans
}