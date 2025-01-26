func minimumCost(m int, n int, horizontalCut []int, verticalCut []int) int64 {
    sort.Ints(horizontalCut)
    sort.Ints(verticalCut)
    ans := int64(0)
    x, y := 1, 1
    for len(horizontalCut) + len(verticalCut) > 0 {
        h, v := len(horizontalCut), len(verticalCut)
        if h > 0 && v > 0 {
            if horizontalCut[h-1] > verticalCut[v-1] {
                ans += int64(y * horizontalCut[h-1])
                horizontalCut = horizontalCut[:h-1]
                x++
            } else {
                ans += int64(x * verticalCut[v-1])
                verticalCut = verticalCut[:v-1]
                y++
            }
        } else if h > 0 {
            ans += int64(y * sum(horizontalCut))
            break
        } else {
            ans += int64(x * sum(verticalCut))
            break
        }
    }
    return ans
}

func sum(nums []int) int {
    ans := 0
    for _, num := range nums {
        ans += num
    }
    return ans
}