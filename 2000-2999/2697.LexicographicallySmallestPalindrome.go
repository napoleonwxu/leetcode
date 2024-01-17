func makeSmallestPalindrome(s string) string {
    n := len(s)
    bytes := []byte(s)
    for i := 0; i < n/2; i++ {
        if bytes[i] != bytes[n-i-1] {
            bytes[i] = min(bytes[i], bytes[n-i-1])
            bytes[n-i-1] = bytes[i]
        }
    }
    return string(bytes)
}

func min(x, y byte) byte {
    if x < y {
        return x
    }
    return y
}