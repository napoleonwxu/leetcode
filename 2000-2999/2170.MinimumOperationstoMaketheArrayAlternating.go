func minimumOperations(nums []int) int {
    even, odd := make(map[int]int), make(map[int]int)
    for i, num := range nums {
        if i&1 == 0 {
            even[num]++
        } else {
            odd[num]++
        }
    }
    
    for num := range even {
        if even[num] > odd[num] { // must >, wrong with >=
            odd[num] = 0
        } else {
            even[num] = 0
        }
    }
    
    evenCnt, oddCnt := 0, 0
    for _, cnt := range even {
        evenCnt = max(evenCnt, cnt)
    }
    for _, cnt := range odd {
        oddCnt = max(oddCnt, cnt)
    }
    return len(nums) - evenCnt - oddCnt
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}