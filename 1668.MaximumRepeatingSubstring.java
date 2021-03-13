class Solution {
    public int maxRepeating(String sequence, String word) {
        int ans = 0, sub = 0, i = 0;
        //while (sequence.contains(word.repeat(ans))) ans++;
        //return ans-1;
        while (i < sequence.length()) {
            if (isSub(sequence, i, word)) {
                sub++;
                ans = Math.max(ans, sub);
                i += word.length();
            } else {
                sub = 0;
                i++;
            }
        }
        return ans;
    }
    
    private boolean isSub(String seq, int idx, String word) {
        for (int i = 0; i < word.length(); i++) {
            if (idx+i >= seq.length() || seq.charAt(idx+i) != word.charAt(i)) {
                return false;
            }
        }
        return true;
    }
}