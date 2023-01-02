func convertTime(current string, correct string) int {
    cnt, diff := 0, getMinutes(correct) - getMinutes(current)
    ops := []int{60, 15, 5, 1}
    for _, op := range ops {
        cnt += diff/op
        diff %= op
    }
    return cnt
}

func getMinutes(t string) int {
    hm := strings.Split(t, ":")
    h, _ := strconv.Atoi(hm[0])
    m, _ := strconv.Atoi(hm[1])
    return 60*h + m
}