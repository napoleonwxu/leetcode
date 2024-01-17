func substringXorQueries(s string, queries [][]int) [][]int {
    ans := make([][]int, len(queries))
    // m: len(s), n: len(queries)
    // O(30m) + O(n)
    m := make(map[int][]int)
    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == '0' {
            m[0] = []int{i, i}
            continue
        }
        num := 0
        for j := i; j < len(s); j++ {
            num = num<<1 + int(s[j]-'0')
            if num > 1<<32 {
                break
            }
            m[num] = []int{i, j}
        }
    }
    for i, q := range queries {
        ans[i] = []int{-1, -1}
        if v, ok := m[q[0]^q[1]]; ok {
            ans[i] = v
        }
    }
    /* O(30mn)
    for i, q := range queries {
        bin := ""
        for x := q[0]^q[1]; x > 0; x >>= 1 {
            if x&1 == 1 {
                bin = "1" + bin
            } else {
                bin = "0" + bin
            }
        }
        if bin == "" {
            bin = "0"
        }
        idx := strings.Index(s, bin)
        if idx == -1 {
            ans[i] = []int{-1, -1}
        } else {
            ans[i] = []int{idx, idx+len(bin)-1}
        }
    }*/
    return ans
}