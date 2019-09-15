/*
Consider a matrix M with dimensions width * height, such that every cell has value 0 or 1, and any square sub-matrix of M of size sideLength * sideLength has at most maxOnes ones.
Return the maximum possible number of ones that the matrix M can have.

Example 1:
Input: width = 3, height = 3, sideLength = 2, maxOnes = 1
Output: 4
Explanation:
In a 3*3 matrix, no 2*2 sub-matrix can have more than 1 one.
The best solution that has 4 ones is:
[1,0,1]
[0,0,0]
[1,0,1]

Example 2:
Input: width = 3, height = 3, sideLength = 2, maxOnes = 2
Output: 6
Explanation:
[1,0,1]
[1,0,1]
[1,0,1]

Constraints:
1 <= width, height <= 100
1 <= sideLength <= width, height
0 <= maxOnes <= sideLength * sideLength
*/
func maximumNumberOfOnes(width int, height int, sideLength int, maxOnes int) int {
    d1, m1 := width/sideLength, width%sideLength
    d2, m2 := height/sideLength, height%sideLength
    t := m1 * m2
    ans := 0
    if maxOnes <= t {
        return maxOnes * (d1+1) * (d2+1)
    } else {
        maxOnes -= t
        ans += t * (d1+1) * (d2+1)
    }
    if d1*(d2+1) < d2*(d1+1) {
        d1, d2 = d2, d1
        m1, m2 = m2, m1
    }
    t = (sideLength - m1) * m2
    if maxOnes <= t {
        return ans + maxOnes * d1 * (d2+1)
    } else {
        maxOnes -= t
        ans += t * d1 * (d2+1)
    }
    t = (sideLength - m2) * m1
    if maxOnes <= t {
        return ans + maxOnes * d2 * (d1+1)
    } else {
        maxOnes -= t
        ans += t * d2 * (d1+1)
    }
    return ans + maxOnes * d1 * d2
}
