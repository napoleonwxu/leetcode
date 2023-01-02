func largestInteger(num int) int {
    bytes := []byte(strconv.Itoa(num))
    odds, evens := make([]byte, 0, len(bytes)), make([]byte, 0, len(bytes))
    for _, b := range bytes {
        if b&1 == 1 {
            odds = append(odds, b)
        } else {
            evens = append(evens, b)
        }
    }
    sortBytes(odds)
    sortBytes(evens)
    i, j := 0, 0
    for k, b := range bytes {
        if b&1 == 1 {
            bytes[k] = odds[i]
            i++
        } else {
            bytes[k] = evens[j]
            j++
        }
    }
    ans, _ := strconv.Atoi(string(bytes))
    return ans
}

func sortBytes(bytes []byte) []byte {
    sort.Slice(bytes, func(i, j int) bool {
        return bytes[i] > bytes[j]
    })
    return bytes
}