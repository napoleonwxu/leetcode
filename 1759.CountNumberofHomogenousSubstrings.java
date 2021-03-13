class Solution {
    private static final int MOD = (int)1e9 + 7;
    
    public int countHomogenous(String s) {
        int ans = 0, cnt = 0;
        for (int i = 0; i < s.length(); i++) {
            cnt = i>0 && s.charAt(i)==s.charAt(i-1) ? cnt+1 : 1;
            ans = (ans + cnt) % MOD;
        }
        return ans;
    }
}