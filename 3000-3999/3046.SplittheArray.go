func isPossibleToSplit(nums []int) bool {
    cnt := make(map[int]int, len(nums))
    for _, num := range nums {
        cnt[num]++
        if cnt[num] > 2 {
            return false
        }
    }
    return true
}