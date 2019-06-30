func twoSum(nums []int, target int) []int {
    Map := make(map[int]int)
    for i, v := range nums {
        r, ok := Map[target-v]
        if ok {
            return []int{r, i}
        }
        Map[v] = i
    }
    return nil
}