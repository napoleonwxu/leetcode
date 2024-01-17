func findChampion(grid [][]int) int {
    for i, nums := range grid {
        sum := 0
        for _, num := range nums {
            sum += num
        }
        if sum >= len(grid)-1 {
            return i
        }
    }
    return -1
}