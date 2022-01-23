func maxSatisfied(customers []int, grumpy []int, X int) int {
    n := len(customers)
    sum0, max1, tmp := 0, 0, 0
    for i := 0; i < n; i++ {
        if grumpy[i] == 0 {
            sum0 += customers[i]
        } else {
            tmp += customers[i]
        }
        if i >= X && grumpy[i-X] == 1 {
            tmp -= customers[i-X]
        }
        if tmp > max1 {
            max1 = tmp
        }
    }
    return sum0 + max1
}