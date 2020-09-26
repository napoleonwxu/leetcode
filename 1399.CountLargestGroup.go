func countLargestGroup(n int) int {
    cnts := make([]int, 37)
    for i := 1; i <= n; i++ {
        cnts[sum(i)]++
    }
    max, ans := 0, 0
    for i := 1; i < len(cnts) && cnts[i] > 0; i++ {
        if cnts[i] > max {
            max, ans = cnts[i], 1
        } else if cnts[i] == max {
            ans++
        }
    }
    return ans
}

func sum(x int) int {
    ans := 0
    for ; x > 0; x /= 10 {
        ans += x%10
    }
    return ans
}