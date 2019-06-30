func toLowerCase(str string) string {
    //return strings.ToLower(str)
    ch := []string{}
    add := byte('a' - 'A')
    for i := 0; i < len(str); i++ {
        if str[i] >= 'A' && str[i] <= 'Z' {
            ch = append(ch, string(str[i]+add))
        } else {
            ch = append(ch, string(str[i]))
        }
    }
    return strings.Join(ch, "")
}