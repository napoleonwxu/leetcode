import "sort"

const MOD = 1e9 + 7

func maximizeSquareArea(m int, n int, hFences []int, vFences []int) int {
    h := getDiff(m, hFences)
    v := getDiff(n, vFences)
    ma := -1
    for l := range h {
        if v[l] {
            ma = max(ma, l)
        }
    }
    if ma == -1 {
        return -1
    }
    return ma * ma % MOD
}

func getDiff(n int, nums []int) map[int]bool {
    nums = append(nums, 1, n)
    sort.Ints(nums)
    m := make(map[int]bool)
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            m[nums[j]-nums[i]] = true
        }
    }
    return m
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}