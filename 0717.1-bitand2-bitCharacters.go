func isOneBitCharacter(bits []int) bool {
    i, n := 0, len(bits)
    for i < n-1 {
        i += bits[i] + 1
        /*
        if bits[i] == 1 {
            i += 2
        } else {
            i++
        } */
    }
    return i == n-1
}