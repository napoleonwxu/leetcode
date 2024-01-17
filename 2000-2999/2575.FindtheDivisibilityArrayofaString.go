    func divisibilityArray(word string, m int) []int {
    n, num := len(word), 0
    ans := make([]int, n)
    for i := 0; i < n; i++ {
        num = (10*num + int(word[i]-'0')) % m
        if num == 0 {
            ans[i] = 1
        }
    }
    return ans
    }