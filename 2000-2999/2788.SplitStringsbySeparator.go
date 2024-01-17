import "strings"

func splitWordsBySeparator(words []string, separator byte) []string {
    ans := []string{}
    sep := string(separator)
    for _, word := range words {
        for _, w := range strings.Split(word, sep) {
            if w != "" {
                ans = append(ans, w)
            }
        }
    }
    return ans
}