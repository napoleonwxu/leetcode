class Solution {
    public int minimumDeletions(String s) {
        int b = 0, cnt = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'a') {
                if (b > 0) {
                    b--;
                    cnt++;
                }
            } else {
                b++;
            }
        }
        return cnt;
    }
}