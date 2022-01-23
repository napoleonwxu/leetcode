/*
There are N courses, labelled from 1 to N.
We are given relations[i] = [X, Y], representing a prerequisite relationship between course X and course Y: course X has to be studied before course Y.
In one semester you can study any number of courses as long as you have studied all the prerequisites for the course you are studying.
Return the minimum number of semesters needed to study all courses.  If there is no way to study all the courses, return -1.

Example 1:
Input: N = 3, relations = [[1,3],[2,3]]
Output: 2
Explanation: 
In the first semester, courses 1 and 2 are studied. In the second semester, course 3 is studied.

Example 2:
Input: N = 3, relations = [[1,2],[2,3],[3,1]]
Output: -1
Explanation: 
No course can be studied because they depend on each other.
 
Note:
1 <= N <= 5000
1 <= relations.length <= 5000
relations[i][0] != relations[i][1]
There are no repeated relations in the input.
*/

func minimumSemesters(N int, relations [][]int) int {
    // BFS
    cur := make(map[int]bool)
    // In case some courses are not in relations
    for c := 1; c <= N; c++ {
        cur[c] = true
    }
    nxt := make([][]int, N+1)
    pre_cnt := make([]int, N+1)
    for _, r := range relations {
        if cur[r[1]] {
            delete(cur, r[1])
        }
        nxt[r[0]] = append(nxt[r[0]], r[1])
        pre_cnt[r[1]]++
    }
    ans, cnt := 0, 0
    for len(cur) > 0 {
        tmp := make(map[int]bool)
        for c, _ := range cur {
            for _, n_c := range nxt[c] {
                pre_cnt[n_c]--
                if pre_cnt[n_c] == 0 {
                    tmp[n_c] = true
                }
            }
        }
        ans++
        cnt += len(cur)
        cur = tmp
    }
    if cnt == N {
        return ans
    }
    return -1
}

