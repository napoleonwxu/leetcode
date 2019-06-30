func distributeCandies(candies []int) int {
    Map := make(map[int]bool)
    for _, candy := range candies {
        Map[candy] = true
    }
    if len(Map) >= len(candies)>>1 {
        return len(candies)>>1
    }
    return len(Map)
}