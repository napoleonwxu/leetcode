import "math"

func pivotInteger(n int) int {
    sum := (n*n + n) / 2
    sq := int(math.Sqrt(float64(sum)))
    if sq*sq == sum {
        return sq
    }
    return -1
}