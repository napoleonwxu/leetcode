func bitwiseComplement(N int) int {
    f := 1
    for f < N {
        f = f<<1 + 1
    }
    //return f - N
    return f ^ N
}