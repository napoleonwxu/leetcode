func countOfSubstrings(word string, k int) int {
    ans := 0
    for x := 0; x < len(word); x++ {
        a := 0
        e := 0
        i := 0
        o := 0
        u := 0
        c := 0
        for y := x; y < len(word); y++ {
            if word[y] == 'a' {
                a++
            } else if word[y] == 'e' {
                e++
            } else if word[y] == 'i' {
                i++
            } else if word[y] == 'o' {
                o++
            } else if word[y] == 'u' {
                u++
            } else {
                c++
            }
            if a > 0 && e > 0 && i > 0 && o > 0 && u > 0 && c == k {
                ans++
            }
            if c > k {
                break
            }
        }
    }
    return ans
}