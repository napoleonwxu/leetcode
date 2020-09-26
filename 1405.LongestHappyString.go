func longestDiverseString(a int, b int, c int) string {
    return help(a, b, c, "a", "b", "c")
}

func help(a, b, c int, sa, sb, sc string) string {
    if a < b {
        return help(b, a, c, sb, sa, sc)
    }
    if b < c {
        return help(a, c, b, sa, sc, sb)
    }
    use_a := min(a, 2)
    if b == 0 {
        return strings.Repeat(sa, use_a)
    }
    if a-use_a >= b {
        return strings.Repeat(sa, use_a) + sb + help(a-use_a, b-1, c, sa, sb, sc)
    }
    return strings.Repeat(sa, use_a) + help(a-use_a, b, c, sa, sb, sc)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
