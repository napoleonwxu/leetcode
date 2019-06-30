func fib(N int) int {
    if N == 0 || N == 1{
        return N
    }
    // return fib(N-1) + fib(N-2) //slow
    f := make([]int, N+1)
    f[1] = 1
    for i := 2; i <= N; i++ {
        f[i] = f[i-1] + f[i-2]
    }
    return f[N]
}