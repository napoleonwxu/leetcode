import "sort"

func minimumAddedCoins(coins []int, target int) int {
    sort.Ints(coins)
    max, ans := 0, 0
    i := 0
    for max < target {
        if i < len(coins) && coins[i] <= max+1 {
            max += coins[i]
            i++
        } else {
            max += max + 1
            ans++
        }
    }
    return ans
}