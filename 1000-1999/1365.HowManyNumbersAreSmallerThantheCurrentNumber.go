func smallerNumbersThanCurrent(nums []int) []int {
    // O(n)
    cnt := make([]int, 101)
    for _, num := range nums {
        cnt[num]++
    }
    for i := 1; i < len(cnt); i++ {
        cnt[i] += cnt[i-1]
    }
    ans := make([]int, len(nums))
    for i, num := range nums {
        if num != 0 {
            ans[i] = cnt[num-1]
        }
    }
    /* O(nlgn)
    tmp := make([]int, len(nums))
    copy(tmp, nums)
    sort.Ints(tmp)
    cnt := 0
    Map := make(map[int]int)
    for i, num := range tmp {
        if i == 0 || num != tmp[i-1] {
            Map[num] = cnt
        }
        cnt++
    }
    ans := make([]int, len(nums))
    for i, num := range nums {
        ans[i] = Map[num]
    } */
    return ans
}