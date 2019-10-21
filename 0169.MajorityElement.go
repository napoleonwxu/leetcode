func majorityElement(nums []int) int {
    // O(n) + O(1)
    major, cnt := nums[0], 1
    for i := 1; i < len(nums); i++ {
        if nums[i] == major {
            cnt++
        } else if cnt > 0 {
            cnt--
        } else {
            major, cnt = nums[i], 1
        }
    }
    return major
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