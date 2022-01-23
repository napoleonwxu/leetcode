func checkValidString(s string) bool {
    // O(n)
    n := len(s)
    l, r := 0, 0
    for i := 0; i < n; i++ {
        if s[i] == ')' {
            l--
        } else {
            l++
        }
        if s[n-i-1] == '(' {
            r--
        } else {
            r++
        }
        if l < 0 || r < 0 {
            return false
        }
    }
    return true
    /* O(n)
    lo, hi := 0, 0
    for _, ch := range s {
        if ch == '(' {
            lo++
            hi++
        } else if ch == ')' {
            lo = max(lo-1, 0)
            hi--
        } else {
            lo = max(lo-1, 0)
            hi++
        }
        if hi < 0 {
            return false
        }
    }
    return lo == 0
    */
    // worst case: O(3^n)
    //return dfs(s, 0, 0)
}

func dfs(s string, i, cnt int) bool {
    if cnt < 0 {
        return false
    }
    if i >= len(s) {
        return cnt == 0
    }
    if s[i] == '(' {
        return dfs(s, i+1, cnt+1)
    } else if s[i] == ')' {
        return dfs(s, i+1, cnt-1)
    }
    return dfs(s, i+1, cnt-1) || dfs(s, i+1, cnt) || dfs(s, i+1, cnt+1)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}