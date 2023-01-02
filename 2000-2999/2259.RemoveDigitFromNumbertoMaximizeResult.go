func removeDigit(number string, digit byte) string {
    n, idx := len(number), -1
    for i := range number {
        if number[i] == digit {
            if i < n-1 && number[i] < number[i+1] {
                return number[:i] + number[i+1:]
            }
            idx = i
        }
    }
    return number[:idx] + number[idx+1:]
}