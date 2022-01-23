func backspaceCompare(S string, T string) bool {
    // O(S+T), O(1)
    i, j := len(S), len(T)
    for i >= 0 || j >= 0 {
        index(&S, &i)
        index(&T, &j)
        if i >= 0 && j >= 0 && S[i] != T[j] {
            return false
        }
    }
    return i == j
    // O(S+T), O(s+t)
    //return backspace(S) == backspace(T)
}

func index(s *string, i *int) {
    *i--
    cnt := 0
    for *i >= 0 {
        if (*s)[*i] == '#' {
            cnt++
        } else if cnt > 0 {
            cnt--
        } else {
            return
        }
        *i--
    }
}

func backspace(s string) string {
    ans := []string{}
    for _, ch := range s {
        if ch != '#' {
            ans = append(ans, string(ch))
        } else if len(ans) > 0 {
            ans = ans[:len(ans)-1]
        }
    }
    return strings.Join(ans, "")
}