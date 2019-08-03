func complexNumberMultiply(a string, b string) string {
    a1, a2 := toInt(a)
    b1, b2 := toInt(b)
    return fmt.Sprintf("%d+%di", a1*b1-a2*b2, a1*b2+a2*b1)
    //return strconv.Itoa(a1*b1-a2*b2) + "+" + strconv.Itoa(a1*b2+a2*b1) + "i"
}

func toInt(c string) (x, y int) {
    s := strings.Split(c, "+")
    x, _ = strconv.Atoi(s[0])
    y, _ = strconv.Atoi(s[1][:len(s[1])-1])
    return
}