func checkAlmostEquivalent(word1 string, word2 string) bool {
	cnts := make([]int, 26)
	chs1, chs2 := []byte(word1), []byte(word2)
	for i := range chs1 {
		cnts[int(chs1[i]-'a')]++
		cnts[int(chs2[i]-'a')]--
	}
	for _, cnt := range cnts {
		if cnt > 3 || cnt < -3 {
			return false
		}
	}
	return true
}