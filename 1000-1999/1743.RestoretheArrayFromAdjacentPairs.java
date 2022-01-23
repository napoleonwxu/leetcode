class Solution {
    public int[] restoreArray(int[][] adjacentPairs) {
        Map<Integer, Set<Integer>> map = new HashMap<>();
        for (int[] adj : adjacentPairs) {
            map.computeIfAbsent(adj[0], v -> new HashSet<>()).add(adj[1]);
            map.computeIfAbsent(adj[1], v -> new HashSet<>()).add(adj[0]);
        }
        int n = adjacentPairs.length + 1;
        int[] ans = new int[n];
        for (int k : map.keySet()) {
            if (map.get(k).size() == 1) {
                ans[0] = k;
                break;
            }
        }
        int pre = 1 << 31, cur = ans[0];
        for (int i = 1; i < n; i++) {
            for (int next : map.remove(cur)) {
                if (next != pre) {
                    ans[i] = next;
                    pre = cur;
                    cur = next;
                    break;
                }
            }
        }
        return ans;
    }
}