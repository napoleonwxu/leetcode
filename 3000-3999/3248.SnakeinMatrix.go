func finalPositionOfSnake(n int, commands []string) int {
    i, j := 0, 0
    for _, cmd := range commands {
        if cmd == "UP" {
            i--
        } else if cmd == "DOWN" {
            i++
        } else if cmd == "LEFT" {
            j--
        } else if cmd == "RIGHT" {
            j++
        }
    }
    return i*n + j
}