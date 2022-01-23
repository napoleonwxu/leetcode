func closestDivisors(num int) []int {
    c1, c2 := closest(num+1), closest(num+2)
    if c1[1]-c1[0] <= c2[1]-c2[0] {
        return c1
    }
    return c2
}

func closest(num int) []int {
    ans := []int{1, num}
    for i := 2; i*i <= num; i++ {
        if num%i == 0 {
            ans = []int{i, num/i}
        }
    }
    return ans
}