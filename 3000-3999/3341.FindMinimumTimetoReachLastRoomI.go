type Point struct {
    row, col, time int
}

type PriorityQueue []Point

func (pq PriorityQueue) Len() int {return len(pq)}
func (pq PriorityQueue) Less(i, j int) bool {return pq[i].time < pq[j].time}
func (pq PriorityQueue) Swap(i, j int) {pq[i], pq[j] = pq[j], pq[i]}
func (pq *PriorityQueue) Push(x interface{}) {*pq = append(*pq, x.(Point))}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[:n-1]
    return item
}

func minTimeToReach(moveTime [][]int) int {
    m, n := len(moveTime), len(moveTime[0])
    dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
    dis := make([][]int, m)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = math.MaxInt64
        }
    }
    dis[0][0] = 0

    pq := &PriorityQueue{}
    heap.Push(pq, Point{row: 0, col: 0, time: 0})
    for pq.Len() > 0 {
        p := heap.Pop(pq).(Point)
        r, c, t := p.row, p.col, p.time
        if r == m-1 && c == n-1 {
            return t
        }
        for _, d := range dirs {
            row, col := r + d[0], c + d[1]
            if row >= 0 && row < m && col >= 0 && col < n {
                nextTime := max(t, moveTime[row][col]) + 1
                if nextTime < dis[row][col] {
                    dis[row][col] = nextTime
                    heap.Push(pq, Point{row, col, nextTime})
                }
            }
        }
    }
    return -1
}