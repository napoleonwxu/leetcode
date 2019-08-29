func romanToInt(s string) int {
    ans, n := 0, len(s)
    Map := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
    for i := 0; i < n-1; i++ {
        if Map[s[i]] < Map[s[i+1]] {
            ans -= Map[s[i]]
        } else {
            ans += Map[s[i]]
        }
    }
    ans += Map[s[n-1]]
    /*
    Map := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
    for i := 0; i < n; i++ {
        if i+1 < n && Map[s[i:i+2]] != 0 {
            ans += Map[s[i:i+2]]
            i++
        } else {
            ans += Map[string(s[i])]
        }
    } */
    return ans
}