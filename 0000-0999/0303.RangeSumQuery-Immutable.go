type NumArray struct {
    dp []int
}


func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums)+1)
    for i := 0; i < len(nums); i++ {
        sums[i+1] = sums[i] + nums[i]
    }
    return NumArray{dp: sums}
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.dp[j+1] - this.dp[i]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */