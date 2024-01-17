import "sort"

func maximumTastiness(price []int, k int) int {
    sort.Ints(price)
    l, r := 0, int(1e9)
    for l < r {
        mid := (l + r) / 2
        if check(price, k, mid) {
            l = mid + 1
        } else {
            r = mid
        }
    }
    return l - 1
}

func check(price []int, k, diff int) bool {
    cnt, pre := 1, price[0]
    for i := 1; i < len(price); i++ {
        if price[i]-pre >= diff {
            cnt++
            pre = price[i]
        }
    }
    return cnt >= k
}