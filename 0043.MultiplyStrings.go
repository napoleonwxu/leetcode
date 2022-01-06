func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	multi := make([]byte, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			multi[i+j+1] += (num1[i] - '0') * (num2[j] - '0')
			multi[i+j] += multi[i+j+1] / 10
			multi[i+j+1] %= 10
		}
	}
	idx := 0
	for ; idx < m+n-1 && multi[idx] == 0; idx++ {
	}
	for i := idx; i < m+n; i++ {
		multi[i] += '0'
	}
	return string(multi[idx:])
}