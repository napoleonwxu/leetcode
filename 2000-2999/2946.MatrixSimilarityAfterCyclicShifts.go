func areSimilar(mat [][]int, k int) bool {
    n := len(mat[0])
    for _, row := range mat {
        for j, num := range row {
            if num != row[(j+k)%n] {
                return false
            }
        }
    }
    return true
}