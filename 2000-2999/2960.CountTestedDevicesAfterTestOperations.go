func countTestedDevices(batteryPercentages []int) int {
    cnt := 0
    for _, num := range batteryPercentages {
        if num > cnt {
            cnt++
        }
    }
    return cnt
}