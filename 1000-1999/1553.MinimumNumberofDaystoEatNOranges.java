class Solution {
    private HashMap<Integer, Integer> map = new HashMap<>();
    
    public int minDays(int n) {
        if (n <= 1) {
            return n;
        }
        if (!map.containsKey(n)) {
            map.put(n, 1 + Math.min(n%2+minDays(n/2), n%3+minDays(n/3)));
        }
        return map.get(n);
    }
}