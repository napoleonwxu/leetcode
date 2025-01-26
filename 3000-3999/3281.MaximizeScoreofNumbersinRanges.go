func maxPossibleScore(start []int, d int) int {
    sort.Ints(start)
    //ans := 0
    l, r := 0, start[len(start)-1]-start[0]+d
    for l <= r {
        mid := (l + r) / 2
        if check(start, mid, d) {
            //ans = mid
            l = mid + 1
        } else {
            r = mid - 1
        }
    }
    return l - 1
}

func check(start []int, mid, d int) bool {
    pre := start[0]
    for i := 1; i < len(start); i++ {
        if pre+mid > start[i]+d {
            return false
        }
        pre = max(pre+mid, start[i])
    }
    return true
}