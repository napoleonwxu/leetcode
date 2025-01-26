func multiply(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    mul := make([]byte, m+n)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            mul[i+j+1] += (num1[i]-'0') * (num2[j]-'0')
            mul[i+j] += mul[i+j+1] / 10
            mul[i+j+1] %= 10
        }
    }
    start := 0
    for ; start < m+n-1 && mul[start] == 0; start++ {}
    for i := start; i < m+n; i++ {
        mul[i] += '0'
    }
    return string(mul[start:])
}