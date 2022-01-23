func orderOfLargestPlusSign(N int, mines [][]int) int {
    zero := make(map[int]bool, len(mines))
    for _, mine := range mines {
        zero[mine[0]*N+mine[1]] = true
    }
    order := make([][]int, N)
    for i := range order {
        order[i] = make([]int, N)
    }
    ans := 0
    for i := 0; i < N; i++ {
        cnt := 0
        for j := 0; j < N; j++ {
            if zero[i*N+j] {
                cnt = 0
            } else {
                cnt++
            }
            order[i][j] = cnt
        }
        cnt = 0
        for j := N-1; j >= 0; j-- {
            if zero[i*N+j] {
                cnt = 0
            } else {
                cnt++
            }
            order[i][j] = min(order[i][j], cnt)
        }
    }
    for j := 0; j < N; j++ {
        cnt := 0
        for i := 0; i < N; i++ {
            if zero[i*N+j] {
                cnt = 0
            } else {
                cnt++
            }
            order[i][j] = min(order[i][j], cnt)
        }
        cnt = 0
        for i := N-1; i >= 0; i-- {
            if zero[i*N+j] {
                cnt = 0
            } else {
                cnt++
            }
            order[i][j] = min(order[i][j], cnt)
            if order[i][j] > ans {
                ans = order[i][j]
            }
        }
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}