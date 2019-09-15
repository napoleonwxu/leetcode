func maxNumberOfBalloons(text string) int {
    cnt := make([]int, 5)
    for _, ch := range text {
        switch ch {
        case 'b':
            cnt[0]++
        case 'a':
            cnt[1]++
        case 'l':
            cnt[2]++
        case 'o':
            cnt[3]++
        case 'n':
            cnt[4]++
        }
    }
    cnt[2] /= 2
    cnt[3] /= 2
    min := cnt[0]
    for i := 1; i < 5; i++ {
        if cnt[i] < min {
            min = cnt[i]
        }
    }
    return min
}