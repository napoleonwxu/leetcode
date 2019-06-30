func predictPartyVictory(senate string) string {
    var queue []int
    cnt, ban := []int{0, 0}, []int{0, 0}
    for _, s := range senate {
        x := 0
        if s == 'R' {
            x = 1
        }
        cnt[x]++
        queue = append(queue, x)
    }
    for cnt[0] > 0 && cnt[1] > 0 {
        x := queue[0]
        queue = queue[1:]
        if ban[x] > 0 {
            ban[x]--
            cnt[x]--
        } else {
            ban[x^1]++
            queue = append(queue, x)
        }
    }
    if cnt[1] > 0 {
        return "Radiant"
    } else {
        return "Dire"
    }
}