import (
    "strconv"
    "strings"
)

func discountPrices(sentence string, discount int) string {
    words := strings.Split(sentence, " ")
    for i, word := range words {
        price, err := strconv.Atoi(word[1:])
        if word[0] == '$' && err == nil {
            words[i] = "$" + strconv.FormatFloat(float64(price*(100-discount))/100, 'f', 2, 64)
        }
    }
    return strings.Join(words, " ")
}