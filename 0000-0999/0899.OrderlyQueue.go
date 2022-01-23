import (
	"sort"
	"strings"
)

func orderlyQueue(s string, k int) string {
	chs := []byte(s)
	if k > 1 {
		sort.Slice(chs, func(i, j int) bool {
			return chs[i] < chs[j]
		})
		return string(chs)
	}
	ans := s
	for i := 1; i < len(chs); i++ {
		tmp := string(chs[i:]) + string(chs[:i])
		if strings.Compare(tmp, ans) < 0 {
			ans = tmp
		}
	}
	return ans
}