func convertToTitle(columnNumber int) string {
	ans := ""
	for columnNumber > 0 {
		ans = string((columnNumber-1)%26+'A') + ans
		columnNumber = (columnNumber - 1) / 26
	}
	return ans
}