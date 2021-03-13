class Solution {
    public int[] frequencySort(int[] nums) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int num : nums) {
            map.put(num, map.getOrDefault(num, 0) + 1);
        }
        var pq = new PriorityQueue<Integer>((i, j) -> map.get(i)==map.get(j) ? j-i : map.get(i)-map.get(j));
        for (int num :nums) {
            pq.offer(num);
        }
        int[] ans = new int[nums.length];
        int i = 0;
        while (!pq.isEmpty()) {
            ans[i++] = pq.poll();
        }
        return ans;
    }
}