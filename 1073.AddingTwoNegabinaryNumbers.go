/*
2*(-2)^i = -(-2)^(i+1)
         = (-2)^(i+2) + (-2)^(i+1)
*/
func addNegabinary(arr1 []int, arr2 []int) []int {
    n1, n2 := len(arr1), len(arr2)
    n := max(n1, n2) + 2
    ans := make([]int, n)
    for i1, i := n1-1, n-1; i1 >= 0; i1, i = i1-1, i-1 {
        ans[i] += arr1[i1]
    }
    for i2, i := n2-1, n-1; i2 >= 0; i2, i = i2-1, i-1 {
        ans[i] += arr2[i2]
    }
    for i := n-1; i >= 2; i-- {
        if ans[i] >= 2 {
            ans[i] -= 2
            if ans[i-1] >= 1 {
                ans[i-1]--
            } else {
                ans[i-1]++
                ans[i-2]++
            }
        }
    }
    i := 0
    for ; i < n-1 && ans[i] == 0; i++ {}
    return ans[i:]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}