import "sort"

func maxIncreasingGroups(usageLimits []int) int {
    sort.Ints(usageLimits)
    sum, ans := 0, 0
    for _, cnt := range usageLimits {
        sum += cnt
        if sum >= (1+(ans+1))*(ans+1)/2 {
            ans++
        }
    }
    return ans
}