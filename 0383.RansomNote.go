func canConstruct(ransomNote string, magazine string) bool {
	cnts := make([]int, 26)
	for _, ch := range magazine {
		cnts[int(ch-'a')]++
	}
	for _, ch := range ransomNote {
		cnts[int(ch-'a')]--
		if cnts[int(ch-'a')] < 0 {
			return false
		}
	}
	return true
}