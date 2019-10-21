/*
A Stepping Number is an integer such that all of its adjacent digits have an absolute difference of exactly 1. For example, 321 is a Stepping Number while 421 is not.
Given two integers low and high, find and return a sorted list of all the Stepping Numbers in the range [low, high] inclusive.

Example 1:
Input: low = 0, high = 21
Output: [0,1,2,3,4,5,6,7,8,9,10,12,21]

Constraints:
0 <= low <= high <= 2 * 10^9
*/

func countSteppingNumbers(low int, high int) []int {
    ans := []int{}
    for i := 1; i <= 9; i++ {
        search(i, low, high, &ans)
    }
    if low == 0 {
        ans = append(ans, 0)
    }
    sort.Ints(ans)
    return ans
}

func search(num, low, high int, ans *[]int) {
    if num > high {
        return
    }
    if num >= low {
        *ans = append(*ans, num)
    }
    if num%10 < 9 {
        search(10*num+num%10+1, low, high, ans)
    }
    if num%10 > 0 {
        search(10*num+num%10-1, low, high, ans)
    }
}