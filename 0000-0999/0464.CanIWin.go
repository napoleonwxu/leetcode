func canIWin(maxChoosableInteger int, desiredTotal int) bool {
    if desiredTotal <= maxChoosableInteger {
        return true
    }
    sum := (1+maxChoosableInteger)*maxChoosableInteger/2
    if sum < desiredTotal {
        return false
    }
    used := make([]bool, maxChoosableInteger+1)
    Map := make(map[int]bool)
    return dfs(desiredTotal, maxChoosableInteger, used, Map)
}

func dfs(desire, max int, used []bool, Map map[int]bool) bool {
    if desire <= 0 {
        return false
    }
    key := toKey(used)
    _, ok := Map[key]
    if ok {
        return Map[key]
    }
    for i := 1; i <= max; i++ {
        if used[i] {
            continue
        }
        used[i] = true
        if (!dfs(desire-i, max, used, Map)) {
            Map[key] = true
            used[i] = false
            return true
        }
        used[i] = false
    }
    Map[key] = false
    return false
}

func toKey(used []bool) int {
    key := 0
    for i := range used {
        key <<= 1
        if used[i] {
            key |= 1
        }
    }
    return key
}