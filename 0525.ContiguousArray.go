func canReorderDoubled(A []int) bool {
    Map := make(map[int]int)
    for _, a := range A {
        Map[a]++
    }
    // O(N + KlgK), K: len of different num in A
    nums := []int{}
    for num, _ := range Map {
        nums = append(nums, num)
    }
    sort.Slice(nums, func(i, j int) bool {
        return abs(nums[i]) < abs(nums[j])
    })
    for _, num := range nums {
        if Map[num] > Map[2*num] {
            return false
        }
        Map[2*num] -= Map[num]
    }
    /* O(NlgN)
    sort.Slice(A, func(i, j int) bool {
        return abs(A[i]) < abs(A[j])
    })
    for _, a := range A {
        if Map[2*a] < Map[a] {
            return false
        }
        Map[2*a] -= Map[a]
        Map[a] = 0
    }
    */
    return true
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}