func numJewelsInStones(J string, S string) int {
    Map := make(map[byte]bool, len(J))
    for i := range J {
        Map[J[i]] = true
    }
    cnt := 0
    for i := range S {
        if Map[S[i]] {
            cnt++
        }
    }
    return cnt
}