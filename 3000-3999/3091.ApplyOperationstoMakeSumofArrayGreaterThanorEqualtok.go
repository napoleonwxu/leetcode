func minOperations(k int) int {
    sq := int(math.Sqrt(float64(k)))
    return sq - 1 + (k-1)/sq
}