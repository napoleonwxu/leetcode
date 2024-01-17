import "strings"

func garbageCollection(garbage []string, travel []int) int {
    n := len(garbage)
    for i := 1; i < n-1; i++ {
        travel[i] += travel[i-1]
    }
    ans := 0
    for _, g := range []string{"M", "P", "G"} {
        found := false
        for i := n - 1; i >= 0; i-- {
            if !found && strings.Contains(garbage[i], g) {
                found = true
                if i > 0 {
                    ans += travel[i-1]
                }
            }
            if found {
                ans += strings.Count(garbage[i], g)
            }
        }
    }
    return ans
}