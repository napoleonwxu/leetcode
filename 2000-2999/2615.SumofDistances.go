func distance(nums []int) []int64 {
    m := make(map[int][]int)
    for i, num := range nums {
        m[num] = append(m[num], i)
    }
    ans := make([]int64, len(nums))
    for _, indexes := range m {
        sum := int64(0)
        for _, j := range indexes {
            sum += int64(j)
        }
        leftSum := int64(0)
        for i, j := range indexes {
            rightSum := sum - leftSum - int64(j)
            ans[j] += int64(i*j) - leftSum
            ans[j] += rightSum - int64((len(indexes)-i-1)*j)
            leftSum += int64(j)
        }
    }
    return ans
}
