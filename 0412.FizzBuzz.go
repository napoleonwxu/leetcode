func fizzBuzz(n int) []string {
    // Not using "%" operation
    ans := make([]string, n)
    fizz, buzz := 3, 5
    for i := 1; i <= n; i++ {
        if i == fizz && i == buzz {
            ans[i-1] = "FizzBuzz"
            fizz += 3
            buzz += 5
        } else if i == fizz {
            ans[i-1] = "Fizz"
            fizz += 3
        } else if i == buzz {
            ans[i-1] = "Buzz"
            buzz += 5
        } else {
            ans[i-1] = strconv.Itoa(i)
        }
    }
    return ans
}