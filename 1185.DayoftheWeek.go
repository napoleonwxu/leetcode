func dayOfTheWeek(day int, month int, year int) string {
    week := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
    t := []int{0, 3, 2, 5, 0, 3, 5, 1, 4, 6, 2, 4};
    if month < 3 {
        year--
    }
    return week[(year + year/4 - year/100 + year/400 + t[month-1] + day) % 7]
}