func longestConsecutive(nums []int) int {
    /* O(n) + O(n)
    Map := make(map[int]int)
    cnt := 0
    for _, num := range nums {
        if Map[num] == 0 {
            left, right := Map[num-1], Map[num+1]
            Map[num] = left + right + 1
            Map[num-left] = Map[num]
            Map[num+right] = Map[num]
            cnt = max(cnt, Map[num])
        }
    }*/
    // O(nlgn) + O(1)
    if len(nums) <= 1 {
        return len(nums)
    }
    sort.Ints(nums)
    cnt, tmp := 0, 1
    for i := 1; i < len(nums); i++ {
        if nums[i]-1 == nums[i-1] {
            tmp++
        } else if nums[i] != nums[i-1] { // deal with [0, 1, 1, 2]
            tmp = 1
        }
        cnt = max(cnt, tmp)
    }
    return cnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}