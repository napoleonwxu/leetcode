func prisonAfterNDays(cells []int, N int) []int {
    Map := make(map[int]int)
    for ; N > 0; N-- {
        c := toInt(cells)
        if Map[c] > 0 {
            N %= Map[c] - N
            if N == 0 {
                break
            }
        }
        Map[c] = N
        nxt := make([]int, 8)
        for i := 1; i < 7; i++ {
            nxt[i] = 1 - cells[i-1]^cells[i+1]
        }
        cells = nxt
    }
    return cells
}

func toInt(cells []int) int {
    c := 0
    for _, cell := range cells {
        c = (c << 1) + cell
    }
    return c
}