import "strings"

func findMinimumOperations(s1 string, s2 string, s3 string) int {
    for i := len(s1); i >= 1; i-- {
        if strings.HasPrefix(s2, s1[:i]) && strings.HasPrefix(s3, s1[:i]) {
            return len(s1) + len(s2) + len(s3) - 3*i
        }
    }
    return -1
}