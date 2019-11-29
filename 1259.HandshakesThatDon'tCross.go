/*
You are given an even number of people num_people that stand around a circle and each person shakes hands with someone else, so that there are num_people / 2 handshakes total.
Return the number of ways these handshakes could occur such that none of the handshakes cross.
Since this number could be very big, return the answer mod 10^9 + 7

Example 1:
Input: num_people = 2
Output: 1

Example 2:
Input: num_people = 4
Output: 2
Explanation: There are two ways to do it, the first way is [(1,2),(3,4)] and the second one is [(2,3),(4,1)].

Example 3:
Input: num_people = 6
Output: 5

Example 4:
Input: num_people = 8
Output: 14
 
Constraints:
2 <= num_people <= 1000
num_people % 2 == 0
*/

//https://coordinate.wang/index.php/archives/2765/
const mod = 1e9 + 7

func numberOfWays(num_people int) int {
    inv := make([]int, num_people/2 + 2)
    inv[1] = 1
    res := 1
    for i := 2; i < num_people/2 + 2; i++ {
        inv[i] = mod - mod / i * inv[mod%i] % mod
    }
    for i := 1; i < num_people/2 + 1; i++ {
        res = res * (i + num_people/2) % mod
        res = res * inv[i] % mod
    }
    return res * inv[num_people/2 + 1] % mod
}