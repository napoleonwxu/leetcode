func corpFlightBookings(bookings [][]int, n int) []int {
    ans := make([]int, n)
    // O(n)
    for _, b := range bookings {
        ans[b[0]-1] += b[2]
        if b[1] < n {
            ans[b[1]] -= b[2]
        }
    }
    for i := 1; i < n; i++ {
        ans[i] += ans[i-1]
    }
    /* O(n^2)
    for _, b := range bookings {
        for i := b[0]; i <= b[1]; i++ {
            ans[i-1] += b[2]
        }
    }
    */
    return ans
}