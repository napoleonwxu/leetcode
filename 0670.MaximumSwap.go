func maximumSwap(num int) int {
    nums := []int{}
    tmp := num
    for tmp > 0 {
        nums = append(nums, tmp%10)
        tmp /= 10
    }
    // O(n)
    index := make(map[int]int, 10)
    for i := len(nums)-1; i >= 0; i-- {
        index[nums[i]] = i
    }
    for i := len(nums)-1; i >= 0; i-- {
        for n := 9; n > nums[i]; n-- {
            if _, ok := index[n]; ok && index[n] < i {
                nums[i], nums[index[n]] = nums[index[n]], nums[i]
                return toInt(nums)
            }
        }
    }
    /* O(n^2)
    for i := len(nums)-1; i > 0; i-- {
        idx := i
        for j := i-1; j >= 0; j-- {
            if nums[j] > nums[idx] || (nums[j] == nums[idx] && idx != i) {
                idx = j
            }
        }
        if idx != i {
            nums[i], nums[idx] = nums[idx], nums[i]
            return toInt(nums)
        }
    } */
    return num
}

func toInt(nums []int) int {
    num := 0
    for i := len(nums)-1; i >= 0; i-- {
        num = 10*num + nums[i]
    }
    return num
}