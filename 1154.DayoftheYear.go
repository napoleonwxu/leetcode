func dayOfYear(date string) int {
    y, _ := strconv.Atoi(date[:4])
    m, _ := strconv.Atoi(date[5:7])
    d, _ := strconv.Atoi(date[8:])
    day := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
    for i := 1; i < m; i++ {
        d += day[i-1]
        if i == 2 && (y%400 == 0 || (y%4 == 0 && y%100 != 0)) {
            d++
        }
    }
    return d
}