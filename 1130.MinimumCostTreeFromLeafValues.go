func mctFromLeafValues(arr []int) int {
    ans := 0
    // O(n) + O(n)
    stack := make([]int, len(arr)+1)
    stack[0] = math.MaxInt32
    n := 0
    for _, a := range arr {
        for stack[n] < a {
            ans += stack[n] * min(stack[n-1], a)
            n--
        }
        n++
        stack[n] = a
    }
    for ; n > 1; n-- {
        ans += stack[n] * stack[n-1]
    }
    /* O(n^2) + O(1)
    for len(arr) > 1 {
        i := minIdx(arr)
        if i == 0 || (i+1 < len(arr) && arr[i+1] < arr[i-1]) {
            ans += arr[i] * arr[i+1]
        } else {
            ans += arr[i] * arr[i-1]
        }
        arr = append(arr[:i], arr[i+1:]...)
    } */
    return ans
}

func minIdx(arr []int) int {
    idx := 0
    for i := 1; i < len(arr); i++ {
        if arr[i] < arr[idx] {
            idx = i
        }
    }
    return idx
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}