func topKFrequent(nums []int, k int) []int {
    Map := make(map[int]int)
    for _, num := range nums {
        Map[num]++
    }
    ans := make([]int, 0, k)
    bucket := make([][]int, len(nums)+1)
    for num, c := range Map {
        bucket[c] = append(bucket[c], num)
    }
    c := len(nums)
    for k > 0 {
        if len(bucket[c]) > 0 {
            ans = append(ans, bucket[c]...)
            k -= len(bucket[c])
        }
        c--
    }
    /* sort, O(ulgu) u: len(unique)
    cnt := make([]int, 0, len(Map))
    for _, c := range Map {
        cnt = append(cnt, c)
    }
    sort.Ints(cnt)
    min := cnt[len(cnt)-k]
    for num, c := range Map {
        if c >= min {
            ans = append(ans, num)
        }
    } */
    return ans
}