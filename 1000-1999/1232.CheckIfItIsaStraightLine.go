func checkStraightLine(coordinates [][]int) bool {
    n := len(coordinates)
    if n <= 2 {
        return true
    }
    a, b := coordinates[0], coordinates[1]
    for _, c := range coordinates {
        if (c[0]-a[0])*(c[1]-b[1]) != (c[0]-b[0])*(c[1]-a[1]) {
            return false
        }
    }
    /*
    dx, dy := coordinates[1][0]-coordinates[0][0], coordinates[1][1]-coordinates[0][1]
    d := gcd(dx, dy)
    dx, dy = dx/d, dy/d
    for i := 2; i < n; i++ {
        ix, iy := coordinates[i][0]-coordinates[0][0], coordinates[i][1]-coordinates[0][1]
        d = gcd(ix, iy)
        ix, iy = ix/d, iy/d
        if ix != dx || iy != dy {
            return false
        }
    } */
    return true
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}