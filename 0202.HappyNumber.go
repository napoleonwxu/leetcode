func isHappy(n int) bool {
    Map := make(map[int]bool)
    for n != 1 {
        if Map[n] {
            return false
        }
        Map[n] = true
        n = square(n)
    }
    return true
}

func square(n int) int {
    sq := 0
    for n > 0 {
        d := n%10
        sq += d*d
        n /= 10
    }
    return sq
}