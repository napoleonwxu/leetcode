var power map[int]int

func getKth(lo int, hi int, k int) int {
    power = make(map[int]int)
    power[1] = 0
    for i := lo; i <= hi; i++ {
        calPower(i)
    }
    nums := make([]int, hi-lo+1)
    for i := range nums {
        nums[i] = lo + i
    }
    sort.Slice(nums, func(i, j int) bool {
        if power[nums[i]] == power[nums[j]] {
            return nums[i] < nums[j]
        }
        return power[nums[i]] < power[nums[j]]
    })
    return nums[k-1]
}

func calPower(x int) int {
    if _, ok := power[x]; !ok {
        if x&1 == 0 {
            power[x] = calPower(x>>1) + 1
        } else {
            power[x] = calPower(3*x+1) + 1
        }
    }
    return power[x]
}