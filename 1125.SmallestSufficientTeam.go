func smallestSufficientTeam(req_skills []string, people [][]string) []int {
    n := uint(len(req_skills))
    Map := make(map[string]uint, n)
    for i, skill := range req_skills {
        Map[skill] = uint(i)
    }
    dp := make(map[int][]int)
    dp[0] = []int{}
    for p, skills := range people {
        his_skill := 0
        for _, skill := range skills {
            his_skill |= 1 << Map[skill]
        }
        for skill_set, need := range dp {
            with_him := skill_set | his_skill
            if with_him == skill_set {
                continue
            }
            _, ok := dp[with_him]
            if !ok || len(dp[with_him]) > len(need)+1 {
                tmp := make([]int, len(need))
                copy(tmp, need)
                dp[with_him] = append(tmp, p)
            }
        }
    }
    return dp[1<<n - 1]
}