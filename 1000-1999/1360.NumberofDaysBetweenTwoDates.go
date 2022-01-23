func daysBetweenDates(date1 string, date2 string) int {
    return abs(cntDay(date1), cntDay(date2))
}

func cntDay(date string) int {
    d := strings.Split(date, "-")
    year, _ := strconv.Atoi(d[0])
    month, _ := strconv.Atoi(d[1])
    day, _ := strconv.Atoi(d[2])
    
    for y := 1971; y < year; y++ {
        day += 365
        if isLeap(y) {
            day++
        }
    }
    for m := 1; m < month; m++ {
        if m == 2 {
            day += 28
            if isLeap(year) {
                day++
            }
        } else if m==4 || m==6 || m==9 || m==11 {
            day += 30
        } else {
            day += 31
        }
    }
    return day
}

func isLeap(year int) bool {
    return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func abs(x, y int) int {
    if x > y {
        return x-y
    }
    return y-x
}