func isNStraightHand(hand []int, W int) bool {
    if len(hand)%W != 0 {
        return false
    }
    Map := make(map[int]int)
    for _, card := range hand {
        Map[card]++
    }
    sort.Ints(hand)
    for _, card := range hand {
        if Map[card] == 0 {
            continue
        }
        for i := 0; i < W; i++ {
            Map[card+i] -= 1
            if Map[card+i] < 0 {
                return false
            }
        }
    }
    return true
}