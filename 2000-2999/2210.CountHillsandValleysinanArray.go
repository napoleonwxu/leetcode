func countHillValley(nums []int) int {
    tmp := make([]int, 0, len(nums))
    for i, num := range nums {
        if i == 0 || num != nums[i-1] {
            tmp = append(tmp, num)
        }
    }
    cnt := 0
    for i := 1; i < len(tmp)-1; i++ {
        if (tmp[i] > tmp[i-1] && tmp[i] > tmp[i+1]) || (tmp[i] < tmp[i-1] && tmp[i] < tmp[i+1]) {
            cnt++
        }
    }
    return cnt
}