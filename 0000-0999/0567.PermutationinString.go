func checkInclusion(s1 string, s2 string) bool {
    len1, len2 := len(s1), len(s2)
    if len1 > len2 {
        return false
    }
    cnts := make([]int, 26)
    for i := 0; i < len1; i++ {
        cnts[s1[i]-'a']++
        cnts[s2[i]-'a']--
    }
    if allZero(cnts) {
        return true
    }
    for i := len1; i < len2; i++ {
        cnts[s2[i-len1]-'a']++
        cnts[s2[i]-'a']--
        if allZero(cnts) {
            return true
        }
    }
    return false
}

func allZero(nums []int) bool {
    for _, num := range nums {
        if num != 0 {
            return false
        }
    }
    return true
}