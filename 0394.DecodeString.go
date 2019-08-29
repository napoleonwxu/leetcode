func decodeString(s string) string {
    stack_n, stack_s := []int{}, []string{}
    num, cur := 0, ""
    for _, ch := range s {
        if ch >= '0' && ch <= '9' {
            num = 10*num + int(ch-'0')
        } else if ch == '[' {
            stack_n = append(stack_n, num)
            stack_s = append(stack_s, cur)
            num, cur = 0, ""
        } else if ch == ']' {
            pre := stack_s[len(stack_s)-1]
            stack_s = stack_s[:len(stack_s)-1]
            rep := stack_n[len(stack_n)-1]
            stack_n = stack_n[:len(stack_n)-1]
            cur = pre + strings.Repeat(cur, rep)
        } else {
            cur += string(ch)
        }
    }
    return cur
}