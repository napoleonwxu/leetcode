func distanceBetweenBusStops(distance []int, start int, destination int) int {
    if start > destination {
        start, destination = destination, start
    } 
    left, right := 0, 0
    for i := 0; i < start; i++ {
        left += distance[i]
    }
    for i := start; i < destination; i++ {
        right += distance[i]
    }
    for i := destination; i < len(distance); i++ {
        left += distance[i]
    }
    if left < right {
        return left
    }
    return right
}