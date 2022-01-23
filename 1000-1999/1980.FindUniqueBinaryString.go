func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	ans := make([]byte, n)
	for i, num := range nums {
		ans[i] = '1' - ([]byte(num)[i] - '0')
	}
	return string(ans)
	/*
	   m := make(map[string]bool, n)
	   for _, num := range nums {
	       m[num] = true
	   }
	   for i := int64(0); i < 2<<n; i++ {
	       bin := strconv.FormatInt(i, 2)
	       bin = strings.Repeat("0", n-len(bin)) + bin
	       if !m[bin] {
	           return bin
	       }
	   }
	   return ""
	*/
}