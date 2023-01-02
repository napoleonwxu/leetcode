func countOperations(num1 int, num2 int) int {
    cnt := 0
    for num1 > 0 && num2 > 0 {
        if num1 < num2 {
            num1, num2 = num2, num1
        }
        cnt += num1 / num2
        num1 %= num2
    }
    return cnt
}