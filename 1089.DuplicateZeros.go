func duplicateZeros(arr []int)  {
    n := len(arr)
    // O(n)
    i, j := 0, 0
    for ; i < n; i, j = i+1, j+1 {
        if arr[i] == 0 {
            j++
        }
    }
    for i, j = i-1, j-1; i >= 0; i, j = i-1, j-1 {
        if j < n {
            arr[j] = arr[i]
        }
        if arr[i] == 0 {
            j--
            if j < n {
                arr[j] = 0
            }
        }
    }
    /* O(n*n)
    for i := 0; i < n; i++ {
        if arr[i] == 0 {
            for j := n-1; j > i; j-- {
                arr[j] = arr[j-1]
            }
            i++
        }
    }
    */
}