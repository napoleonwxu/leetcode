func removeComments(source []string) []string {
    ans := []string{}
    line := []string{}
    //line := ""
    inblock := false
    for _, s := range source {
        if !inblock {
            line = []string{}
            //line = ""
        }
        i, n := 0, len(s)
        for i < n {
            if !inblock && i+1<n && s[i:i+2] == "//" {
                break
            } else if !inblock && i+1<n && s[i:i+2] == "/*" {
                inblock = true
                i += 2
                continue
            } else if inblock && i+1<n && s[i:i+2] == "*/" {
                inblock = false
                i += 2
                continue
            } else if !inblock {
                line = append(line, string(s[i]))
                //line += string(s[i])
            }
            i++
        }
        if !inblock && len(line) > 0 {
            ans = append(ans, strings.Join(line, ""))
            //ans = append(ans, line)
        }
    }
    return ans
}