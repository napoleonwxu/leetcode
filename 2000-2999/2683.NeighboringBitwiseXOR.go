func doesValidArrayExist(derived []int) bool {
	xor := 0
	for _, d := range derived {
		xor ^= d
	}
	return xor == 0
	/*
	   n := len(derived)
	   origin := make([]int, n)
	   origin[n-1] = derived[n-1] ^ derived[0]
	   for i := n-2; i >= 0; i-- {
	       origin[i] = derived[i] ^ origin[i+1]
	   }
	   tmp := origin[0]
	   for i := 0; i < n-1; i++ {
	       origin[i] ^= origin[i+1]
	   }
	   origin[n-1] ^= tmp
	   for i := 0; i < n; i++ {
	       if origin[i] != derived[i] {
	           return false
	       }
	   }
	   return true
	*/
}