func numberOfSteps (num int) int {
    step := 0
    for num > 0 {
        if num&1 == 0 || num == 1 {
            step++
        } else {
            step += 2
        }
        num >>= 1
    }
    return step
}