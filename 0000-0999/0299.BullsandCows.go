import "fmt"

func getHint(secret string, guess string) string {
	x, y := 0, 0
	cnts := make([]int, 10)
	sChs, gChs := []byte(secret), []byte(guess)
	for i := 0; i < len(sChs); i++ {
		if sChs[i] == gChs[i] {
			x++
		} else {
			if cnts[int(sChs[i]-'0')] < 0 {
				y++
			}
			if cnts[int(gChs[i]-'0')] > 0 {
				y++
			}
			cnts[int(sChs[i]-'0')]++
			cnts[int(gChs[i]-'0')]--
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}