func largestMultipleOfThree(digits []int) string {
    sum := 0
    zero, one, two := []int{}, []int{}, []int{}
    for _, d := range digits {
        sum += d
        if d%3 == 0 {
            zero = append(zero, d)
        } else if d%3 == 1 {
            one = append(one, d)
        } else {
            two = append(two, d)
        }
    }
    if sum == 0 {
        return "0"
    }
    if sum%3 == 0 {
        zero = digits
    } else if sum%3 == 1 {
        if len(one) >= 1 {
            sort.Ints(one)
            zero = append(zero, one[1:]...)
            zero = append(zero, two...)
        } else if len(two) >= 2 {
            sort.Ints(two)
            zero = append(zero, one...)
            zero = append(zero, two[2:]...)
        }
    } else if sum%3 == 2 {
        if len(two) >= 1 {
            sort.Ints(two)
            zero = append(zero, one...)
            zero = append(zero, two[1:]...)
        } else if len(one) >= 2 {
            sort.Ints(one)
            zero = append(zero, one[2:]...)
            zero = append(zero, two...)
        }
    }
    sort.Ints(zero)
    ans := ""
    for i := len(zero)-1; i >= 0; i-- {
        ans += strconv.Itoa(zero[i])
    }
    return ans
}