func parseBoolExpr(expression string) bool {
    if expression == "t" {
        return true
    }
    if expression == "f" {
        return false
    }
    op := expression[0]
    ans := op != '|'
    cnt := 0
    for i, pre := 1, 2; i < len(expression); i++ {
        if expression[i] == '(' {
            cnt++
        } else if expression[i] == ')' {
            cnt--
        }
        if (expression[i] == ',' && cnt == 1) || (expression[i] == ')' && cnt == 0) {
            nxt := parseBoolExpr(expression[pre:i])
            pre = i + 1
            if op == '&' {
                ans = ans && nxt
            } else if op == '|' {
                ans = ans || nxt
            } else {
                ans = !nxt
            }
        }
    }
    return ans
}