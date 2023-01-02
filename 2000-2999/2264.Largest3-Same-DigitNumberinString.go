func largestGoodInteger(num string) string {
    ans := ""
    for i := 0; i < len(num)-2; i++ {
        if num[i] == num[i+1] && num[i] == num[i+2] {
            if ans == "" || ans[0] < num[i] {
                ans = num[i:i+3]
            }
        }
    }
    return ans
}