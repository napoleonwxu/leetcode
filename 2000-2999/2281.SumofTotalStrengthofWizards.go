const MOD = 1e9 + 7

func totalStrength(strength []int) int {
    n := len(strength)
    sumL, sumR := make([]int, n+1), make([]int, n+1)
    mulL, mulR := make([]int, n+1), make([]int, n+1)
    for i := 0; i < n; i++ {
        sumL[i+1] = (sumL[i] + strength[i]) % MOD
        mulL[i+1] = (mulL[i] + (i+1)*strength[i]) % MOD
    }
    for i := n - 1; i >= 0; i-- {
        sumR[i] = (sumR[i+1] + strength[i]%MOD) % MOD
        mulR[i] = (mulR[i+1] + (n-i)*strength[i]%MOD) % MOD
    }
    ans := 0
    stack := make([]int, 0, n)
    for j := 0; j <= n; j++ {
        for len(stack) > 0 && (j == n || strength[stack[len(stack)-1]] >= strength[j]) {
            piv := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            i := 0
            if len(stack) > 0 {
                i = stack[len(stack)-1] + 1
            }
            sumLeft := (MOD + mulL[piv+1] - mulL[i] - i*(sumL[piv+1]-sumL[i])%MOD) % MOD
            sumRight := (MOD + mulR[piv+1] - mulR[j] - (n-j)*(sumR[piv+1]-sumR[j])%MOD) % MOD
            sum := (sumRight*(piv-i+1) + sumLeft*(j-piv)) % MOD
            ans = (ans + sum*strength[piv]%MOD) % MOD
        }
        stack = append(stack, j)
    }
    return ans
}
