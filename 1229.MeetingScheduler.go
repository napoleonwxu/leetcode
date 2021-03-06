/*
Given the availability time slots arrays slots1 and slots2 of two people and a meeting duration duration, return the earliest time slot that works for both of them and is of duration duration.
If there is no common time slot that satisfies the requirements, return an empty array.
The format of a time slot is an array of two elements [start, end] representing an inclusive time range from start to end.  
It is guaranteed that no two availability slots of the same person intersect with each other. That is, for any two time slots [start1, end1] and [start2, end2] of the same person, either start1 > end2 or start2 > end1.

Example 1:
Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 8
Output: [60,68]

Example 2:
Input: slots1 = [[10,50],[60,120],[140,210]], slots2 = [[0,15],[60,70]], duration = 12
Output: []

Constraints:
1 <= slots1.length, slots2.length <= 10^4
slots1[i].length, slots2[i].length == 2
slots1[i][0] < slots1[i][1]
slots2[i][0] < slots2[i][1]
0 <= slots1[i][j], slots2[i][j] <= 10^9
1 <= duration <= 10^6 
*/

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
    // solution2, meeting room: O(2nlog2n), n: len(slot1)+len(slot2)
    times := make([][]int, 0, 2*(len(slots1)+len(slots2)))
    for _, slot := range slots1 {
        times = append(times, []int{slot[0], 1}, []int{slot[1], -1})
    }
    for _, slot := range slots2 {
        times = append(times, []int{slot[0], 1}, []int{slot[1], -1})
    }
    sort.Slice(times, func(i, j int) bool {
        return times[i][0] < times[j][0]
    })
    now, cnt := 0, 0
    for _, time := range times {
        if cnt == 2 && now+duration <= time[0] {
            return []int{now, now+duration}
        }
        now = time[0]
        cnt += time[1]
    }
    /* solution1: O(nlogn), n: len(slot1)+len(slot2)
    slot := append(slots1, slots2...)
    sort.Slice(slot, func(i, j int) bool {
        return slot[i][0] < slot[j][0]
    })
    for i := 1; i < len(slot); i++ {
        if slot[i][1]-slot[i][0] >= duration && slot[i-1][1]-slot[i][0] >= duration {
            return []int{slot[i][0], slot[i][0]+duration}
        }
    } */
    return nil
}
