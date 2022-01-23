func sequentialDigits(low int, high int) []int {
    ans := []int{}
    l1, l2 := length(low), length(high)
    mod := 1
    for i := 0; i < l1-1; i++ {
        mod *= 10
    }
    for l := l1; l <= l2; l++ {
        num := 0
        d := 1
        for ; d <= l-1; d++ {
            num = 10*num + d
        }
        for ; d <= 9; d++ {
            num = 10*(num%mod) + d
            if num > high {
                break
            }
            if num >= low {
                ans = append(ans, num)
            }
        }
        mod *= 10
    }
    return ans
}

func length(num int) int {
    l := 0
    for ; num > 0; num /= 10 {
        l++
    }
    return l
}