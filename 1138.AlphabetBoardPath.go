func alphabetBoardPath(target string) string {
    path := []string{}
    x0, y0 := 0, 0
    for _, ch := range target {
        d := int(ch - 'a')
        x, y := d/5, d%5
        // L U before R D to deal with 'z'
        if x < x0 {
            path = append(path, strings.Repeat("U", x0-x))
        }
        if y < y0 {
            path = append(path, strings.Repeat("L", y0-y))
        }
        if x > x0 {
            path = append(path, strings.Repeat("D", x-x0))
        }
        if y > y0 {
            path = append(path, strings.Repeat("R", y-y0))
        }
        path = append(path, "!")
        x0, y0 = x, y
    }
    /*
    n := len(target)
    pos := [][]int{[]int{0, 0, 0}}
    for i := 0; i < n; i++ {
        d := int(target[i] - 'a')
        if d != 25 || (i > 0 && target[i-1] == 'z') {
            pos = append(pos, []int{d/5, d%5, 1})
        } else {
            pos = append(pos, []int{4, 0, 0})
            pos = append(pos, []int{5, 0, 1})
            if i != n-1 && target[i+1] != 'z' { 
                pos = append(pos, []int{4, 0, 0})
            }
        }
    }
    for i := 1; i < len(pos); i++ {
        if pos[i][0] > pos[i-1][0] {
            path = append(path, strings.Repeat("D", pos[i][0]-pos[i-1][0]))
        } else {
            path = append(path, strings.Repeat("U", pos[i-1][0]-pos[i][0]))
        }
        if pos[i][1] > pos[i-1][1] {
            path = append(path, strings.Repeat("R", pos[i][1]-pos[i-1][1]))
        } else {
            path = append(path, strings.Repeat("L", pos[i-1][1]-pos[i][1]))
        }
        if pos[i][2] == 1 {
            path = append(path, "!")
        }
    } */
    return strings.Join(path, "")
}