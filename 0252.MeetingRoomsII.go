/*
Given an array of meeting time intervals consisting of start and end times [[s1,e1],[s2,e2],...] (si < ei), find the minimum number of conference rooms required.

Example 1:
Input: [[0, 30],[5, 10],[15, 20]]
Output: 2

Example 2:
Input: [[7,10],[2,4]]
Output: 1
*/

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	time := make([][]int, 0, n<<1)
	for _, interval := intervals {
		time = append(time, []int{interval[0], 1}, []int{interval[1], -1})
	}
	sort.Slice(time, func(i, j int) bool {
		return time[i][0] < time[j][0]
	})
	ans, cnt := 0, 0
	for _, t := range time {
		cnt += t[1]
		ans = max(ans, cnt)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}