func repeatedSubstringPattern(s string) bool {
    // amazing !
    s2 := strings.Repeat(s, 2)
    return strings.Contains(s2[1:len(s2)-1], s)
    /*
    n := len(s)
    if n <= 1 {
        return false
    }
    sq := int(math.Sqrt(float64(n)))
    ds := []int{1}
    for i := 2; i <= sq; i++ {
        if n%i == 0 {
            ds = append(ds, i, n/i)
        }
    }
    for _, d := range ds {
        if strings.Repeat(s[:d], n/d) == s {
            return true
        }
    }
    return false
    */
}
