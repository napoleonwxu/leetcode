// 2D sub matrix sum: 1074
func subarraySum(nums []int, k int) int {
    sum, cnt := 0, 0
    Map := make(map[int]int)
    Map[0] = 1
    for _, num := range nums {
        sum += num
        cnt += Map[sum-k]
        Map[sum]++
    }
    return cnt
}