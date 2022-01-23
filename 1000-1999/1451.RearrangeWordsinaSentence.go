import (
	"sort"
	"strings"
)

type word_idx struct {
	word string
	idx  int
}

func arrangeWords(text string) string {
	words := strings.Split(text, " ")
	words[0] = strings.ToLower(words[0])
	word_idxs := make([]word_idx, len(words))
	for i, w := range words {
		word_idxs[i] = word_idx{word: w, idx: i}
	}
	sort.Slice(word_idxs, func(i, j int) bool {
		if len(word_idxs[i].word) == len(word_idxs[j].word) {
			return word_idxs[i].idx < word_idxs[j].idx
		}
		return len(word_idxs[i].word) < len(word_idxs[j].word)
	})
	for i, word_idx := range word_idxs {
		words[i] = word_idx.word
	}
	words[0] = strings.ToUpper(words[0][:1]) + words[0][1:]
	return strings.Join(words, " ")
}