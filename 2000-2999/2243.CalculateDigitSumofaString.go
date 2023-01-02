func digitSum(s string, k int) string {
    for len(s) > k {
        tmp := ""
        for i := 0; i < len(s); i += k {
            sum := 0
            for j := i; j < len(s) && j < i+k; j++ {
                sum += int(s[j] - '0')
            }
            tmp += strconv.Itoa(sum)
        }
        s = tmp
    }
    return s
}