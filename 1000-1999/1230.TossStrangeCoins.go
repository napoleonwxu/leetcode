/*
You have some coins.  The i-th coin has a probability prob[i] of facing heads when tossed.
Return the probability that the number of coins facing heads equals target if you toss every coin exactly once.

Example 1:
Input: prob = [0.4], target = 1
Output: 0.40000

Example 2:
Input: prob = [0.5,0.5,0.5,0.5,0.5], target = 0
Output: 0.03125
 
Constraints:
1 <= prob.length <= 1000
0 <= prob[i] <= 1
0 <= target <= prob.length
Answers will be accepted as correct if they are within 10^-5 of the correct answer.
*/

func probabilityOfHeads(prob []float64, target int) float64 {
    // DP, O(n*target)
    dp := make([]float64, target+1)
    dp[0] = 1
    for _, p := range prob {
        for c := target; c > 0; c-- {
            dp[c] = dp[c-1]*p + dp[c]*(1-p)
        }
        dp[0] = dp[0]*(1-p)
    }
    return dp[target]
    /* O(n*2^n), TLE
    n := len(prob)
    ans := 0.0
    for i := 0; i < 1<<uint(n); i++ {
        if targetOne(i, target) {
            ii, tmp := i, 1.0
            for j := 0; j < n; j++ {
                if ii&1 == 1 {
                    tmp *= prob[j]
                } else {
                    tmp *= 1 - prob[j]
                }
                ii >>= 1
            }
            ans += tmp
        }
    }
    return ans
    */
}

func targetOne(n, target int) bool {
    for n > 0 {
        target--
        if target < 0 {
            return false
        }
        n &= n-1
    }
    return target==0
}