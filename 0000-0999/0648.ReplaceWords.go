func replaceWords(dict []string, sentence string) string {
    Map := make(map[string]bool, len(dict))
    for _, root := range dict {
        Map[root] = true
    }
    words := strings.Split(sentence, " ")
    for i, word := range words {
        for j := 1; j < len(word); j++ {
            if Map[word[:j]] {
                words[i] = word[:j]
                break
            }
        }
    }    
    /* Without hashmap
    words := strings.Split(sentence, " ")
    for i := range words {
        for _, root := range dict {
            if strings.HasPrefix(words[i], root) {
                words[i] = root
            }
        }
    }
    */
    return strings.Join(words, " ")
}