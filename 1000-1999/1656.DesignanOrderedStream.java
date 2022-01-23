class OrderedStream {
    private String[] strs;
    private int n, ptr;

    public OrderedStream(int n) {
        strs = new String[n];
        ptr = 0;
        this.n = n;
    }
    
    public List<String> insert(int id, String value) {
        strs[id-1] = value;
        List<String> ans = new ArrayList<>();
        for (; ptr < n && strs[ptr] != null; ptr++) {
            ans.add(strs[ptr]);
        }
        return ans;
    }
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * OrderedStream obj = new OrderedStream(n);
 * List<String> param_1 = obj.insert(id,value);
 */