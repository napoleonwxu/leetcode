func checkIfExist(arr []int) bool {
    Map := make(map[int]bool)
    for _, num := range arr {
        if Map[2*num] || (num&1 == 0 && Map[num/2]) {
            return true
        }
        Map[num] = true
    }
    return false
}