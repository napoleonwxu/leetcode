func convertDateToBinary(date string) string {
    fields := strings.Split(date, "-")
    for i, field := range fields {
        num, _ := strconv.Atoi(field)
        fields[i] = strconv.FormatInt(int64(num), 2)
    }
    return strings.Join(fields, "-")
}