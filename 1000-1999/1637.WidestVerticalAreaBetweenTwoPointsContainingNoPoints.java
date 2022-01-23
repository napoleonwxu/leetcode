class Solution {
    public int maxWidthOfVerticalArea(int[][] points) {
        Set<Integer> set = new HashSet<>();
        for (int[] point : points) {
            set.add(point[0]);
        }
        List<Integer> x = new ArrayList<>(set);
        Collections.sort(x);
        int ans = 0;
        for (int i = 1; i < x.size(); i++) {
            ans = Math.max(ans, x.get(i) - x.get(i-1));
        }
        return ans;
    }
}