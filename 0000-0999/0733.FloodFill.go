func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc] != newColor {
        dfs(image, sr, sc, image[sr][sc], newColor)
    }
    return image
}

func dfs(image [][]int, i, j, old, color int) {
    if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) || image[i][j] != old {
        return
    }
    image[i][j] = color
    dfs(image, i-1, j, old, color)
    dfs(image, i+1, j, old, color)
    dfs(image, i, j-1, old, color)
    dfs(image, i, j+1, old, color)
}