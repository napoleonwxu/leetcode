func reverse(x int) int {
    var rev int32
    for x != 0 {
        nxt := 10*rev + int32(x%10)
        if (nxt - int32(x%10))/10 != rev {
            return 0
        }
        rev = nxt
        x /= 10
    }
    return int(rev)
}