/*
Given an integer array arr, in one move you can select a palindromic subarray arr[i], arr[i+1], ..., arr[j] where i <= j, and remove that subarray from the given array. Note that after removing a subarray, the elements on the left and on the right of that subarray move to fill the gap left by the removal.
Return the minimum number of moves needed to remove all numbers from the array.

Example 1:
Input: arr = [1,2]
Output: 2

Example 2:
Input: arr = [1,3,4,1,5]
Output: 3
Explanation: Remove [4] then remove [1,3,1] then remove [5].

Constraints:
1 <= arr.length <= 100
1 <= arr[i] <= 20
*/

func minimumMoves(arr []int) int {
    // O(n^3)
    n := len(arr)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        dp[i][i] = 1
    }
    for size := 2; size <= n; size++ {
        for i := 0; i <= n-size; i++ {
            j := i + size - 1
            dp[i][j] = dp[i+1][j] + 1
            if arr[i] == arr[i+1] {
                dp[i][j] = min(dp[i][j], dp[i+2][j]+1)
            }
            for k := i+2; k <= j; k++ {
                if arr[i] == arr[k] {
                    dp[i][j] = min(dp[i][j], dp[i+1][k-1]+dp[k+1][j])
                }
            }
        }
    }
    return dp[0][n-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}