func kthCharacter(k int) byte {
    //return 'a' + byte(bitCount(k-1))
    chs := make([]byte, k)
    chs[0] = 'a'
    for i := 1; i < len(chs); {
        tmp := i
        for j := 0; j < tmp && i < len(chs); j++ {
            if chs[j] == 'z' {
                chs[i] = 'a'
            } else {
                chs[i] = chs[j] + 1
            }
            i++
        }
    }
    return chs[k-1]
}

func bitCount(num int) int {
    cnt := 0
    for num > 0 {
        cnt++
        num &= (num-1)
    }
    return cnt
}