// 2D sub matrix sum: 1074
func subarraySum(nums []int, k int) int {
    cnt, sum := 0, 0
    Map := make(map[int]int)
    Map[0] = 1
    for _, n := range(nums) {
        sum += n
        cnt += Map[sum-k]
        Map[sum]++
    }
    return cnt
}