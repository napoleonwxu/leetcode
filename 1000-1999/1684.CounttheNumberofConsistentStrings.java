class Solution {
    public int countConsistentStrings(String allowed, String[] words) {
        Set<Character> set = new HashSet<>();
        for (int i = 0; i < allowed.length(); i++) {
            set.add(allowed.charAt(i));
        }
        int cnt = 0;
        for (String word: words) {
            boolean consistent = true;
            for (int i = 0; i < word.length(); i++) {
                if (!set.contains(word.charAt(i))) {
                    consistent = false;
                    break;
                }
            }
            if (consistent) {
                cnt++;
            }
        }
        return cnt;
    }
}