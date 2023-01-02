func minimumCardPickup(cards []int) int {
    ans := len(cards) + 1
    idx := make(map[int]int)
    for i, card := range cards {
        if pre, ok := idx[card]; ok {
            ans = min(ans, i-pre+1)
        }
        idx[card] = i
    }
    if ans == len(cards)+1 {
        return -1
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}