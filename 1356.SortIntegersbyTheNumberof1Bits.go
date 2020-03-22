func sortByBits(arr []int) []int {
    sort.Slice(arr, func(i, j int) bool {
        if cntOne(arr[i]) == cntOne(arr[j]) {
            return arr[i] < arr[j]
        }
        return cntOne(arr[i]) < cntOne(arr[j])
    })
    return arr
}

func cntOne(num int) int {
    cnt := 0
    for num > 0 {
        cnt++
        num &= num-1
    }
    return cnt
}