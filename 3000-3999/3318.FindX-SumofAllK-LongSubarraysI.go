type Cnt struct {
    val int
    cnt int
}

func findXSum(nums []int, k int, x int) []int {
    n := len(nums)
    m := make(map[int]*Cnt, n)
    for i := 0; i < k; i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = &Cnt{val: nums[i]}
        }
        m[nums[i]].cnt++
    }
    ans := make([]int, 0, n-k+1)
    ans = append(ans, xSum(m, x))
    for i := k; i < n; i++ {
        if _, ok := m[nums[i]]; !ok {
            m[nums[i]] = &Cnt{val: nums[i]}
        }
        m[nums[i]].cnt++
        m[nums[i-k]].cnt--
        ans = append(ans, xSum(m, x))
    }
    return ans
}

func xSum(m map[int]*Cnt, x int) int {
    cnts := make([]*Cnt, 0, len(m))
    for _, cnt := range m {
        cnts = append(cnts, cnt)
    }
    sort.Slice(cnts, func(i, j int) bool {
        if cnts[i].cnt == cnts[j].cnt {
            return cnts[i].val > cnts[j].val
        }
        return cnts[i].cnt > cnts[j].cnt
    })
    sum := 0
    for i := 0; i < x && i < len(cnts); i++ {
        sum += cnts[i].val * cnts[i].cnt
    }
    return sum
}