/*
Given two strings str1 and str2 of the same length, determine whether you can transform str1 into str2 by doing zero or more conversions.
In one conversion you can convert all occurrences of one character in str1 to any other lowercase English character.
Return true if and only if you can transform str1 into str2.

Example 1:
Input: str1 = "aabcc", str2 = "ccdee"
Output: true
Explanation: Convert 'c' to 'e' then 'b' to 'd' then 'a' to 'c'. Note that the order of conversions matter.

Example 2:
Input: str1 = "leetcode", str2 = "codeleet"
Output: false
Explanation: There is no way to transform str1 to str2.
 
Note:
1 <= str1.length == str2.length <= 10^4
Both str1 and str2 contain only lowercase English letters.
*/

func canConvert(str1 string, str2 string) bool {
    if str1 == str2 {
        return true
    }
    // If str1 != str2 and str2 has all 26 letters, return false
    Map := make(map[byte]bool)
    for i := 0; i < len(str2); i++ {
        Map[str2[i]] = true
    }
    if len(Map) >= 26 {
        return false
    }
    
    trans := make(map[byte]byte)
    for i := 0; i < len(str1); i++ {
        to, ok := trans[str1[i]]
        if ok && str2[i] != to {
            return false
        }
        trans[str1[i]] = str2[i]
    }
    return true
}