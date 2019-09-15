func findNumOfValidWords(words []string, puzzles []string) []int {
    nums := words2nums(words)
    n := len(puzzles)
    ans := make([]int, n)
    for i, p := range puzzles {
        num_p := 0
        for j := range p {
            num_p |= 1 << (p[j]-'a')
        }
        first := 1 << (p[0]-'a')
        cnt := 0
        for _, num := range nums {
            if num&num_p == num && num&first == first {
                cnt++
            }
        }
        ans[i] = cnt
    }
    return ans
}

func words2nums (words []string) []int {
    nums := make([]int, len(words))
    for i, w := range words {
        num := 0
        for j := range w {
            num |= 1 << (w[j]-'a')
        }
        nums[i] = num
    }
    return nums
}