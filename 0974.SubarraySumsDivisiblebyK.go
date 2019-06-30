func subarraysDivByK(A []int, K int) int {
    Map := make(map[int]int)
    Map[0] = 1
    cnt, sum := 0, 0
    for _, a := range A {
        sum += a
        Map[(sum%K+K)%K]++
    }
    for _, c := range Map {
        cnt += c*(c-1)/2
    }
    return cnt
}