func fullJustify(words []string, maxWidth int) []string {
    ans := []string{}
    line, width := []string{}, 0
    for _, word := range words {
        n := len(line)
        if width + len(word) + n <= maxWidth {
            line = append(line, word)
            width += len(word)
        } else {
            if n == 1 {
                ans = append(ans, line[0]+strings.Repeat(" ", maxWidth-width))
            } else {
                space, extra := (maxWidth-width)/(n-1), (maxWidth-width)%(n-1)
                for i := 0; i < extra; i++ {
                    line[i] += " "
                }
                ans = append(ans, strings.Join(line, strings.Repeat(" ", space)))
            }
            line, width = []string{word}, len(word)
        }
    }
    last := strings.Join(line, " ")
    ans = append(ans, last+strings.Repeat(" ", maxWidth-len(last)))
    return ans
}