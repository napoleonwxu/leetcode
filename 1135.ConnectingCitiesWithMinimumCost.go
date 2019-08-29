/*
There are N cities numbered from 1 to N.
You are given connections, where each connections[i] = [city1, city2, cost] represents the cost to connect city1 and city2 together.  (A connection is bidirectional: connecting city1 and city2 is the same as connecting city2 and city1.)
Return the minimum cost so that for every pair of cities, there exists a path of connections (possibly of length 1) that connects those two cities together.  The cost is the sum of the connection costs used. If the task is impossible, return -1.

Example 1:
Input: N = 3, connections = [[1,2,5],[1,3,6],[2,3,1]]
Output: 6
Explanation: 
Choosing any 2 edges will connect all cities so we choose the minimum 2.

Example 2:
Input: N = 4, connections = [[1,2,3],[3,4,4]]
Output: -1
Explanation: 
There is no way to connect all cities even if all edges are used.

Note:
1 <= N <= 10000
1 <= connections.length <= 10000
1 <= connections[i][0], connections[i][1] <= N
0 <= connections[i][2] <= 10^5
connections[i][0] != connections[i][1]
*/

// https://mp.weixin.qq.com/s/5oPMt3cUetqWBu6hX0ChYw
func minimumCost(N int, conections [][]int) int {
    if len(conections) < N-1 {
        return -1
    }
    sort.Slice(conections, func(i, j int) bool {
        return conections[i][2] < conections[j][2]
    })
    
    par := make([]int, N+1)
    sz := make([]int, N+1)
    for i := range par {
        par[i] = i
        sz[i] = 1
    }
    
    ans, cnt := 0, 0
    for _, con := range conections {
        if find(par, con[0]) != find(par, con[1]) {
            ans += con[2]
            cnt++
            if cnt == N-1 {
                return ans
            }
            union(par, sz, con[0], con[1])
        }
    }
    return -1
}

func find(par []int, city int) int {
    if par[city] != city {
        par[city] = find(par, par[city])
    }
    return par[city]
}

func union(par, sz []int, c1, c2 int) {
    p1, p2 := find(par, c1), find(par, c2)
    if p1 == p2 {
        return
    }
    if par[p1] < par[p2] {
        p1, p2 = p2, p1
    }
    par[p2] = p1
    sz[p1] += sz[p2]
}
