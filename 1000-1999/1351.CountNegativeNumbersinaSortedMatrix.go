func countNegatives(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    cnt := 0
    // O(M+N)
    i, j := m-1, 0
    for i >= 0 && j < n {
        if grid[i][j] < 0 {
            cnt += n-j
            i--
        } else {
            j++
        }
    }
    /* O(MlgN)
    for i := 0; i < m; i++ {
        cnt += n - binSearch(grid[i]) - 1
    } */
    return cnt
}

func binSearch(nums []int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l+r) / 2
        if nums[mid] < 0 {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return r
}