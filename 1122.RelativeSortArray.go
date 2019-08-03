func relativeSortArray(arr1 []int, arr2 []int) []int {
    // O(n), n: len(arr1)
    cnt := make([]int, 1001)
    for _, num := range arr1 {
        cnt[num]++
    }
    ans := make([]int, len(arr1))
    i := 0
    for _, num := range arr2 {
        for ; cnt[num] > 0; i++ {
            ans[i] = num
            cnt[num]--
        }
    }
    for num := range cnt {
        for ; cnt[num] > 0; i++ {
            ans[i] = num
            cnt[num]--
        }
    }
    return ans
    /* O(nlogn) with sort.Slice(), n: len(arr1)
    map2 := make(map[int]int, len(arr2))
    for i, num := range arr2 {
        map2[num] = i
    }
    sort.Slice(arr1, func(i, j int) bool {
        idxi, oki := map2[arr1[i]]
        idxj, okj := map2[arr1[j]]
        if oki && okj {
            return idxi < idxj
        } else if oki {
            return true
        } else if okj {
            return false
        }
        return arr1[i] < arr1[j]
    })
    return arr1
    */
}