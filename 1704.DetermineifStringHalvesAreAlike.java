class Solution {
    public boolean halvesAreAlike(String s) {
        var set = Set.of('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U');
        int pre = 0, tail = 0, n = s.length();
        for (int i = 0; i < n/2; i++) {
            pre += set.contains(s.charAt(i)) ? 1 : 0;
            tail += set.contains(s.charAt(n-1-i)) ? 1 : 0;
        }
        return pre == tail;
    }
}