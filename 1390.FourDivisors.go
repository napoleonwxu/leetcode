func sumFourDivisors(nums []int) int {
    ans := 0
    for _, num := range nums {
        cnt, sum := 0, 0
        i := 1
        for ; i*i < num; i++ {
            if num%i == 0 {
                cnt += 2
                if cnt > 4 {
                    break
                }
                sum += i + num/i
            }
        }
        if i*i != num && cnt == 4 {
            ans += sum
        }
    }
    return ans
}