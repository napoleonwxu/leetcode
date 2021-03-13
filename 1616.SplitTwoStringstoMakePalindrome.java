class Solution {
    public boolean checkPalindromeFormation(String a, String b) {
        return check(a, b) || check(b, a);
    }
    
    private boolean check(String a, String b) {
        for (int i = 0, j = a.length()-1; i < j; i++, j--) {
            if (a.charAt(i) != b.charAt(j)) {
                return isPalindrome(a, i, j) || isPalindrome(b, i, j);
            }
        }
        return true;
    }

    private boolean isPalindrome(String s, int i, int j) {
        for (; i < j && s.charAt(i) == s.charAt(j); i++, j--) {}
        return i >= j;
    }
}