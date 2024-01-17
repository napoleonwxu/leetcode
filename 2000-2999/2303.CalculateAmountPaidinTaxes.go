func calculateTax(brackets [][]int, income int) float64 {
    tax := float64(0)
    pre := 0
    for _, b := range brackets {
        tax += float64((min(income, b[0])-pre)*b[1]) / 100
        if income <= b[0] {
            break
        }
        pre = b[0]
    }
    return tax
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}