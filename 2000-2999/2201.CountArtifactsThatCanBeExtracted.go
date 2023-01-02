func digArtifacts(n int, artifacts [][]int, dig [][]int) int {
    set := make(map[int]bool)
    for _, d := range dig {
        set[d[0]*n+d[1]] = true
    }
    cnt := 0
    for _, art := range artifacts {
        flag := true
        for i := art[0]; i <= art[2]; i++ {
            for j := art[1]; j <= art[3]; j++ {
                if !set[i*n+j] {
                    flag = false
                }
            }
        }
        if flag {
            cnt++
        }
    }
    return cnt
}