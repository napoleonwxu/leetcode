import "strings"

func removeOccurrences(s string, part string) string {
	idx := strings.Index(s, part)
	if idx == -1 {
		return s
	}
	chs := []byte(s)
	chs = append(chs[:idx], chs[idx+len(part):]...)
	return removeOccurrences(string(chs), part)
}