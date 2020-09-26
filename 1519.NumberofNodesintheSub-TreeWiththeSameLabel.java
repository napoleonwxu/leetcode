class Solution {
    private int n;
    private String labels;
    private int[] ans;
    private HashMap<Integer, ArrayList<Integer>> graph;
    
    public int[] countSubTrees(int n, int[][] edges, String labels) {
        this.n = n;
        this.labels = labels;
        this.ans = new int[n];
        this.graph = new HashMap();
        for (int[] edge : edges) {
            graph.computeIfAbsent(edge[0], k -> new ArrayList<>()).add(edge[1]);
            graph.computeIfAbsent(edge[1], k -> new ArrayList<>()).add(edge[0]);
        }
        
        dfs(0, -1);
        return ans;
    }
    
    private int[] dfs(int node, int parent) {
        int[] cnt = new int[26];
        for (int child : graph.getOrDefault(node, new ArrayList<>())) {
            if (child != parent) {
                int[] sub = dfs(child, node);
                for (int i = 0; i < 26; i++) {
                    cnt[i] += sub[i];
                }
            }
        }
        char ch = labels.charAt(node);
        cnt[ch-'a']++;
        ans[node] = cnt[ch-'a'];
        return cnt;
    }
}