class Solution {
    public List<Integer> findSmallestSetOfVertices(int n, List<List<Integer>> edges) {
        HashSet<Integer> toNodes = new HashSet<>();
        for (List<Integer> edge : edges) {
            toNodes.add(edge.get(1));
        }
        List<Integer> ans = new ArrayList<>();
        for (int node = 0; node < n; node++) {
            if (!toNodes.contains(node)) {
                ans.add(node);
            }
        }
        return ans;
    }
}