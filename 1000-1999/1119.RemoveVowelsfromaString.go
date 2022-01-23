/*
Given a string S, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and return the new string.
Note:
S consists of lowercase English letters only.
1 <= S.length <= 1000
*/
func removeVowels(S string) string {
    vowels := []string{"a", "e", "i", "o", "u"}
    for _, vowel := range vowels {
        S = strings.Replace(S, vowel, "", -1)
    }
    return S
}