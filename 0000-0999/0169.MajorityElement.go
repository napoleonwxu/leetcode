func majorityElement(nums []int) int {
    // O(n) + O(1)
    major, cnt := 0, 0
    for _, num := range nums {
        if num == major {
            cnt++
        } else if cnt > 0 {
            cnt--
        } else {
            major, cnt = num, 1
        }
    }
    /* O(n) + O(n)
    cnt := make(map[int]int)
    for _, num := range nums {
        cnt[num]++
        if cnt[num] > len(nums)>>1 {
            return num
        }
    }
    return -1
    */
}