func minimumBoxes(apple []int, capacity []int) int {
    sum := 0
    for _, cnt := range apple {
        sum += cnt
    }
    sort.Ints(capacity)
    i := len(capacity) - 1
    for ; i >= 0 && sum > 0; i-- {
        sum -= capacity[i]
    }
    return len(capacity) - i - 1
}