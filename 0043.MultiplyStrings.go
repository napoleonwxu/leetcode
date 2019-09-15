func multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    m, n := len(num1), len(num2)
    mul := make([]int, m+n)
    for i := m-1; i >= 0; i-- {
        for j := n-1; j >= 0; j-- {
            mul[i+j+1] += int((num1[i]-'0')*(num2[j]-'0'))
            mul[i+j] += mul[i+j+1] / 10
            mul[i+j+1] %= 10
        }
	}
	// len(ans) must be len(num1)+len(num2) or len(num1)+len(num2)-1
    if mul[0] == 0 {
        mul = mul[1:]
    }
    ans := make([]byte, len(mul))
    for i := 0; i < len(mul); i++ {
        ans[i] = '0' + byte(mul[i])
    }
    return string(ans)
}