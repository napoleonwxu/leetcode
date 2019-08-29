func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    Map := make(map[string]bool, len(wordList))
    for _, word := range wordList {
        Map[word] = true
    }
    //Map[beginWord] = false
    ans := [][]string{}
    step, minStep := 1, math.MaxInt32
    paths := [][]string{[]string{beginWord}}
    visited := []string{}
    for len(paths) > 0 {
        path := paths[0]
        paths = paths[1:]
        if len(path) > step {
            if len(path) > minStep {
                break
            } else {
                step = len(path)
            }
            for _, word := range visited {
                Map[word] = false
            }
            visited = []string{}
        }
        cur := path[len(path)-1]
        for i := 0; i < len(cur); i++ {
            for ch := 'a'; ch <= 'z'; ch++ {
                nxt := cur[:i] + string(ch) + cur[i+1:]
                if Map[nxt] {
                    newPath := make([]string, len(path)+1)
                    copy(newPath, path)
                    newPath[len(path)] = nxt
                    visited = append(visited, nxt)
                    if nxt == endWord {
                        minStep = step
                        ans = append(ans, newPath)
                    } else {
                        paths = append(paths, newPath)
                    }
                }
            }
        }
    }
    return ans
}