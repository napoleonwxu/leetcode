func splitIntoFibonacci(S string) []int {
    n := len(S)
    MAX := 1<<31 - 1 //2^31-1
    for i := 0; i < min(10, n); i++ {
        if S[0] == '0' && i > 0 {
            break
        }
        x, _ := strconv.Atoi(S[:i+1])
        for j := i+1; j < min(i+10, n); j++ {
            if S[i+1] == '0' && j > i+1 {
                break
            }
            y, _ := strconv.Atoi(S[i+1:j+1])
            f := []int{x, y}
            fn := 2
            k := j + 1
            for k < n {
                z := f[fn-1] + f[fn-2]
                zs := strconv.Itoa(z)
                if z <= MAX && strings.HasPrefix(S[k:], zs) {
                    f = append(f, z)
                    fn++
                    k += len(zs)
                } else {
                    goto search
                }
            }
            if fn >= 3 {
                return f
            }
            search:
        }
    }
    return []int{}
}

func min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}