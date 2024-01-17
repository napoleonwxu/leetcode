func fillCups(amount []int) int {
    ma, sum := 0, 0
    for _, a := range amount {
        ma = max(ma, a)
        sum += a
    }
    return max(ma, (sum+1)/2)
    /*
    cnt := 0
    for {
        sort.Ints(amount)
        if amount[1] >= 1 {
            amount[1]--
            amount[2]--
            cnt++
        } else {
            cnt += amount[2]
            break
        }
    }
    return cnt
    */
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}