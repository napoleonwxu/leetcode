func queensAttacktheKing(queens [][]int, king []int) [][]int {
    Map := make(map[int]bool)
    for _, queen := range queens {
        Map[queen[0]*8+queen[1]] = true
    }
    ans := make([][]int, 0, 8)
    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue
            }
            for i, j := king[0], king[1]; i >= 0 && i < 8 && j >= 0 && j < 8; i, j = i+dx, j+dy {
                if Map[i*8+j] {
                    ans = append(ans, []int{i, j})
                    break
                }
            }
        }
    }
    return ans
}