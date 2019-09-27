/*
In an infinite chess board with coordinates from -infinity to +infinity, you have a knight at square [0, 0].
A knight has 8 possible moves it can make, as illustrated below. Each move is two squares in a cardinal direction, then one square in an orthogonal direction.

Return the minimum number of steps needed to move the knight to the square [x, y].  It is guaranteed the answer exists.

Example 1:
Input: x = 2, y = 1
Output: 1
Explanation: [0, 0] → [2, 1]

Example 2:
Input: x = 5, y = 5
Output: 4
Explanation: [0, 0] → [2, 1] → [4, 2] → [3, 4] → [5, 5]

Constraints:
|x| + |y| <= 300
*/

func minKnightMoves(x int, y int) int {
    if x < 0 {
        x = -x
    }
    if y < 0 {
        y = -y
    }
    step := 0
    queue := [][2]int{{0, 0}}
    Map := make(map[[2]int]bool)
    Map[queue[0]] = true
    dx := []int{-2, -2, -1, -1, 1, 1, 2, 2}
    dy := []int{-1, 1, -2, 2, -2, 2, -1, 1}
    for len(queue) > 0 {
        nxt := [][2]int{}
        for _, pos := range queue {            
            if int(pos[0]) == x && int(pos[1]) == y {
                return step
            }
            for i := 0; i < len(dx); i++ {
                tmp := [2]int{pos[0]+dx[i], pos[1]+dy[i]}
                if tmp[0] >= 0 && tmp[1] >= 0 && tmp[0]+tmp[1] <= 300 && !Map[tmp] {
                    Map[tmp] = true
                    nxt = append(nxt, tmp)
                }
            }
        }
        step++
        queue = nxt
    }
    return -1
}
