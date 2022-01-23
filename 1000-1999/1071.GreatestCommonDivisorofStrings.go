func gcdOfStrings(str1 string, str2 string) string {
    //gcd
    d := gcd(len(str1), len(str2))
    if str1[:d] == str2[:d] {
        return str1[:d]
    }
    return ""
    /* recursive
    if len(str1) == len(str2) {
        if str1 == str2 {
            return str1
        } else {
            return ""
        }
    }
    if len(str1) < len(str2) {
        str1, str2 = str2, str1
    }
    if str1[:len(str2)] == str2 {
        return gcdOfStrings(str1[len(str2):], str2)
    } else {
        return ""
    }
    */
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}