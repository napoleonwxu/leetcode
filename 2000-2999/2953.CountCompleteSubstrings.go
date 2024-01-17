func countCompleteSubstrings(word string, k int) int {
    idx := make([]int, 0, len(word))
    for i := range word {
        if i > 0 && abs(int(word[i])-int(word[i-1])) > 2 {
            idx = append(idx, i)
        }
    }
    ans := 0
    for times := 1; times <= 26; times++ {
        cnt := make(map[byte]int, 0)
        l, r := 0, times*k-1
        if r >= len(word) {
            break
        }
        for i := l; i < r; i++ {
            cnt[word[i]]++
        }
        for ; r < len(word); l, r = l+1, r+1 {
            cnt[word[r]]++
            //if valid(l, r, idx) && check(cnt, k) { // TLE
            if check(cnt, k) && valid(l, r, idx) {
                ans++
            }
            cnt[word[l]]--
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func valid(l, r int, idx []int) bool {
    for _, i := range idx {
        if i > l && i <= r {
            return false
        }
    }
    return true
}

func check(cnt map[byte]int, k int) bool {
    for _, c := range cnt {
        if c != 0 && c != k {
            return false
        }
    }
    return true
}
