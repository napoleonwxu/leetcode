func threeSumClosest(nums []int, target int) int {
    // O(n^2)
    n := len(nums)
    if n < 3 {
        return 0
    }
    ans := nums[0] + nums[1] + nums[2]
    sort.Ints(nums)
    for i := 0; i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == target {
                return target
            }
            if abs(sum-target) < abs(ans-target) {
                ans = sum
            }
            if sum < target {
                l++
            } else {
                r--
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}