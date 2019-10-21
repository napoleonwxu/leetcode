func multiply(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    num := make([]byte, m+n)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            num[i+j+1] += (num1[i]-'0')*(num2[j]-'0')
            num[i+j] += num[i+j+1]/10
            num[i+j+1] %= 10
        }
    }
    i := 0
    for ; i < m+n-1 && num[i] == 0; i++ {}
    for j := i; j < m+n; j++ {
        num[j] += '0'
    }
    return string(num[i:])
}