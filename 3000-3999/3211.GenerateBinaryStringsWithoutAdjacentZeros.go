func validStrings(n int) []string {
    tail0, tail1 := []string{"0"}, []string{"1"}
    for ; n > 1; n-- {
        tail10 := make([]string, 0, len(tail1))
        for _, s := range tail1 {
            tail10 = append(tail10, s+"0")
        }
        tailx1 := make([]string, 0, len(tail0)+len(tail1))
        for _, s := range tail0 {
            tailx1 = append(tailx1, s+"1")
        }
        for _, s := range tail1 {
            tailx1 = append(tailx1, s+"1")
        }
        tail0, tail1 = tail10, tailx1
    }
    return append(tail0, tail1...)
}