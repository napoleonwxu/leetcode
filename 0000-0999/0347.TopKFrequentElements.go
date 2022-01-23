func topKFrequent(nums []int, k int) []int {
    ans := make([]int, 0, k)
    Map := make(map[int]int)
    // O(n) + O(n)
    maxCnt := 0
    for _, num := range nums {
        Map[num]++
        if Map[num] > maxCnt {
            maxCnt = Map[num]
        }
    }
    bucket := make([][]int, maxCnt+1)
    for num, cnt := range Map {
        bucket[cnt] = append(bucket[cnt], num)
    }
    cnt := maxCnt
    for k > 0 {
        ans = append(ans, bucket[cnt]...)
        k -= len(bucket[cnt])
        cnt--
    }
    /* sort, max(O(n), O(ulgu)) + O(n), u: len(unique)
    for _, num := range nums {
        Map[num]++
    }
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
