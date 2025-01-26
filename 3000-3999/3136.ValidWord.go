func isValid(word string) bool {
    if len(word) < 3 {
        return false
    }
    vowels := make(map[byte]bool, 10)
    vowels['a'] = true
    vowels['e'] = true
    vowels['i'] = true
    vowels['o'] = true
    vowels['u'] = true
    vowels['A'] = true
    vowels['E'] = true
    vowels['I'] = true
    vowels['O'] = true
    vowels['U'] = true
    vowel, consonant := false, false
    bytes := []byte(word)
    for _, ch := range bytes {
        if ch >= '0' && ch <= '9' {
            continue
        } else if vowels[ch] {
            vowel = true
        } else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
            consonant = true
        } else {
            return false
        }
    }
    return vowel && consonant
}