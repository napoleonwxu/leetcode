func invalidTransactions(transactions []string) []string {
    n := len(transactions)
    name := make([]string, n)
    time := make([]int, n)
    amount := make([]int, n)
    city := make([]string, n)
    invalid := make([]bool, n)
    for i, tran := range transactions {
        info := strings.Split(tran, ",")
        name[i] = info[0]
        time[i], _ = strconv.Atoi(info[1])
        amount[i], _ = strconv.Atoi(info[2])
        city[i] = info[3]
        if amount[i] > 1000 {
            invalid[i] = true
        }
        for j := 0; j < i; j++ {
            if name[i] == name[j] && abs(time[i]-time[j]) <= 60 && city[i] != city[j] {
                invalid[i], invalid[j] = true, true
            }
        }
    }
    ans := []string{}
    for i, inv := range invalid {
        if inv {
            ans = append(ans, transactions[i])
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}