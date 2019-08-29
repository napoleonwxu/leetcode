func intToRoman(num int) string {
    Int := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    Roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    ans := []string{}
    for i := 0; i < len(Int); i++ {
        for num >= Int[i] {
            ans = append(ans, Roman[i])
            num -= Int[i]
        }
    }
    return strings.Join(ans, "")
}