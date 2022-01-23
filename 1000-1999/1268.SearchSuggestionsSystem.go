func suggestedProducts(products []string, searchWord string) [][]string {
    sort.Strings(products)
    ans := make([][]string, 0, len(searchWord))
    l, r := 0, len(products)-1
    for i := range searchWord {
        for l <= r && (len(products[l]) <= i || products[l][i] != searchWord[i]) {
            l++
        }
        for l <= r && (len(products[r]) <= i || products[r][i] != searchWord[i]) {
            r--
        }
        tmp := make([]string, 0, 3)
        for j := l; j < l+3 && j <= r; j++ {
            tmp = append(tmp, products[j])
        }
        ans = append(ans, tmp)
    }
    return ans
}
