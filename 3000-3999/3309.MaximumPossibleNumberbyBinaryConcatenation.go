func maxGoodNumber(nums []int) int {
    sort.Slice(nums, func(i, j int) bool {
        return join(nums[i], nums[j]) > join(nums[j], nums[i])
    })
    ans := 0
    for _, num := range nums {
        ans = join(ans, num)
    }
    return ans
}

func join(x, y int) int {
    tmp, l := y, 0
    for ; tmp > 0; tmp /= 2 {
        l++
    }
    return x << l + y
}
