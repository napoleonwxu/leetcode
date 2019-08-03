type Node struct {
    Val int
    Cnt int
}

const MOD = 1e9 + 7

func sumSubarrayMins(A []int) int {
    // O(n)
    stack := make([]Node, len(A))
    n := -1
    prod, ans := 0, 0
    for _, a := range A {
        cnt := 1
        for ; n >= 0 && stack[n].Val > a; n-- {
            prod -= stack[n].Val * stack[n].Cnt
            cnt += stack[n].Cnt
        }
        n++
        stack[n] = Node{Val: a, Cnt: cnt}
        prod += a * cnt
        ans = (ans+prod) % MOD
    }
    return ans
}