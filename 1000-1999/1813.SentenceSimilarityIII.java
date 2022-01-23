class Solution {
    public boolean areSentencesSimilar(String sentence1, String sentence2) {
        String[] s1 = sentence1.split(" "), s2 = sentence2.split(" ");
        int n1 = s1.length, n2 = s2.length, i = 0;
        if (n1 > n2) {
            return areSentencesSimilar(sentence2, sentence1);
        }
        for (; i < n1 && s1[i].equals(s2[i]); i++) {}
        for (; i < n1 && s1[i].equals(s2[i+n2-n1]); i++) {}
        return i == n1;
    }
}