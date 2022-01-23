func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    // O(MlgM) + O(1)
    sort.Slice(reservedSeats, func(i, j int) bool {
        return reservedSeats[i][0] < reservedSeats[j][0]
    })
    reservedSeats = append(reservedSeats, []int{0, 0})
    cnt := 2*n
    Map := make(map[int]bool, 10)
    for i := 0; i < len(reservedSeats); i++ {
        if i == 0 || reservedSeats[i][0] == reservedSeats[i-1][0] {
            Map[reservedSeats[i][1]] = true
            continue
        }
        tmp := 0
        if !Map[2] && !Map[3] && !Map[4] && !Map[5] {
            tmp++
        }
        if !Map[6] && !Map[7] && !Map[8] && !Map[9] {
            tmp++
        }
        if !Map[4] && !Map[5] && !Map[6] && !Map[7] && tmp==0 {
            tmp++
        }
        cnt += tmp - 2
        Map = make(map[int]bool, 10)
        Map[reservedSeats[i][1]] = true
    }
    return cnt
}