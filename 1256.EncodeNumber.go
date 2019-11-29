/*
Given a non-negative integer num, Return its encoding string.
The encoding is done by converting the integer to a string using a secret function that you should deduce from the following table:

n	f(n)
0	""
1	"0"
2	"1"
3	"00"
4	"01"
5	"10"
6	"11"
7	"000"

Example 1:
Input: num = 23
Output: "1000"

Example 2:
Input: num = 107
Output: "101100"

Constraints:
0 <= num <= 10^9
*/

func encode(num int) string {
    return bin(num+1)[1:]
    /*
    var width uint
    for width = 0; 1<<width <= num+1; width++ {}
    width--
    return toBin(num-1<<width+1, width)
    */
}

func bin(num int) string {
    ans := ""
    for num > 0 {
        ans = strconv.Itoa(num&1) + ans
        num >>= 1
    }
    return ans
}

func toBin(num int, width uint) string {
    ans := make([]byte, width)
    for i := int(width-1); i >= 0; i-- {
        ans[i] = byte(num&1) + '0'
        num >>= 1
    }
    return string(ans)
}
