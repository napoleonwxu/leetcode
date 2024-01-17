func arrayChange(nums []int, operations [][]int) []int {
    idx := make(map[int]int, len(nums)+len(operations))
    for i, num := range nums {
        idx[num] = i
    }
    for _, op := range operations {
        idx[op[1]] = idx[op[0]]
        nums[idx[op[0]]] = op[1]
    }
    return nums
}