func rand10() int {
    // O(49/40) = O(1)
    rand40 := 40
    for rand40 >= 40 {
        rand40 = 7*(rand7()-1) + (rand7()-1)
    }
    return rand40%10 + 1
}