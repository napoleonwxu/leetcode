func isBalanced(num string) bool {
    even, odd := 0, 0
    for i, ch := range num {
        if i%2 == 0 {
            even += int(ch - '0')
        } else {
            odd += int(ch - '0')
        }
    }
    return even == odd
}