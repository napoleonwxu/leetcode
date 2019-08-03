func repeatedNTimes(A []int) int {
    // O(n), no extra space
    for i := 2; i < len(A); i++ {
        if A[i] == A[i-1] || A[i] == A[i-2] {
            return A[i]
        }
    }
    // In case of [x, x, y, z] or [x, y, z, x], return A[0]
    return A[0]
    
    /* O(n) + O(n)
    N := len(A)/2
    Map := make(map[int]bool, N+1)
    for _, a := range A {
        if Map[a] {
            return a
        }
        Map[a] = true
    } 
    return -1
    */
}