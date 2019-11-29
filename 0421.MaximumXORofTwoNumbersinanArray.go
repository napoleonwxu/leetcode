func findMaximumXOR(nums []int) int {
    ans, mask := 0, 0
    for i := 31; i >= 0; i-- {
        mask |= 1 << uint(i)
        Map := make(map[int]bool)
        for _, num := range nums {
            Map[num&mask] = true
        }
        tmp := ans | 1<<uint(i)
        for prefix := range Map {
            if Map[prefix^tmp] {
                ans = tmp
                break
            }
        }
    }
    return ans
}

func findMaximumAND(nums []int) int {
    ans := 0
    for bit := 1<<31; bit >= 1; bit >>= 1 {
        nxt := []int{}
        for _, num := range nums {
            if num & bit > 0 {
                nxt = append(nxt, num)
            }
        }
        if len(nxt) >= 2 {
            ans |= bit
            nums = nxt
        }
    }
    return ans
}
