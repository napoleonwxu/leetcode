func minOperations(n int) int {
	ans := 0
	for n > 0 {
		if n&3 == 3 {
			ans++
			n++
		} else {
			ans += n & 1
			n >>= 1
		}
	}
	return ans
	/*
	   tmp := 1
	   for tmp < n {
	       tmp <<= 1
	   }
	   if tmp == n {
	       return 1
	   }
	   return min(minOperations(tmp-n), minOperations(n-tmp/2)) + 1
	*/
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}