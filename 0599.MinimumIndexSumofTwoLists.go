func findRestaurant(list1 []string, list2 []string) []string {
    Map := make(map[string]int, len(list1))
    for i, r := range list1 {
        Map[r] = i
    }
    var ans []string
    for j, r := range list2 {
        if _, ok := Map[r]; ok {
            Map[r] += j
            if len(ans) == 0 || Map[r] < Map[ans[0]] {
                ans = []string{r}
            } else if Map[r] == Map[ans[0]] {
                ans = append(ans, r)
            }
        }
    }
    return ans
}