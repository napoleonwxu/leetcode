func distributeCandies(candies int, num_people int) []int {
    ans := make([]int, num_people)
    c := 1
    for candies > 0 {
        for i := 0; i < num_people; i++ {
            if c <= candies {
                ans[i] += c
                candies -= c
                c++
            } else {
                ans[i] += candies
                candies = 0
            }
        }
    }
    return ans
}