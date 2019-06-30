func reorderLogFiles(logs []string) []string {
    sort.SliceStable(logs, func (i, j int) bool {
        split1 := strings.SplitN(logs[i], " ", 2)
        split2 := strings.SplitN(logs[j], " ", 2)
        s1 := "0" + split1[1]
        s2 := "0" + split2[1]
        digit1 := unicode.IsNumber(rune(s1[1]))
        digit2 := unicode.IsNumber(rune(s2[1]))
        if !digit1 && !digit2 && s1 == s2 {
            return split1[0] < split2[0]
        }
        if digit1 {
            s1 = "1"
        }
        if digit2 {
            s2 = "1"
        }
        return s1 < s2
    })
    return logs
}