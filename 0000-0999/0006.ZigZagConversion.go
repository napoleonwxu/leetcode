func convert(s string, numRows int) string {
    n := len(s)
    if numRows == 1 || n <= numRows {
        return s
    }
    chs := make([][]byte, numRows)
    i := 0
    for i < n {
        for j := 0; j < numRows && i < n; i, j = i+1, j+1 {
            chs[j] = append(chs[j], s[i])
        }
        for j := numRows-2; j > 0 && i < n; i, j = i+1, j-1 {
            chs[j] = append(chs[j], s[i])
        }
    }
    for j := 1; j < numRows; j++ {
        chs[0] = append(chs[0], chs[j]...)
    }
    return string(chs[0])
}