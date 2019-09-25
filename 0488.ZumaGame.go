// AC but wrong, refer https://leetcode.com/problems/zuma-game/discuss/97007/Standard-test-program-is-wrong
const MAXCNT = 6

func findMinStep(board string, hand string) int {
    cnt := make([]int, 26)
    for _, c := range hand {
        cnt[c-'A']++
    }
    insert := dfs(board+"#", cnt)
    if insert == MAXCNT {
        return -1
    }
    return insert
}

func dfs(board string, cnt []int) int {
    board = remove(board)
    if board == "#" {
        return 0
    }
    insert := MAXCNT
    for i, j := 0, 0; j < len(board); j++ {
        if board[j] == board[i] {
            continue
        }
        need := 3 - (j-i)
        if cnt[board[i]-'A'] >= need {
            cnt[board[i]-'A'] -= need
            insert = min(insert, need+dfs(board[:i]+board[j:], cnt))
            cnt[board[i]-'A'] += need
        }
        i = j
    }
    return insert
}

func remove(board string) string {
    for i, j := 0, 0; j < len(board); j++ {
        if board[i] == board[j] {
            continue
        }
        if j-i >= 3 {
            return remove(board[:i]+board[j:])
        }
        i = j
    }
    return board
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}