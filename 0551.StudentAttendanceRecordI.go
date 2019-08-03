func checkRecord(s string) bool {
    //return strings.Count(s, "A") <= 1 && !strings.Contains(s, "LLL")
    abs, late := 0, 0
    for _, ch := range s {
        if ch == 'A' {
            abs++
        }
        if ch == 'L' {
            late++
        } else {
            late = 0
        }
        if abs > 1 || late > 2 {
            return false
        }
    }
    return true
}