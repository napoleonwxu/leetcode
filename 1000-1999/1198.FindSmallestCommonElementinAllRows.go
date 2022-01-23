/*
Given a matrix mat where every row is sorted in increasing order, return the smallest common element in all rows.
If there is no common element, return -1.

Example 1:
Input: mat = [[1,2,3,4,5],[2,4,5,8,10],[3,5,7,9,11],[1,3,5,7,9]]
Output: 5

Constraints:
1 <= mat.length, mat[i].length <= 500
1 <= mat[i][j] <= 10^4
mat[i] is sorted in increasing order.
*/

func smallestCommonElement(mat [][]int) int {
    m, n := len(mat), len(mat[0])
    minLast := mat[0][n-1]
    for i := 1; i < m; i++ {
        if mat[i][n-1] < minLast {
            minLast = mat[i][n-1]
        }
    }
    Map := make(map[int]bool)
    for _, num := range mat[0] {
        if num > minLast {
            break
        }
        Map[num] = true
    }
    for i := 1; i < m; i++ {
        nxt := make(map[int]bool)
        for _, num := range mat[i] {
            if num > minLast {
                break
            }
            if Map[num] {
                nxt[num] = true
            }
        }
        if len(nxt) == 0 {
            return -1
        }
        Map = nxt
    }
    ans := math.MaxInt32
    for num := range Map {
        if num < ans {
            ans = num
        }
    }
    return ans
}