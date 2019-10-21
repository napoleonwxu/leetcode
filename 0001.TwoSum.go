func twoSum(nums []int, target int) []int {
    idx := make(map[int]int)
    for i, num := range nums {
        if j, ok := idx[target-num]; ok {
            return []int{j, i}
        }
        idx[num] = i
    }
    return nil
}