var idx []int

func smallestStringWithSwaps(s string, pairs [][]int) string {
    // Union-Find
    n := len(s)
    idx = make([]int, n)
    for i := 0; i < n; i++ {
        idx[i] = i
    }
    for _, pair := range pairs {
        x, y := find(pair[0]), find(pair[1])
        if x != y {
            idx[x] = y
        }
    }
    
    Map := make(map[int][]byte)
    for i := 0; i < n; i++ {
        x := find(i)
        Map[x] = append(Map[x], s[i])
    }
    for _, chars := range Map {
        sort.Slice(chars, func(i, j int) bool {
            return chars[i] < chars[j]
        })
    }
    
    ans := make([]byte, n)
    for i := 0; i < n; i++ {
        x := find(i)
        ans[i] = Map[x][0]
        Map[x] = Map[x][1:]
    }
    return string(ans)
}

func find(x int) int {
    if x != idx[x] {
        idx[x] = find(idx[x])
    }
    return idx[x]
}