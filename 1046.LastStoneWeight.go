func lastStoneWeight(stones []int) int {
    for len(stones) >= 2 {
        sort.Sort(sort.Reverse(sort.IntSlice(stones)))
        for len(stones) >= 2 && stones[0] == stones[1] {
            stones = stones[2:]
        }
        if len(stones) >= 2 {
            stones[1] = stones[0] - stones[1]
            stones = stones[1:]
        }
    }
    if len(stones) == 0 {
        return 0
    }
    return stones[0]
}