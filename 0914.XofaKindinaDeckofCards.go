// O(N*logN*logN)
func hasGroupsSizeX(deck []int) bool {
    Map := make(map[int]int)
    for _, n := range deck {
        Map[n]++
    }
    g := 0
    for _, cnt := range Map {
        g = gcd(cnt, g)
        if g <= 1 {
            return false
        }
    }
    return true
}

func gcd(x int, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}