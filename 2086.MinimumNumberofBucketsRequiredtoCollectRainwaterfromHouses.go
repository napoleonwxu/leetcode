import "strings"

func minimumBuckets(street string) int {
	if street == "H" || strings.HasPrefix(street, "HH") || strings.HasSuffix(street, "HH") || strings.Contains(street, "HHH") {
		return -1
	}
	return strings.Count(street, "H") - strings.Count(street, "H.H")
}