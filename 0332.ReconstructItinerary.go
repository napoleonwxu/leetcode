func findItinerary(tickets [][]string) []string {
    Map := make(map[string][]string)
    for _, t := range tickets {
        Map[t[0]] = append(Map[t[0]], t[1])
    }
    for _, to := range Map {
        sort.Strings(to)
    }
    path := make([]string, 0, len(tickets)+1)
    stack := make([]string, len(tickets)+1)
    stack[0] = "JFK"
    n := 1
    for n > 0 {
        for len(Map[stack[n-1]]) > 0 {
            cur := stack[n-1]
            stack[n] = Map[cur][0]
            n++
            Map[cur] = Map[cur][1:]
        }
        path = append([]string{stack[n-1]}, path...)
        n--
    }
    return path
}