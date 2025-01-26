func minimumPushes(word string) int {
    cnts := make([]int, 26)
    for i := range word {
        cnts[int(word[i]-'a')]++
    }
    sort.Ints(cnts)
    ans := 0
    push, left := 1, 8
    for i := len(cnts)-1; i >= 0; i-- {
        ans += push * cnts[i]
        left--
        if left <= 0 {
            push++
            left = 8
        }
    }
    return ans
}