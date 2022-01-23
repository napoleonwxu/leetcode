// same with 767, 656ms, need optimize
func rearrangeBarcodes(barcodes []int) []int {
    n := len(barcodes)
    if n <= 2 {
        return barcodes
    }
    Map := make(map[int]int)
    for _, b := range barcodes {
        Map[b]++
    }
    cnt := [][]int{}
    for b, c := range Map {
        cnt = append(cnt, []int{c, b})
    }
    sort.Slice(cnt, func (i, j int) bool {
        return cnt[i][0] > cnt[j][0]
    })
    ans := make([]int, n)
    i := 0
    for _, c := range cnt {
        for j := c[0]; j > 0; j-- {
            ans[i] = c[1]
            i += 2
            if i >= n {
                i = 1
            }
        }
    }
    return ans
}