func finalString(s string) string {
    chs, flip := []byte(s), false
    ans := []byte{}
    for _, ch := range chs {
        if ch == 'i' {
            flip = !flip
        } else {
            if !flip {
                ans = append(ans, ch)
            } else {
                ans = append([]byte{ch}, ans...)
            }
        }
    }
    if flip {
        reverse(ans)
    }
    return string(ans)
}

func reverse(chs []byte) {
    for i, j := 0, len(chs)-1; i < j; i, j = i+1, j-1 {
        chs[i], chs[j] = chs[j], chs[i]
    }
}
