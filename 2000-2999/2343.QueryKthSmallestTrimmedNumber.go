import "sort"

type NumIdx struct {
    num string
    idx int
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
    m, n := len(nums), len(queries)
    ans := make([]int, n)
    for i, q := range queries {
        tmp := make([]NumIdx, m)
        for j, num := range nums {
            tmp[j] = NumIdx{num: num[len(num)-q[1]:], idx: j}
        }
        sort.Slice(tmp, func(i, j int) bool {
            if tmp[i].num == tmp[j].num {
                return tmp[i].idx < tmp[j].idx
            }
            return tmp[i].num < tmp[j].num
        })
        ans[i] = tmp[q[0]-1].idx
    }
    return ans
}