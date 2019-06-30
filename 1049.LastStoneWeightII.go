/*
partition an array into 2 subsets whose difference is minimal
sum = s1 + s2
diff = s1 - s2 = sum - 2*s2
minimize diff ==> maximize s2(range from 0 to sum/2)
same as 494.Target Sum
*/
func lastStoneWeightII(stones []int) int {
    sum := 0
    for _, s := range stones {
        sum += s
    }
    dp := make([]bool, sum+1)
    dp[0] = true
    for _, s := range stones {
        for i := sum; i >= s; i-- {
            dp[i] = dp[i] || dp[i-s]
        }
    }
    for i := sum>>1; i > 0; i-- {
        if dp[i] {
            return sum - i - i
        }
    }
    return 0
}