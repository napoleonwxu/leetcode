/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */
 func findSecretWord(wordlist []string, master *Master) {
    for i := 0; i < 10; i++ {
        guess := wordlist[rand.Intn(len(wordlist))]
        x := master.guess(guess)
        if x == 6 {
            return
        }
        nxt := make([]string, 0, len(wordlist))
        for _, word := range wordlist {
            if match(word, guess) == x {
                nxt = append(nxt, word)
            }
        }
        wordlist = nxt
    }
}

func match(word1, word2 string) int {
    cnt := 0
    for i := 0; i < len(word1); i++ {
        if word1[i] == word2[i] {
            cnt++
        }
    }
    return cnt
}