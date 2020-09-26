class Solution {
    public String thousandSeparator(int n) {
        if (n < 1000) {
            return String.valueOf(n);
        }
        String ans = "";
        int cnt = 0;
        while (n > 0) {
            ans = String.valueOf(n%10) + ans;
            n /= 10;
            cnt++;
            if (cnt == 3 && n > 0) {
                ans = "." + ans;
                cnt = 0;
            }
        }
        return ans;
    }
}