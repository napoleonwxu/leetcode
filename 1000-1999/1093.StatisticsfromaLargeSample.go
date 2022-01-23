func sampleStats(count []int) []float64 {
    n := 256
    ans := make([]float64, 5)
    cnt, sum := 0, 0
    mode, max := 0, 0
    i, j := 0, n-1
    for ; i < n && count[i] == 0; i++ {}
    for ; j >= 0 && count[j] == 0; j-- {}
    ans[0] = float64(i)
    ans[1] = float64(j)
    for k := i; k <= j; k++ {
        cnt += count[k]
        sum += k * count[k]
        if count[k] > max {
            max = count[k]
            mode = k
        }
    }
    ans[2] = float64(sum)/float64(cnt)
    ans[4] = float64(mode)
    m1, m2 := cnt >> 1, cnt >> 1
    if cnt&1 == 0 {
        m2++
    }
    cnt = 0
    for k := i; k <= j; k++ {
        if cnt < m1 && cnt+count[k] >= m1 {
            ans[3] += float64(k)/2.0
        }
        if cnt < m2 && cnt+count[k] >= m2 {
            ans[3] += float64(k)/2.0
            break
        }
        cnt += count[k]
    }
    return ans
}