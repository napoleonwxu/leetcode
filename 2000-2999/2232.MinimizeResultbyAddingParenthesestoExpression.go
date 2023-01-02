func minimizeResult(expression string) string {
    nums := strings.Split(expression, "+")
    num1, num2 := nums[0], nums[1]
    min, ans := math.MaxInt32, ""
    for i := 0; i < len(num1); i++ {
        for j := 1; j <= len(num2); j++ {
            a, _ := strconv.Atoi(num1[:i])
            b, _ := strconv.Atoi(num1[i:])
            c, _ := strconv.Atoi(num2[:j])
            d, _ := strconv.Atoi(num2[j:])
            tmp := b + c
            if a > 0 {
                tmp *= a
            }
            if d > 0 {
                tmp *= d
            }
            if tmp < min {
                min = tmp
                ans = num1[:i] + "(" + num1[i:] + "+" + num2[:j] + ")" + num2[j:]
            }
        }
    }
    return ans
}