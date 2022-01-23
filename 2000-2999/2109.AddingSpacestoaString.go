import "strings"

func addSpaces(s string, spaces []int) string {
	words := make([]string, 0, len(spaces)+1)
	idx := 0
	for _, space := range spaces {
		words = append(words, s[idx:space])
		idx = space
	}
	words = append(words, s[idx:len(s)])
	return strings.Join(words, " ")
}